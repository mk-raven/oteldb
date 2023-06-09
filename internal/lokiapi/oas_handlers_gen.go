// Code generated by ogen, DO NOT EDIT.

package lokiapi

import (
	"context"
	"net/http"
	"time"

	"github.com/go-faster/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.19.0"
	"go.opentelemetry.io/otel/trace"

	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/otelogen"
)

// handleGetLabelValuesRequest handles GetLabelValues operation.
//
// Get values of label.
//
// GET /loki/api/v1/label/{name}/values
func (s *Server) handleGetLabelValuesRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("GetLabelValues"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/loki/api/v1/label/{name}/values"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "GetLabelValues",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		s.duration.Record(ctx, elapsedDuration.Microseconds(), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "GetLabelValues",
			ID:   "GetLabelValues",
		}
	)
	params, err := decodeGetLabelValuesParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response *Values
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "GetLabelValues",
			OperationID:   "GetLabelValues",
			Body:          nil,
			Params: middleware.Parameters{
				{
					Name: "X-Grafana-User",
					In:   "header",
				}: params.XGrafanaUser,
				{
					Name: "start",
					In:   "query",
				}: params.Start,
				{
					Name: "end",
					In:   "query",
				}: params.End,
				{
					Name: "name",
					In:   "path",
				}: params.Name,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetLabelValuesParams
			Response = *Values
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetLabelValuesParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetLabelValues(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetLabelValues(ctx, params)
	}
	if err != nil {
		recordError("Internal", err)
		if errRes, ok := errors.Into[*ErrorStatusCode](err); ok {
			encodeErrorResponse(errRes, w, span)
			return
		}
		if errors.Is(err, ht.ErrNotImplemented) {
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
		encodeErrorResponse(s.h.NewError(ctx, err), w, span)
		return
	}

	if err := encodeGetLabelValuesResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// handleQueryRangeRequest handles QueryRange operation.
//
// Query range.
//
// GET /loki/api/v1/query_range
func (s *Server) handleQueryRangeRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("QueryRange"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/loki/api/v1/query_range"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "QueryRange",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		s.duration.Record(ctx, elapsedDuration.Microseconds(), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "QueryRange",
			ID:   "QueryRange",
		}
	)
	params, err := decodeQueryRangeParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response *QueryResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "QueryRange",
			OperationID:   "QueryRange",
			Body:          nil,
			Params: middleware.Parameters{
				{
					Name: "X-Grafana-User",
					In:   "header",
				}: params.XGrafanaUser,
				{
					Name: "start",
					In:   "query",
				}: params.Start,
				{
					Name: "query",
					In:   "query",
				}: params.Query,
				{
					Name: "step",
					In:   "query",
				}: params.Step,
				{
					Name: "limit",
					In:   "query",
				}: params.Limit,
				{
					Name: "end",
					In:   "query",
				}: params.End,
				{
					Name: "direction",
					In:   "query",
				}: params.Direction,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = QueryRangeParams
			Response = *QueryResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackQueryRangeParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.QueryRange(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.QueryRange(ctx, params)
	}
	if err != nil {
		recordError("Internal", err)
		if errRes, ok := errors.Into[*ErrorStatusCode](err); ok {
			encodeErrorResponse(errRes, w, span)
			return
		}
		if errors.Is(err, ht.ErrNotImplemented) {
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
		encodeErrorResponse(s.h.NewError(ctx, err), w, span)
		return
	}

	if err := encodeQueryRangeResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// handleSeriesRequest handles Series operation.
//
// Get series.
//
// GET /loki/api/v1/series
func (s *Server) handleSeriesRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("Series"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/loki/api/v1/series"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "Series",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		s.duration.Record(ctx, elapsedDuration.Microseconds(), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "Series",
			ID:   "Series",
		}
	)
	params, err := decodeSeriesParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response *Maps
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "Series",
			OperationID:   "Series",
			Body:          nil,
			Params: middleware.Parameters{
				{
					Name: "start",
					In:   "query",
				}: params.Start,
				{
					Name: "end",
					In:   "query",
				}: params.End,
				{
					Name: "match",
					In:   "query",
				}: params.Match,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = SeriesParams
			Response = *Maps
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackSeriesParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.Series(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.Series(ctx, params)
	}
	if err != nil {
		recordError("Internal", err)
		if errRes, ok := errors.Into[*ErrorStatusCode](err); ok {
			encodeErrorResponse(errRes, w, span)
			return
		}
		if errors.Is(err, ht.ErrNotImplemented) {
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
		encodeErrorResponse(s.h.NewError(ctx, err), w, span)
		return
	}

	if err := encodeSeriesResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}
