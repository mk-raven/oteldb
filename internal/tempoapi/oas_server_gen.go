// Code generated by ogen, DO NOT EDIT.

package tempoapi

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// Search implements search operation.
	//
	// Execute TraceQL query.
	//
	// GET /api/search
	Search(ctx context.Context, params SearchParams) (*Traces, error)
	// SearchTagValues implements search_tag_values operation.
	//
	// This endpoint retrieves all discovered values for the given tag, which can be used in search.
	//
	// GET /api/search/tag/{tag_name}/values
	SearchTagValues(ctx context.Context, params SearchTagValuesParams) (*TagValues, error)
	// SearchTagValuesV2 implements search_tag_values_v2 operation.
	//
	// This endpoint retrieves all discovered values and their data types for the given TraceQL
	// identifier.
	//
	// GET /api/v2/search/tag/{tag_name}/values
	SearchTagValuesV2(ctx context.Context, params SearchTagValuesV2Params) (*TagValuesV2, error)
	// SearchTags implements search_tags operation.
	//
	// This endpoint retrieves all discovered tag names that can be used in search.
	//
	// GET /api/search/tags
	SearchTags(ctx context.Context) (*TagNames, error)
	// TraceByID implements traceByID operation.
	//
	// Querying traces by id.
	//
	// GET /api/traces/{traceID}
	TraceByID(ctx context.Context, params TraceByIDParams) (TraceByIDRes, error)
	// NewError creates *ErrorStatusCode from error returned by handler.
	//
	// Used for common default response.
	NewError(ctx context.Context, err error) *ErrorStatusCode
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h Handler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		baseServer: s,
	}, nil
}
