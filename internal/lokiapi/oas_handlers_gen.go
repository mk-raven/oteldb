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

// handleLabelValuesRequest handles labelValues operation.
//
// Get values of label.
//
// GET /loki/api/v1/label/{name}/values
func (s *Server) handleLabelValuesRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("labelValues"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/loki/api/v1/label/{name}/values"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "LabelValues",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
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
			Name: "LabelValues",
			ID:   "labelValues",
		}
	)
	params, err := decodeLabelValuesParams(args, argsEscaped, r)
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
			OperationName: "LabelValues",
			OperationID:   "labelValues",
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
					Name: "since",
					In:   "query",
				}: params.Since,
				{
					Name: "query",
					In:   "query",
				}: params.Query,
				{
					Name: "name",
					In:   "path",
				}: params.Name,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = LabelValuesParams
			Response = *Values
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackLabelValuesParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.LabelValues(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.LabelValues(ctx, params)
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

	if err := encodeLabelValuesResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// handleLabelsRequest handles labels operation.
//
// Get labels.
// Used by Grafana to test connection to Loki.
//
// GET /loki/api/v1/labels
func (s *Server) handleLabelsRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("labels"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/loki/api/v1/labels"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "Labels",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
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
			Name: "Labels",
			ID:   "labels",
		}
	)
	params, err := decodeLabelsParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response *Labels
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "Labels",
			OperationID:   "labels",
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
					Name: "since",
					In:   "query",
				}: params.Since,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = LabelsParams
			Response = *Labels
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackLabelsParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.Labels(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.Labels(ctx, params)
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

	if err := encodeLabelsResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// handlePushRequest handles push operation.
//
// Push data.
//
// POST /loki/api/v1/push
func (s *Server) handlePushRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("push"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/loki/api/v1/push"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "Push",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
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
			Name: "Push",
			ID:   "push",
		}
	)
	request, close, err := s.decodePushRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *PushNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "Push",
			OperationID:   "push",
			Body:          request,
			Params:        middleware.Parameters{},
			Raw:           r,
		}

		type (
			Request  = PushReq
			Params   = struct{}
			Response = *PushNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.Push(ctx, request)
				return response, err
			},
		)
	} else {
		err = s.h.Push(ctx, request)
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

	if err := encodePushResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// handleQueryRangeRequest handles queryRange operation.
//
// Query range.
//
// GET /loki/api/v1/query_range
func (s *Server) handleQueryRangeRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("queryRange"),
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
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
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
			ID:   "queryRange",
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
			OperationID:   "queryRange",
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

// handleSeriesRequest handles series operation.
//
// Get series.
//
// GET /loki/api/v1/series
func (s *Server) handleSeriesRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("series"),
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
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
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
			ID:   "series",
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
			OperationID:   "series",
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
					Name: "since",
					In:   "query",
				}: params.Since,
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
