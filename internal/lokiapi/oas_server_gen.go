// Code generated by ogen, DO NOT EDIT.

package lokiapi

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// LabelValues implements labelValues operation.
	//
	// Get values of label.
	//
	// GET /loki/api/v1/label/{name}/values
	LabelValues(ctx context.Context, params LabelValuesParams) (*Values, error)
	// Labels implements labels operation.
	//
	// Get labels.
	// Used by Grafana to test connection to Loki.
	//
	// GET /loki/api/v1/labels
	Labels(ctx context.Context, params LabelsParams) (*Labels, error)
	// Push implements push operation.
	//
	// Push data.
	//
	// POST /loki/api/v1/push
	Push(ctx context.Context, req PushReq) error
	// QueryRange implements queryRange operation.
	//
	// Query range.
	//
	// GET /loki/api/v1/query_range
	QueryRange(ctx context.Context, params QueryRangeParams) (*QueryResponse, error)
	// Series implements series operation.
	//
	// Get series.
	//
	// GET /loki/api/v1/series
	Series(ctx context.Context, params SeriesParams) (*Maps, error)
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
