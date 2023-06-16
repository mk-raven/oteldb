// Code generated by ogen, DO NOT EDIT.

package promapi

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// GetLabelValues implements getLabelValues operation.
	//
	// GET /api/v1/label/{label}/values
	GetLabelValues(ctx context.Context, params GetLabelValuesParams) (*LabelValuesResponse, error)
	// GetLabels implements getLabels operation.
	//
	// GET /api/v1/labels
	GetLabels(ctx context.Context, params GetLabelsParams) (*LabelsResponse, error)
	// GetMetadata implements getMetadata operation.
	//
	// GET /api/v1/metadata
	GetMetadata(ctx context.Context, params GetMetadataParams) (*MetadataResponse, error)
	// GetQuery implements getQuery operation.
	//
	// Query Prometheus.
	//
	// GET /api/v1/query
	GetQuery(ctx context.Context, params GetQueryParams) (*QueryResponse, error)
	// GetQueryExemplars implements getQueryExemplars operation.
	//
	// Query Prometheus.
	//
	// GET /api/v1/query_exemplars
	GetQueryExemplars(ctx context.Context, params GetQueryExemplarsParams) (*QueryExemplarsResponse, error)
	// GetQueryRange implements getQueryRange operation.
	//
	// Query Prometheus.
	//
	// GET /api/v1/query_range
	GetQueryRange(ctx context.Context, params GetQueryRangeParams) (*QueryResponse, error)
	// GetRules implements getRules operation.
	//
	// GET /api/v1/rules
	GetRules(ctx context.Context, params GetRulesParams) (*RulesResponse, error)
	// GetSeries implements getSeries operation.
	//
	// Query Prometheus.
	//
	// GET /api/v1/series
	GetSeries(ctx context.Context, params GetSeriesParams) (*SeriesResponse, error)
	// PostLabels implements postLabels operation.
	//
	// POST /api/v1/labels
	PostLabels(ctx context.Context) (*LabelsResponse, error)
	// PostQuery implements postQuery operation.
	//
	// Query Prometheus.
	//
	// POST /api/v1/query
	PostQuery(ctx context.Context) (*QueryResponse, error)
	// PostQueryExemplars implements postQueryExemplars operation.
	//
	// Query Prometheus.
	//
	// POST /api/v1/query_exemplars
	PostQueryExemplars(ctx context.Context) (*QueryExemplarsResponse, error)
	// PostQueryRange implements postQueryRange operation.
	//
	// Query Prometheus.
	//
	// POST /api/v1/query_range
	PostQueryRange(ctx context.Context) (*QueryResponse, error)
	// PostSeries implements postSeries operation.
	//
	// Query Prometheus.
	//
	// POST /api/v1/series
	PostSeries(ctx context.Context) (*SeriesResponse, error)
	// NewError creates *FailStatusCode from error returned by handler.
	//
	// Used for common default response.
	NewError(ctx context.Context, err error) *FailStatusCode
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
