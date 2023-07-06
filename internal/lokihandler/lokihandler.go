// Package lokihandler provides Loki API implementation.
package lokihandler

import (
	"context"
	"net/http"
	"time"

	ht "github.com/ogen-go/ogen/http"
	"go.uber.org/zap"

	"github.com/go-faster/errors"
	"github.com/go-faster/sdk/zctx"

	"github.com/go-faster/oteldb/internal/iterators"
	"github.com/go-faster/oteldb/internal/logql/logqlengine"
	"github.com/go-faster/oteldb/internal/logstorage"
	"github.com/go-faster/oteldb/internal/lokiapi"
	"github.com/go-faster/oteldb/internal/otelstorage"
)

var _ lokiapi.Handler = (*LokiAPI)(nil)

// LokiAPI implements lokiapi.Handler.
type LokiAPI struct {
	q      logstorage.Querier
	engine *logqlengine.Engine
}

// NewLokiAPI creates new LokiAPI.
func NewLokiAPI(q logstorage.Querier, engine *logqlengine.Engine) *LokiAPI {
	return &LokiAPI{
		q:      q,
		engine: engine,
	}
}

// LabelValues implements labelValues operation.
// Get values of label.
//
// GET /loki/api/v1/label/{name}/values
func (h *LokiAPI) LabelValues(ctx context.Context, params lokiapi.LabelValuesParams) (*lokiapi.Values, error) {
	lg := zctx.From(ctx)

	start, end, err := parseTimeRange(
		time.Now(),
		params.Start,
		params.End,
		params.Since,
	)
	if err != nil {
		return nil, errors.Wrap(err, "parse time range")
	}

	iter, err := h.q.LabelValues(ctx, params.Name, logstorage.LabelsOptions{
		Start: otelstorage.NewTimestampFromTime(start),
		End:   otelstorage.NewTimestampFromTime(end),
	})
	if err != nil {
		return nil, errors.Wrap(err, "query")
	}
	defer func() {
		_ = iter.Close()
	}()

	var values []string
	if err := iterators.ForEach(iter, func(tag logstorage.Label) error {
		values = append(values, tag.Value)
		return nil
	}); err != nil {
		return nil, errors.Wrap(err, "map tags")
	}
	lg.Debug("Got tag values",
		zap.String("label_name", params.Name),
		zap.Int("count", len(values)),
	)

	return &lokiapi.Values{
		Status: "success",
		Data:   values,
	}, nil
}

// Labels implements labels operation.
//
// Get labels.
// Used by Grafana to test connection to Loki.
//
// GET /loki/api/v1/labels
func (h *LokiAPI) Labels(ctx context.Context, params lokiapi.LabelsParams) (*lokiapi.Labels, error) {
	lg := zctx.From(ctx)

	start, end, err := parseTimeRange(
		time.Now(),
		params.Start,
		params.End,
		params.Since,
	)
	if err != nil {
		return nil, errors.Wrap(err, "parse time range")
	}

	names, err := h.q.LabelNames(ctx, logstorage.LabelsOptions{
		Start: otelstorage.NewTimestampFromTime(start),
		End:   otelstorage.NewTimestampFromTime(end),
	})
	if err != nil {
		return nil, errors.Wrap(err, "query")
	}
	lg.Debug("Got label names", zap.Int("count", len(names)))

	return &lokiapi.Labels{
		Status: "success",
		Data:   names,
	}, nil
}

// Push implements push operation.
//
// Push data.
//
// POST /loki/api/v1/push
func (h *LokiAPI) Push(context.Context, lokiapi.PushReq) error {
	return ht.ErrNotImplemented
}

// Query implements query operation.
//
// Query.
//
// GET /loki/api/v1/query
func (h *LokiAPI) Query(ctx context.Context, params lokiapi.QueryParams) (*lokiapi.QueryResponse, error) {
	lg := zctx.From(ctx)

	ts, err := parseTimestamp(params.Time.Value, time.Now())
	if err != nil {
		return nil, errors.Wrap(err, "parse time")
	}

	streams, err := h.engine.Eval(ctx, params.Query, logqlengine.EvalParams{
		Start:     otelstorage.NewTimestampFromTime(ts),
		End:       otelstorage.NewTimestampFromTime(ts),
		Step:      0,
		Direction: string(params.Direction.Or("backward")),
		Limit:     params.Limit.Or(100),
	})
	if err != nil {
		return nil, errors.Wrap(err, "eval")
	}
	lg.Debug("Query", zap.Int("streams", len(streams)))

	return &lokiapi.QueryResponse{
		Status: "success",
		Data: lokiapi.QueryResponseData{
			ResultType: lokiapi.QueryResponseDataResultTypeStreams,
			Result:     streams,
		},
	}, nil
}

// QueryRange implements queryRange operation.
//
// Query range.
//
// GET /loki/api/v1/query_range
func (h *LokiAPI) QueryRange(ctx context.Context, params lokiapi.QueryRangeParams) (*lokiapi.QueryResponse, error) {
	lg := zctx.From(ctx)

	start, end, err := parseTimeRange(
		time.Now(),
		params.Start,
		params.End,
		params.Since,
	)
	if err != nil {
		return nil, errors.Wrap(err, "parse time range")
	}

	step, err := parseStep(params.Step, start, end)
	if err != nil {
		return nil, errors.Wrap(err, "parse step")
	}

	streams, err := h.engine.Eval(ctx, params.Query, logqlengine.EvalParams{
		Start:     otelstorage.NewTimestampFromTime(start),
		End:       otelstorage.NewTimestampFromTime(end),
		Step:      step,
		Direction: string(params.Direction.Or("backward")),
		Limit:     params.Limit.Or(100),
	})
	if err != nil {
		return nil, errors.Wrap(err, "eval")
	}
	lg.Debug("Query range", zap.Int("streams", len(streams)))

	return &lokiapi.QueryResponse{
		Status: "success",
		Data: lokiapi.QueryResponseData{
			ResultType: lokiapi.QueryResponseDataResultTypeStreams,
			Result:     streams,
		},
	}, nil
}

// Series implements series operation.
//
// Get series.
//
// GET /loki/api/v1/series
func (h *LokiAPI) Series(context.Context, lokiapi.SeriesParams) (*lokiapi.Maps, error) {
	return nil, ht.ErrNotImplemented
}

// NewError creates *ErrorStatusCode from error returned by handler.
//
// Used for common default response.
func (h *LokiAPI) NewError(_ context.Context, err error) *lokiapi.ErrorStatusCode {
	return &lokiapi.ErrorStatusCode{
		StatusCode: http.StatusBadRequest,
		Response:   lokiapi.Error(err.Error()),
	}
}
