package ytstore

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-logfmt/logfmt"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
	"go.ytsaurus.tech/yt/go/ypath"
	"go.ytsaurus.tech/yt/go/yt"
	"golang.org/x/exp/maps"

	"github.com/go-faster/errors"
	"github.com/go-faster/sdk/zctx"

	"github.com/go-faster/oteldb/internal/tempoapi"
)

// TempoAPI implements tempoapi.Handler.
type TempoAPI struct {
	yc     yt.Client
	tables tables
}

// NewTempoAPI creates new TempoAPI.
func NewTempoAPI(yc yt.Client, prefix ypath.Path) *TempoAPI {
	return &TempoAPI{
		yc:     yc,
		tables: newTables(prefix),
	}
}

var _ tempoapi.Handler = (*TempoAPI)(nil)

// Echo request for testing, issued by Grafana.
//
// GET /api/echo
func (h *TempoAPI) Echo(_ context.Context) (tempoapi.EchoOK, error) {
	return tempoapi.EchoOK{Data: strings.NewReader("echo")}, nil
}

func (h TempoAPI) querySpans(ctx context.Context, query string, cb func(Span) error) error {
	r, err := h.yc.SelectRows(ctx, query, nil)
	if err != nil {
		return errors.Wrap(err, "select")
	}
	defer func() {
		_ = r.Close()
	}()
	for r.Next() {
		var span Span
		if err := r.Scan(&span); err != nil {
			return errors.Wrap(err, "scan")
		}
		if err := cb(span); err != nil {
			return errors.Wrap(err, "callback")
		}
	}
	if err := r.Err(); err != nil {
		return errors.Wrap(err, "iter err")
	}
	return nil
}

func (h *TempoAPI) searchTags(ctx context.Context, params tempoapi.SearchParams) (map[TraceID]tempoapi.TraceSearchMetadata, error) {
	lg := zctx.From(ctx)

	var query strings.Builder
	fmt.Fprintf(&query, "* from [%s] where true", h.tables.spans)
	if s, ok := params.Start.Get(); ok {
		n := s.UnixNano()
		fmt.Fprintf(&query, " and start >= %d", n)
	}
	if s, ok := params.End.Get(); ok {
		n := s.UnixNano()
		fmt.Fprintf(&query, " and end <= %d", n)
	}
	if d, ok := params.MinDuration.Get(); ok {
		n := d.Nanoseconds()
		fmt.Fprintf(&query, " and (end-start) >= %d", n)
	}
	if d, ok := params.MaxDuration.Get(); ok {
		n := d.Nanoseconds()
		fmt.Fprintf(&query, " and (end-start) <= %d", n)
	}
	if tags, ok := params.Tags.Get(); ok {
		d := logfmt.NewDecoder(strings.NewReader(tags))
		for d.ScanRecord() {
			for d.ScanKeyval() {
				if string(d.Key()) == "name" {
					fmt.Fprintf(&query, " and name = %q", d.Value())
					continue
				}

				query.WriteString(" and (")
				for i, column := range []string{
					"attrs",
					"scope_attrs",
					"resource_attrs",
				} {
					if i != 0 {
						query.WriteString(" or ")
					}
					yp := append([]byte{'/'}, d.Key()...)
					yp = append(yp, "/1"...)
					fmt.Fprintf(&query, "try_get_string(%s, %q) = %q", column, yp, d.Value())
				}
				query.WriteByte(')')
			}
		}
		if err := d.Err(); err != nil {
			return nil, errors.Wrap(err, "parse tags")
		}
	}
	fmt.Fprintf(&query, " limit %d", params.Limit.Or(20))

	lg.Debug("Search traces",
		zap.Stringer("query", &query),
	)
	var c metadataCollector
	if err := h.querySpans(ctx, query.String(), c.AddSpan); err != nil {
		return nil, err
	}
	return c.Result(), nil
}

func (h *TempoAPI) queryParentSpans(ctx context.Context, metadatas map[TraceID]tempoapi.TraceSearchMetadata) error {
	traces := map[TraceID]struct{}{}

	for id, m := range metadatas {
		if m.StartTimeUnixNano.IsZero() {
			traces[id] = struct{}{}
		}
	}
	if len(traces) == 0 {
		return nil
	}

	var query strings.Builder
	fmt.Fprintf(&query, "* from [%s] where is_null(parent_span_id) and trace_id in (", h.tables.spans)
	n := 0
	for id := range traces {
		if n != 0 {
			query.WriteByte(',')
		}
		fmt.Fprintf(&query, "%q", id)
		n++
	}
	query.WriteByte(')')

	zctx.From(ctx).Debug("Query missing parent spans",
		zap.Stringer("query", &query),
		zap.Int("count", len(traces)),
	)

	return h.querySpans(ctx, query.String(), func(span Span) error {
		traceID := span.TraceID
		m := metadatas[traceID]
		span.FillTraceMetadata(&m)
		metadatas[traceID] = m
		return nil
	})
}

// Search implements search operation.
// Execute TraceQL query.
//
// GET /api/search
func (h *TempoAPI) Search(ctx context.Context, params tempoapi.SearchParams) (resp *tempoapi.Traces, _ error) {
	ctx = zctx.With(ctx,
		zap.String("q", params.Q.Value),
		zap.String("tags", params.Tags.Value),
	)

	metadatas, err := h.searchTags(ctx, params)
	if err != nil {
		return resp, errors.Wrap(err, "search tags")
	}
	if err := h.queryParentSpans(ctx, metadatas); err != nil {
		return resp, errors.Wrap(err, "query missing parent spans")
	}
	return &tempoapi.Traces{
		Traces: maps.Values(metadatas),
	}, nil
}

// SearchTagValues implements search_tag_values operation.
//
// This endpoint retrieves all discovered values for the given tag, which can be used in search.
//
// GET /api/search/tag/{tag_name}/values
func (h *TempoAPI) SearchTagValues(ctx context.Context, params tempoapi.SearchTagValuesParams) (resp *tempoapi.TagValues, _ error) {
	lg := zctx.From(ctx)

	query := fmt.Sprintf("value from [%s] where name = %q", h.tables.tags, params.TagName)
	r, err := h.yc.SelectRows(ctx, query, nil)
	if err != nil {
		return resp, err
	}
	defer func() {
		_ = r.Close()
	}()

	var values []string
	for r.Next() {
		var tag Tag
		if err := r.Scan(&tag); err != nil {
			return resp, err
		}
		values = append(values, tag.Value)
	}
	if err := r.Err(); err != nil {
		return resp, err
	}
	lg.Debug("Got tag values",
		zap.String("tag_name", params.TagName),
		zap.Int("count", len(values)),
	)

	return &tempoapi.TagValues{
		TagValues: values,
	}, nil
}

// SearchTagValuesV2 implements search_tag_values_v2 operation.
//
// This endpoint retrieves all discovered values and their data types for the given TraceQL
// identifier.
//
// GET /api/v2/search/tag/{tag_name}/values
func (h *TempoAPI) SearchTagValuesV2(ctx context.Context, params tempoapi.SearchTagValuesV2Params) (resp *tempoapi.TagValuesV2, _ error) {
	lg := zctx.From(ctx)

	query := fmt.Sprintf("type, value from [%s] where name = %q", h.tables.tags, params.TagName)
	r, err := h.yc.SelectRows(ctx, query, nil)
	if err != nil {
		return resp, err
	}
	defer func() {
		_ = r.Close()
	}()

	var values []tempoapi.TagValue
	for r.Next() {
		var tag Tag
		if err := r.Scan(&tag); err != nil {
			return resp, err
		}

		// TODO(tdakkota): handle duration/status and things
		// https://github.com/grafana/tempo/blob/991d72281e5168080f426b3f1c9d5c4b88f7c460/modules/ingester/instance_search.go#L379
		var typ string
		switch pcommon.ValueType(tag.Type) {
		case pcommon.ValueTypeStr:
			typ = "string"
		case pcommon.ValueTypeInt:
			typ = "int"
		case pcommon.ValueTypeDouble:
			typ = "float"
		case pcommon.ValueTypeBool:
			typ = "bool"
		case pcommon.ValueTypeBytes:
			typ = "string"
		case pcommon.ValueTypeMap, pcommon.ValueTypeSlice:
			// what?
			continue
		}

		values = append(values, tempoapi.TagValue{
			Type:  typ,
			Value: tag.Value,
		})
	}
	if err := r.Err(); err != nil {
		return resp, err
	}
	lg.Debug("Got tag types and values",
		zap.String("tag_name", params.TagName),
		zap.Int("count", len(values)),
	)

	return &tempoapi.TagValuesV2{
		TagValues: values,
	}, nil
}

// SearchTags implements search_tags operation.
//
// This endpoint retrieves all discovered tag names that can be used in search.
//
// GET /api/search/tags
func (h *TempoAPI) SearchTags(ctx context.Context) (resp *tempoapi.TagNames, _ error) {
	lg := zctx.From(ctx)

	query := fmt.Sprintf("name from [%s]", h.tables.tags)
	r, err := h.yc.SelectRows(ctx, query, nil)
	if err != nil {
		return resp, err
	}
	defer func() {
		_ = r.Close()
	}()

	var names []string
	for r.Next() {
		var tag Tag
		if err := r.Scan(&tag); err != nil {
			return resp, err
		}
		names = append(names, tag.Name)
	}
	if err := r.Err(); err != nil {
		return resp, err
	}
	lg.Debug("Got tag names", zap.Int("count", len(names)))

	return &tempoapi.TagNames{
		TagNames: names,
	}, nil
}

// TraceByID implements traceByID operation.
//
// Querying traces by id.
//
// GET /api/traces/{traceID}
func (h *TempoAPI) TraceByID(ctx context.Context, params tempoapi.TraceByIDParams) (resp tempoapi.TraceByIDRes, _ error) {
	lg := zctx.From(ctx)
	var (
		start = zap.Skip()
		end   = zap.Skip()

		query = fmt.Sprintf("* from [%s] where trace_id = %q", h.tables.spans, params.TraceID[:])
	)

	if s, ok := params.Start.Get(); ok {
		n := s.UnixNano()
		query += fmt.Sprintf(" and start >= %d", n)
		start = zap.Int64("start", n)
	}
	if s, ok := params.End.Get(); ok {
		n := s.UnixNano()
		query += fmt.Sprintf(" and end <= %d", n)
		end = zap.Int64("end", n)
	}
	lg = lg.With(
		zap.Stringer("look_for", params.TraceID),
		start,
		end,
	)

	var c batchCollector
	if err := h.querySpans(ctx, query, c.AddSpan); err != nil {
		return resp, err
	}
	traces := c.Result()
	spanCount := traces.SpanCount()

	lg.Debug("Got trace by ID", zap.Int("span_count", spanCount))
	if spanCount < 1 {
		return &tempoapi.TraceByIDNotFound{}, nil
	}

	m := ptrace.ProtoMarshaler{}
	data, err := m.MarshalTraces(traces)
	if err != nil {
		return resp, errors.Wrap(err, "marshal traces")
	}
	return &tempoapi.TraceByID{Data: bytes.NewReader(data)}, nil
}

// NewError creates *ErrorStatusCode from error returned by handler.
//
// Used for common default response.
func (h *TempoAPI) NewError(_ context.Context, err error) *tempoapi.ErrorStatusCode {
	return &tempoapi.ErrorStatusCode{
		StatusCode: http.StatusBadRequest,
		Response:   tempoapi.Error(err.Error()),
	}
}
