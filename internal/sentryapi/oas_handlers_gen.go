// Code generated by ogen, DO NOT EDIT.

package sentryapi

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

// handleDummyRequest handles dummy operation.
//
// Dummy endpoint to generate schema definitions.
//
// GET /dummy
func (s *Server) handleDummyRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("dummy"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/dummy"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "Dummy",
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

	var response *Event
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "Dummy",
			OperationID:   "dummy",
			Body:          nil,
			Params:        middleware.Parameters{},
			Raw:           r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = *Event
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
				response, err = s.h.Dummy(ctx)
				return response, err
			},
		)
	} else {
		response, err = s.h.Dummy(ctx)
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

	if err := encodeDummyResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// handleEnvelopeRequest handles envelope operation.
//
// Sentry push.
//
// POST /envelope
func (s *Server) handleEnvelopeRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("envelope"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/envelope"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "Envelope",
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
			Name: "Envelope",
			ID:   "envelope",
		}
	)
	request, close, err := s.decodeEnvelopeRequest(r)
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

	var response *EnvelopeOK
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "Envelope",
			OperationID:   "envelope",
			Body:          request,
			Params:        middleware.Parameters{},
			Raw:           r,
		}

		type (
			Request  = *EnvelopeReqWithContentType
			Params   = struct{}
			Response = *EnvelopeOK
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
				err = s.h.Envelope(ctx, request)
				return response, err
			},
		)
	} else {
		err = s.h.Envelope(ctx, request)
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

	if err := encodeEnvelopeResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}
