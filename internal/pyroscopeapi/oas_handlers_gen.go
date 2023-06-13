// Code generated by ogen, DO NOT EDIT.

package pyroscopeapi

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

// handleGetAppsRequest handles getApps operation.
//
// Returns list of application metadata.
// Used by Grafana to test connection to Pyroscope.
//
// GET /api/apps
func (s *Server) handleGetAppsRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getApps"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/api/apps"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "GetApps",
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
		err error
	)

	var response []ApplicationMetadata
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "GetApps",
			OperationID:   "getApps",
			Body:          nil,
			Params:        middleware.Parameters{},
			Raw:           r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = []ApplicationMetadata
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
				response, err = s.h.GetApps(ctx)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetApps(ctx)
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

	if err := encodeGetAppsResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// handleLabelValuesRequest handles labelValues operation.
//
// Returns list of label values.
//
// GET /label-values
func (s *Server) handleLabelValuesRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("labelValues"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/label-values"),
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

	var response LabelValues
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "LabelValues",
			OperationID:   "labelValues",
			Body:          nil,
			Params: middleware.Parameters{
				{
					Name: "label",
					In:   "query",
				}: params.Label,
				{
					Name: "from",
					In:   "query",
				}: params.From,
				{
					Name: "until",
					In:   "query",
				}: params.Until,
				{
					Name: "query",
					In:   "query",
				}: params.Query,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = LabelValuesParams
			Response = LabelValues
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
// Returns list of labels.
//
// GET /labels
func (s *Server) handleLabelsRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("labels"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/labels"),
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

	var response Labels
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "Labels",
			OperationID:   "labels",
			Body:          nil,
			Params: middleware.Parameters{
				{
					Name: "from",
					In:   "query",
				}: params.From,
				{
					Name: "until",
					In:   "query",
				}: params.Until,
				{
					Name: "query",
					In:   "query",
				}: params.Query,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = LabelsParams
			Response = Labels
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

// handleRenderRequest handles render operation.
//
// Renders given query.
// One of `query` or `key` is required.
//
// GET /render
func (s *Server) handleRenderRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("render"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/render"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "Render",
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
			Name: "Render",
			ID:   "render",
		}
	)
	params, err := decodeRenderParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response *FlamebearerProfileV1
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "Render",
			OperationID:   "render",
			Body:          nil,
			Params: middleware.Parameters{
				{
					Name: "from",
					In:   "query",
				}: params.From,
				{
					Name: "until",
					In:   "query",
				}: params.Until,
				{
					Name: "query",
					In:   "query",
				}: params.Query,
				{
					Name: "name",
					In:   "query",
				}: params.Name,
				{
					Name: "groupBy",
					In:   "query",
				}: params.GroupBy,
				{
					Name: "maxNodes",
					In:   "query",
				}: params.MaxNodes,
				{
					Name: "format",
					In:   "query",
				}: params.Format,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = RenderParams
			Response = *FlamebearerProfileV1
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackRenderParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.Render(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.Render(ctx, params)
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

	if err := encodeRenderResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}
