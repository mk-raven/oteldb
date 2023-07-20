// Code generated by ogen, DO NOT EDIT.

package ytqueryapi

import (
	"net/http"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// AbortQueryParams is parameters of abortQuery operation.
type AbortQueryParams struct {
	// Query ID to get.
	QueryID QueryID
	// Stage.
	Stage OptString
}

func unpackAbortQueryParams(packed middleware.Parameters) (params AbortQueryParams) {
	{
		key := middleware.ParameterKey{
			Name: "query_id",
			In:   "query",
		}
		params.QueryID = packed[key].(QueryID)
	}
	{
		key := middleware.ParameterKey{
			Name: "stage",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Stage = v.(OptString)
		}
	}
	return params
}

func decodeAbortQueryParams(args [0]string, argsEscaped bool, r *http.Request) (params AbortQueryParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: query_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "query_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotQueryIDVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotQueryIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.QueryID = QueryID(paramsDotQueryIDVal)
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "query_id",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: stage.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "stage",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotStageVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotStageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Stage.SetTo(paramsDotStageVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "stage",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetQueryParams is parameters of getQuery operation.
type GetQueryParams struct {
	// Query ID to get.
	QueryID QueryID
	// Stage.
	Stage OptString
}

func unpackGetQueryParams(packed middleware.Parameters) (params GetQueryParams) {
	{
		key := middleware.ParameterKey{
			Name: "query_id",
			In:   "query",
		}
		params.QueryID = packed[key].(QueryID)
	}
	{
		key := middleware.ParameterKey{
			Name: "stage",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Stage = v.(OptString)
		}
	}
	return params
}

func decodeGetQueryParams(args [0]string, argsEscaped bool, r *http.Request) (params GetQueryParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: query_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "query_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotQueryIDVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotQueryIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.QueryID = QueryID(paramsDotQueryIDVal)
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "query_id",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: stage.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "stage",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotStageVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotStageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Stage.SetTo(paramsDotStageVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "stage",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// ReadQueryResultParams is parameters of readQueryResult operation.
type ReadQueryResultParams struct {
	// Query ID to get result.
	QueryID QueryID
	// Index of a result to read, defaults to 0.
	ResultIndex OptInt
	// Output format.
	OutputFormat OutputFormat
	// Stage.
	Stage OptString
}

func unpackReadQueryResultParams(packed middleware.Parameters) (params ReadQueryResultParams) {
	{
		key := middleware.ParameterKey{
			Name: "query_id",
			In:   "query",
		}
		params.QueryID = packed[key].(QueryID)
	}
	{
		key := middleware.ParameterKey{
			Name: "result_index",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.ResultIndex = v.(OptInt)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "output_format",
			In:   "query",
		}
		params.OutputFormat = packed[key].(OutputFormat)
	}
	{
		key := middleware.ParameterKey{
			Name: "stage",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Stage = v.(OptString)
		}
	}
	return params
}

func decodeReadQueryResultParams(args [0]string, argsEscaped bool, r *http.Request) (params ReadQueryResultParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: query_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "query_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotQueryIDVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotQueryIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.QueryID = QueryID(paramsDotQueryIDVal)
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "query_id",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: result_index.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "result_index",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotResultIndexVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotResultIndexVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ResultIndex.SetTo(paramsDotResultIndexVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "result_index",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: output_format.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "output_format",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.OutputFormat = OutputFormat(c)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if err := params.OutputFormat.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "output_format",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: stage.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "stage",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotStageVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotStageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Stage.SetTo(paramsDotStageVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "stage",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// StartQueryParams is parameters of startQuery operation.
type StartQueryParams struct {
	// Query.
	Query string
	// Query engine to run.
	Engine Engine
	// Stage.
	Stage OptString
}

func unpackStartQueryParams(packed middleware.Parameters) (params StartQueryParams) {
	{
		key := middleware.ParameterKey{
			Name: "query",
			In:   "query",
		}
		params.Query = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "engine",
			In:   "query",
		}
		params.Engine = packed[key].(Engine)
	}
	{
		key := middleware.ParameterKey{
			Name: "stage",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Stage = v.(OptString)
		}
	}
	return params
}

func decodeStartQueryParams(args [0]string, argsEscaped bool, r *http.Request) (params StartQueryParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: query.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "query",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Query = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "query",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: engine.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "engine",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Engine = Engine(c)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if err := params.Engine.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "engine",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: stage.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "stage",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotStageVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotStageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Stage.SetTo(paramsDotStageVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "stage",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}
