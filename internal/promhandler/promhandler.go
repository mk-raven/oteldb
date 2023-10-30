// Package promhandler provides Prometheus API implementation.
package promhandler

import (
	"context"
	"time"

	"github.com/go-faster/errors"
	ht "github.com/ogen-go/ogen/http"
	"github.com/prometheus/prometheus/model/labels"
	"github.com/prometheus/prometheus/promql"
	"github.com/prometheus/prometheus/storage"
	"github.com/prometheus/prometheus/util/annotations"
	"golang.org/x/exp/maps"

	"github.com/go-faster/oteldb/internal/promapi"
)

var _ promapi.Handler = (*PromAPI)(nil)

// Engine is a Prometheus engine interface.
type Engine interface {
	NewInstantQuery(ctx context.Context, q storage.Queryable, opts promql.QueryOpts, qs string, ts time.Time) (promql.Query, error)
	NewRangeQuery(ctx context.Context, q storage.Queryable, opts promql.QueryOpts, qs string, start, end time.Time, interval time.Duration) (promql.Query, error)
}

// PromAPI implements promapi.Handler.
type PromAPI struct {
	eng   Engine
	store storage.Queryable

	lookbackDelta time.Duration
}

// NewPromAPI creates new PromAPI.
func NewPromAPI(eng Engine, store storage.Queryable, opts PromAPIOptions) *PromAPI {
	opts.setDefaults()
	return &PromAPI{
		eng:           eng,
		store:         store,
		lookbackDelta: opts.LookbackDelta,
	}
}

// GetLabelValues implements getLabelValues operation.
// GET /api/v1/label/{label}/values
func (h *PromAPI) GetLabelValues(ctx context.Context, params promapi.GetLabelValuesParams) (*promapi.LabelValuesResponse, error) {
	mint, err := parseOptTimestamp(params.Start, MinTime)
	if err != nil {
		return nil, validationErr("parse start", err)
	}
	maxt, err := parseOptTimestamp(params.End, MaxTime)
	if err != nil {
		return nil, validationErr("parse end", err)
	}
	sets, err := parseLabelMatchers(params.Match)
	if err != nil {
		return nil, validationErr("parse match", err)
	}

	q, err := h.store.Querier(mint.UnixMilli(), maxt.UnixMilli())
	if err != nil {
		return nil, executionErr("get querier", err)
	}

	// Fast path for cases when match[] is not set.
	if len(sets) < 2 {
		var matchers []*labels.Matcher
		if len(sets) > 0 {
			matchers = sets[0]
		}

		values, warnings, err := q.LabelValues(ctx, params.Label, matchers...)
		if err != nil {
			return nil, executionErr("get label values", err)
		}

		return &promapi.LabelValuesResponse{
			Status:   "success",
			Warnings: warnings.AsStrings("", 0),
			Data:     values,
		}, nil
	}

	var (
		data     = map[string]struct{}{}
		warnings annotations.Annotations
	)
	for _, set := range sets {
		vals, w, err := q.LabelValues(ctx, params.Label, set...)
		if err != nil {
			return nil, executionErr("get label values", err)
		}

		for _, val := range vals {
			data[val] = struct{}{}
		}
		warnings = warnings.Merge(w)
	}

	return &promapi.LabelValuesResponse{
		Status:   "success",
		Warnings: warnings.AsStrings("", 0),
		Data:     maps.Keys(data),
	}, nil
}

// GetLabels implements getLabels operation.
//
// GET /api/v1/labels
func (h *PromAPI) GetLabels(ctx context.Context, params promapi.GetLabelsParams) (*promapi.LabelsResponse, error) {
	mint, err := parseOptTimestamp(params.Start, MinTime)
	if err != nil {
		return nil, validationErr("parse start", err)
	}
	maxt, err := parseOptTimestamp(params.End, MaxTime)
	if err != nil {
		return nil, validationErr("parse end", err)
	}
	sets, err := parseLabelMatchers(params.Match)
	if err != nil {
		return nil, validationErr("parse match", err)
	}

	q, err := h.store.Querier(mint.UnixMilli(), maxt.UnixMilli())
	if err != nil {
		return nil, executionErr("get querier", err)
	}

	// Fast path for cases when match[] is not set.
	if len(sets) < 2 {
		var matchers []*labels.Matcher
		if len(sets) > 0 {
			matchers = sets[0]
		}

		values, warnings, err := q.LabelNames(ctx, matchers...)
		if err != nil {
			return nil, executionErr("label names", err)
		}

		return &promapi.LabelsResponse{
			Status:   "success",
			Warnings: warnings.AsStrings("", 0),
			Data:     values,
		}, nil
	}

	var (
		data     = map[string]struct{}{}
		warnings annotations.Annotations
	)
	for _, set := range sets {
		vals, w, err := q.LabelNames(ctx, set...)
		if err != nil {
			return nil, executionErr("get label names", err)
		}

		for _, val := range vals {
			data[val] = struct{}{}
		}
		warnings = warnings.Merge(w)
	}

	return &promapi.LabelsResponse{
		Status:   "success",
		Warnings: warnings.AsStrings("", 0),
		Data:     maps.Keys(data),
	}, nil
}

// PostLabels implements postLabels operation.
//
// POST /api/v1/labels
func (h *PromAPI) PostLabels(ctx context.Context, req *promapi.LabelsForm) (*promapi.LabelsResponse, error) {
	return h.GetLabels(ctx, promapi.GetLabelsParams{
		Start: req.Start,
		End:   req.End,
		Match: req.Match,
	})
}

// GetQuery implements getQuery operation.
//
// Query Prometheus.
//
// GET /api/v1/query
func (h *PromAPI) GetQuery(ctx context.Context, params promapi.GetQueryParams) (*promapi.QueryResponse, error) {
	t, err := parseOptTimestamp(params.Time, time.Now())
	if err != nil {
		return nil, validationErr("parse time", err)
	}
	opts, err := parseQueryOpts(params.LookbackDelta, params.Stats, h.lookbackDelta)
	if err != nil {
		return nil, err
	}

	rawQuery := params.Query
	q, err := h.eng.NewInstantQuery(ctx, h.store, opts, rawQuery, t)
	if err != nil {
		return nil, executionErr("make instant query", err)
	}
	defer q.Close()

	r := q.Exec(ctx)
	return mapResult(rawQuery, r)
}

// PostQuery implements postQuery operation.
//
// Query Prometheus.
//
// POST /api/v1/query
func (h *PromAPI) PostQuery(ctx context.Context, req *promapi.QueryForm) (*promapi.QueryResponse, error) {
	return h.GetQuery(ctx, promapi.GetQueryParams{
		Query:         req.Query,
		Time:          req.Time,
		LookbackDelta: req.LookbackDelta,
		Stats:         req.Stats,
	})
}

// GetQueryRange implements getQueryRange operation.
//
// Query Prometheus.
//
// GET /api/v1/query_range
func (h *PromAPI) GetQueryRange(ctx context.Context, params promapi.GetQueryRangeParams) (*promapi.QueryResponse, error) {
	start, err := parseTimestamp(params.Start)
	if err != nil {
		return nil, validationErr("parse start", err)
	}
	end, err := parseTimestamp(params.End)
	if err != nil {
		return nil, validationErr("parse end", err)
	}
	step, err := parseStep(params.Step)
	if err != nil {
		return nil, validationErr("parse step", err)
	}
	opts, err := parseQueryOpts(params.LookbackDelta, params.Stats, h.lookbackDelta)
	if err != nil {
		return nil, err
	}

	if end.Before(start) {
		err := errors.New("end timestamp must not be before start time")
		return nil, validationErr("check range", err)
	}

	rawQuery := params.Query
	q, err := h.eng.NewRangeQuery(ctx, h.store, opts, rawQuery, start, end, step)
	if err != nil {
		return nil, executionErr("make range query", err)
	}
	defer q.Close()

	r := q.Exec(ctx)
	return mapResult(rawQuery, r)
}

// PostQueryRange implements postQueryRange operation.
//
// Query Prometheus.
//
// POST /api/v1/query_range
func (h *PromAPI) PostQueryRange(ctx context.Context, req *promapi.QueryRangeForm) (*promapi.QueryResponse, error) {
	return h.GetQueryRange(ctx, promapi.GetQueryRangeParams{
		Query:         req.Query,
		Start:         req.Start,
		End:           req.End,
		Step:          req.Step,
		LookbackDelta: req.LookbackDelta,
		Stats:         req.Stats,
	})
}

// GetQueryExemplars implements getQueryExemplars operation.
//
// Query Prometheus.
//
// GET /api/v1/query_exemplars
func (h *PromAPI) GetQueryExemplars(context.Context, promapi.GetQueryExemplarsParams) (*promapi.QueryExemplarsResponse, error) {
	return nil, ht.ErrNotImplemented
}

// PostQueryExemplars implements postQueryExemplars operation.
//
// Query Prometheus.
//
// POST /api/v1/query_exemplars
func (h *PromAPI) PostQueryExemplars(context.Context) (*promapi.QueryExemplarsResponse, error) {
	return nil, ht.ErrNotImplemented
}

// GetMetadata implements getMetadata operation.
//
// GET /api/v1/metadata
func (h *PromAPI) GetMetadata(context.Context, promapi.GetMetadataParams) (*promapi.MetadataResponse, error) {
	return nil, ht.ErrNotImplemented
}

// GetRules implements getRules operation.
//
// GET /api/v1/rules
func (h *PromAPI) GetRules(context.Context, promapi.GetRulesParams) (*promapi.RulesResponse, error) {
	return nil, ht.ErrNotImplemented
}

// GetSeries implements getSeries operation.
//
// Query Prometheus.
//
// GET /api/v1/series
func (h *PromAPI) GetSeries(context.Context, promapi.GetSeriesParams) (*promapi.SeriesResponse, error) {
	return nil, ht.ErrNotImplemented
}

// PostSeries implements postSeries operation.
//
// Query Prometheus.
//
// POST /api/v1/series
func (h *PromAPI) PostSeries(context.Context) (*promapi.SeriesResponse, error) {
	return nil, ht.ErrNotImplemented
}

// NewError creates *FailStatusCode from error returned by handler.
//
// Used for common default response.
func (h *PromAPI) NewError(_ context.Context, err error) *promapi.FailStatusCode {
	if _, ok := errors.Into[promql.ErrQueryCanceled](err); ok || errors.Is(err, context.Canceled) {
		return fail(promapi.FailErrorTypeCanceled, err)
	}
	if _, ok := errors.Into[promql.ErrQueryTimeout](err); ok || errors.Is(err, context.DeadlineExceeded) {
		return fail(promapi.FailErrorTypeTimeout, err)
	}
	if _, ok := errors.Into[promql.ErrStorage](err); ok {
		return fail(promapi.FailErrorTypeInternal, err)
	}

	if pe, ok := errors.Into[*PromError](err); ok {
		return fail(pe.Kind, err)
	}

	return fail(promapi.FailErrorTypeInternal, err)
}

func fail(kind promapi.FailErrorType, err error) *promapi.FailStatusCode {
	return &promapi.FailStatusCode{
		StatusCode: promapi.FailToCode(kind),
		Response: promapi.Fail{
			Status:    "error",
			Error:     err.Error(),
			ErrorType: kind,
		},
	}
}
