// Code generated by ogen, DO NOT EDIT.

package tempoapi

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// Echo implements echo operation.
//
// Echo request for testing, issued by Grafana.
//
// GET /api/echo
func (UnimplementedHandler) Echo(ctx context.Context) (r EchoOK, _ error) {
	return r, ht.ErrNotImplemented
}

// Search implements search operation.
//
// Execute TraceQL query.
//
// GET /api/search
func (UnimplementedHandler) Search(ctx context.Context, params SearchParams) (r *Traces, _ error) {
	return r, ht.ErrNotImplemented
}

// SearchTagValues implements searchTagValues operation.
//
// This endpoint retrieves all discovered values for the given tag, which can be used in search.
//
// GET /api/search/tag/{tag_name}/values
func (UnimplementedHandler) SearchTagValues(ctx context.Context, params SearchTagValuesParams) (r *TagValues, _ error) {
	return r, ht.ErrNotImplemented
}

// SearchTagValuesV2 implements searchTagValuesV2 operation.
//
// This endpoint retrieves all discovered values and their data types for the given TraceQL
// identifier.
//
// GET /api/v2/search/tag/{tag_name}/values
func (UnimplementedHandler) SearchTagValuesV2(ctx context.Context, params SearchTagValuesV2Params) (r *TagValuesV2, _ error) {
	return r, ht.ErrNotImplemented
}

// SearchTags implements searchTags operation.
//
// This endpoint retrieves all discovered tag names that can be used in search.
//
// GET /api/search/tags
func (UnimplementedHandler) SearchTags(ctx context.Context) (r *TagNames, _ error) {
	return r, ht.ErrNotImplemented
}

// TraceByID implements traceByID operation.
//
// Querying traces by id.
//
// GET /api/traces/{traceID}
func (UnimplementedHandler) TraceByID(ctx context.Context, params TraceByIDParams) (r TraceByIDRes, _ error) {
	return r, ht.ErrNotImplemented
}

// NewError creates *ErrorStatusCode from error returned by handler.
//
// Used for common default response.
func (UnimplementedHandler) NewError(ctx context.Context, err error) (r *ErrorStatusCode) {
	r = new(ErrorStatusCode)
	return r
}
