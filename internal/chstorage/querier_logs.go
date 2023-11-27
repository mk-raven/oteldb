package chstorage

import (
	"context"
	"fmt"
	"strings"

	"github.com/ClickHouse/ch-go"
	"github.com/ClickHouse/ch-go/proto"
	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/go-faster/sdk/zctx"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"github.com/go-faster/oteldb/internal/iterators"
	"github.com/go-faster/oteldb/internal/logql"
	"github.com/go-faster/oteldb/internal/logql/logqlengine"
	"github.com/go-faster/oteldb/internal/logstorage"
	"github.com/go-faster/oteldb/internal/otelstorage"
)

var _ logstorage.Querier = (*Querier)(nil)
var _ logqlengine.Querier = (*Querier)(nil)

// LabelNames implements logstorage.Querier.
func (q *Querier) LabelNames(ctx context.Context, opts logstorage.LabelsOptions) (_ []string, rerr error) {
	table := q.tables.Logs

	ctx, span := q.tracer.Start(ctx, "LabelNames",
		trace.WithAttributes(
			attribute.Int64("chstorage.start_range", int64(opts.Start)),
			attribute.Int64("chstorage.end_range", int64(opts.End)),
			attribute.String("chstorage.table", table),
		),
	)
	defer func() {
		if rerr != nil {
			span.RecordError(rerr)
		}
		span.End()
	}()
	var (
		names proto.ColStr
		out   []string
	)
	if err := q.ch.Do(ctx, ch.Query{
		Logger: zctx.From(ctx).Named("ch"),
		Result: proto.Results{
			{Name: "key", Data: &names},
		},
		OnResult: func(ctx context.Context, block proto.Block) error {
			for i := 0; i < names.Rows(); i++ {
				out = append(out, names.Row(i))
			}
			return nil
		},
		Body: fmt.Sprintf(`SELECT DISTINCT 
arrayJoin(arrayConcat(JSONExtractKeys(attributes), JSONExtractKeys(resource), JSONExtractKeys(scope_attributes))) as key
FROM %s 
WHERE (toUnixTimestamp64Nano(timestamp) >= %d AND toUnixTimestamp64Nano(timestamp) <= %d)
LIMIT 1000`,
			table, opts.Start, opts.End,
		),
	}); err != nil {
		return nil, errors.Wrap(err, "select")
	}

	return out, nil
}

type labelStaticIterator struct {
	name   string
	values []jx.Raw
}

func (l *labelStaticIterator) Next(t *logstorage.Label) bool {
	if len(l.values) == 0 {
		return false
	}
	t.Name = l.name
	e := jx.DecodeBytes(l.values[0])
	switch e.Next() {
	case jx.String:
		t.Type = int32(pcommon.ValueTypeStr)
		s, _ := e.Str()
		t.Value = s
	case jx.Number:
		n, _ := e.Num()
		if n.IsInt() {
			t.Type = int32(pcommon.ValueTypeInt)
			v, _ := n.Int64()
			t.Value = fmt.Sprintf("%d", v)
		} else {
			t.Type = int32(pcommon.ValueTypeDouble)
			v, _ := n.Float64()
			t.Value = fmt.Sprintf("%f", v)
		}
	default:
		t.Type = int32(pcommon.ValueTypeStr)
		t.Value = l.values[0].String()
	}
	l.values = l.values[1:]
	return true
}

func (l *labelStaticIterator) Err() error   { return nil }
func (l *labelStaticIterator) Close() error { return nil }

// LabelValues implements logstorage.Querier.
func (q *Querier) LabelValues(ctx context.Context, labelName string, opts logstorage.LabelsOptions) (_ iterators.Iterator[logstorage.Label], rerr error) {
	table := q.tables.Logs

	ctx, span := q.tracer.Start(ctx, "LabelValues",
		trace.WithAttributes(
			attribute.Int64("chstorage.start_range", int64(opts.Start)),
			attribute.Int64("chstorage.end_range", int64(opts.End)),
			attribute.String("chstorage.table", table),
		),
	)
	defer func() {
		if rerr != nil {
			span.RecordError(rerr)
		}
		span.End()
	}()
	var (
		names proto.ColStr
		out   []jx.Raw
	)
	if err := q.ch.Do(ctx, ch.Query{
		Logger: zctx.From(ctx).Named("ch"),
		Result: proto.Results{
			{Name: "value", Data: &names},
		},
		OnResult: func(ctx context.Context, block proto.Block) error {
			for i := 0; i < names.Rows(); i++ {
				out = append(out, jx.Raw(names.Row(i)))
			}
			return nil
		},
		Body: fmt.Sprintf(`SELECT DISTINCT 
COALESCE(
	JSONExtractRaw(attributes, %[1]s), 
	JSONExtractRaw(scope_attributes, %[1]s),
	JSONExtractRaw(resource, %[1]s)
) as value
FROM %s 
WHERE (toUnixTimestamp64Nano(timestamp) >= %d AND toUnixTimestamp64Nano(timestamp) <= %d) LIMIT 1000`,
			singleQuoted(labelName), table, opts.Start, opts.End,
		),
	}); err != nil {
		return nil, errors.Wrap(err, "select")
	}
	return &labelStaticIterator{
		name:   labelName,
		values: out,
	}, nil
}

// Capabilities implements logqlengine.Querier.
func (q *Querier) Capabilities() (caps logqlengine.QuerierCapabilities) {
	caps.Label.Add(logql.OpEq, logql.OpNotEq, logql.OpRe, logql.OpNotRe)
	caps.Line.Add(logql.OpEq, logql.OpNotEq, logql.OpRe, logql.OpNotRe)
	return caps
}

type logStaticIterator struct {
	data []logstorage.Record
}

func (l *logStaticIterator) Next(t *logstorage.Record) bool {
	if len(l.data) == 0 {
		return false
	}
	*t = l.data[0]
	l.data = l.data[1:]
	return true
}

func (l *logStaticIterator) Err() error   { return nil }
func (l *logStaticIterator) Close() error { return nil }

// SelectLogs implements logqlengine.Querier.
func (q *Querier) SelectLogs(ctx context.Context, start, end otelstorage.Timestamp, params logqlengine.SelectLogsParams) (_ iterators.Iterator[logstorage.Record], rerr error) {
	table := q.tables.Logs

	ctx, span := q.tracer.Start(ctx, "SelectLogs",
		trace.WithAttributes(
			attribute.Int("chstorage.labels_count", len(params.Labels)),
			attribute.Int64("chstorage.start_range", int64(start)),
			attribute.Int64("chstorage.end_range", int64(end)),
			attribute.String("chstorage.table", table),
		),
	)
	defer func() {
		if rerr != nil {
			span.RecordError(rerr)
		}
		span.End()
	}()

	out := newLogColumns()
	var query strings.Builder
	query.WriteString("SELECT ")
	for i, column := range out.StaticColumns() {
		if i != 0 {
			query.WriteByte(',')
		}
		query.WriteString(column)
	}
	fmt.Fprintf(&query, " FROM %s WHERE (toUnixTimestamp64Nano(timestamp) >= %d AND toUnixTimestamp64Nano(timestamp) <= %d)", table, start, end)
	for _, m := range params.Labels {
		switch m.Op {
		case logql.OpEq:
			query.WriteString(" AND (")
		case logql.OpNotEq:
			query.WriteString(" AND NOT (")
		default:
			return nil, errors.Errorf("unexpected op %q", m.Op)
		}
		for i, column := range []string{
			"attributes",
			"resource",
			"scope_attributes",
		} {
			if i != 0 {
				query.WriteString(" OR ")
			}
			// TODO: how to match integers, booleans, floats, arrays?
			fmt.Fprintf(&query, "JSONExtractString(%s, %s) = %s", column, singleQuoted(m.Label), singleQuoted(m.Value))
		}
		query.WriteByte(')')
	}
	for _, m := range params.Line {
		switch m.Op {
		case logql.OpEq:
			query.WriteString(" AND ")
		case logql.OpNotEq:
			query.WriteString(" AND NOT ")
		default:
			return nil, errors.Errorf("unexpected op %q", m.Op)
		}

		// Line filter checks if line contains given value.
		fmt.Fprintf(&query, "positionUTF8(body, %s) > 0", singleQuoted(m.Value))
	}

	// TODO: use streaming.
	var data []logstorage.Record

	if err := q.ch.Do(ctx, ch.Query{
		Logger: zctx.From(ctx).Named("ch"),
		Body:   query.String(),
		Result: out.Result(),
		OnResult: func(ctx context.Context, block proto.Block) error {
			if err := out.ForEach(func(r logstorage.Record) {
				data = append(data, r)
			}); err != nil {
				return errors.Wrap(err, "for each")
			}
			return nil
		},
	}); err != nil {
		return nil, errors.Wrap(err, "select")
	}
	return &logStaticIterator{data: data}, nil
}
