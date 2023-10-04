// Code generated by ogen, DO NOT EDIT.

package ytqueryapi

import (
	"context"
	"net/url"
	"strings"
	"time"

	"github.com/go-faster/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.19.0"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
)

// Invoker invokes operations described by OpenAPI v3 specification.
type Invoker interface {
	// AbortQuery invokes abortQuery operation.
	//
	// Abort running query.
	//
	// POST /api/v4/abort_query
	AbortQuery(ctx context.Context, params AbortQueryParams) error
	// GetQuery invokes getQuery operation.
	//
	// Gets query status.
	//
	// POST /api/v4/get_query
	GetQuery(ctx context.Context, params GetQueryParams) (*QueryStatus, error)
	// ReadQueryResult invokes readQueryResult operation.
	//
	// Read query result.
	//
	// POST /api/v4/read_query_result
	ReadQueryResult(ctx context.Context, params ReadQueryResultParams) (QueryResult, error)
	// StartQuery invokes startQuery operation.
	//
	// Start a new query.
	//
	// POST /api/v4/start_query
	StartQuery(ctx context.Context, params StartQueryParams) (*StartedQuery, error)
}

// Client implements OAS client.
type Client struct {
	serverURL *url.URL
	sec       SecuritySource
	baseClient
}
type errorHandler interface {
	NewError(ctx context.Context, err error) *ErrorStatusCode
}

var _ Handler = struct {
	errorHandler
	*Client
}{}

func trimTrailingSlashes(u *url.URL) {
	u.Path = strings.TrimRight(u.Path, "/")
	u.RawPath = strings.TrimRight(u.RawPath, "/")
}

// NewClient initializes new Client defined by OAS.
func NewClient(serverURL string, sec SecuritySource, opts ...ClientOption) (*Client, error) {
	u, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}
	trimTrailingSlashes(u)

	c, err := newClientConfig(opts...).baseClient()
	if err != nil {
		return nil, err
	}
	return &Client{
		serverURL:  u,
		sec:        sec,
		baseClient: c,
	}, nil
}

type serverURLKey struct{}

// WithServerURL sets context key to override server URL.
func WithServerURL(ctx context.Context, u *url.URL) context.Context {
	return context.WithValue(ctx, serverURLKey{}, u)
}

func (c *Client) requestURL(ctx context.Context) *url.URL {
	u, ok := ctx.Value(serverURLKey{}).(*url.URL)
	if !ok {
		return c.serverURL
	}
	return u
}

// AbortQuery invokes abortQuery operation.
//
// Abort running query.
//
// POST /api/v4/abort_query
func (c *Client) AbortQuery(ctx context.Context, params AbortQueryParams) error {
	_, err := c.sendAbortQuery(ctx, params)
	return err
}

func (c *Client) sendAbortQuery(ctx context.Context, params AbortQueryParams) (res *AbortedQuery, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("abortQuery"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/api/v4/abort_query"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "AbortQuery",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/api/v4/abort_query"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "query_id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "query_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if unwrapped := string(params.QueryID); true {
				return e.EncodeValue(conv.StringToString(unwrapped))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "stage" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "stage",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Stage.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			stage = "Security:YTToken"
			switch err := c.securityYTToken(ctx, "AbortQuery", r); {
			case err == nil: // if NO error
				satisfied[0] |= 1 << 0
			case errors.Is(err, ogenerrors.ErrSkipClientSecurity):
				// Skip this security.
			default:
				return res, errors.Wrap(err, "security \"YTToken\"")
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
				{},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			return res, ogenerrors.ErrSecurityRequirementIsNotSatisfied
		}
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeAbortQueryResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetQuery invokes getQuery operation.
//
// Gets query status.
//
// POST /api/v4/get_query
func (c *Client) GetQuery(ctx context.Context, params GetQueryParams) (*QueryStatus, error) {
	res, err := c.sendGetQuery(ctx, params)
	return res, err
}

func (c *Client) sendGetQuery(ctx context.Context, params GetQueryParams) (res *QueryStatus, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getQuery"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/api/v4/get_query"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "GetQuery",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/api/v4/get_query"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "query_id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "query_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if unwrapped := string(params.QueryID); true {
				return e.EncodeValue(conv.StringToString(unwrapped))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "stage" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "stage",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Stage.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			stage = "Security:YTToken"
			switch err := c.securityYTToken(ctx, "GetQuery", r); {
			case err == nil: // if NO error
				satisfied[0] |= 1 << 0
			case errors.Is(err, ogenerrors.ErrSkipClientSecurity):
				// Skip this security.
			default:
				return res, errors.Wrap(err, "security \"YTToken\"")
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
				{},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			return res, ogenerrors.ErrSecurityRequirementIsNotSatisfied
		}
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetQueryResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ReadQueryResult invokes readQueryResult operation.
//
// Read query result.
//
// POST /api/v4/read_query_result
func (c *Client) ReadQueryResult(ctx context.Context, params ReadQueryResultParams) (QueryResult, error) {
	res, err := c.sendReadQueryResult(ctx, params)
	return res, err
}

func (c *Client) sendReadQueryResult(ctx context.Context, params ReadQueryResultParams) (res QueryResult, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("readQueryResult"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/api/v4/read_query_result"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "ReadQueryResult",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/api/v4/read_query_result"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "query_id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "query_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if unwrapped := string(params.QueryID); true {
				return e.EncodeValue(conv.StringToString(unwrapped))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "result_index" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "result_index",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.ResultIndex.Get(); ok {
				return e.EncodeValue(conv.IntToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "output_format" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "output_format",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(string(params.OutputFormat)))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "stage" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "stage",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Stage.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			stage = "Security:YTToken"
			switch err := c.securityYTToken(ctx, "ReadQueryResult", r); {
			case err == nil: // if NO error
				satisfied[0] |= 1 << 0
			case errors.Is(err, ogenerrors.ErrSkipClientSecurity):
				// Skip this security.
			default:
				return res, errors.Wrap(err, "security \"YTToken\"")
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
				{},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			return res, ogenerrors.ErrSecurityRequirementIsNotSatisfied
		}
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeReadQueryResultResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// StartQuery invokes startQuery operation.
//
// Start a new query.
//
// POST /api/v4/start_query
func (c *Client) StartQuery(ctx context.Context, params StartQueryParams) (*StartedQuery, error) {
	res, err := c.sendStartQuery(ctx, params)
	return res, err
}

func (c *Client) sendStartQuery(ctx context.Context, params StartQueryParams) (res *StartedQuery, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("startQuery"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/api/v4/start_query"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "StartQuery",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/api/v4/start_query"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "query" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "query",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.Query))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "engine" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "engine",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(string(params.Engine)))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "stage" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "stage",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Stage.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			stage = "Security:YTToken"
			switch err := c.securityYTToken(ctx, "StartQuery", r); {
			case err == nil: // if NO error
				satisfied[0] |= 1 << 0
			case errors.Is(err, ogenerrors.ErrSkipClientSecurity):
				// Skip this security.
			default:
				return res, errors.Wrap(err, "security \"YTToken\"")
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
				{},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			return res, ogenerrors.ErrSecurityRequirementIsNotSatisfied
		}
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeStartQueryResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}
