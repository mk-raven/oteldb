// Code generated by ogen, DO NOT EDIT.

package promapi

import (
	"net/http"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// GetQueryParams is parameters of getQuery operation.
type GetQueryParams struct {
	Query string
	// Evaluation timestamp.
	Time OptString
}

func unpackGetQueryParams(packed middleware.Parameters) (params GetQueryParams) {
	{
		key := middleware.ParameterKey{
			Name: "query",
			In:   "query",
		}
		params.Query = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "time",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Time = v.(OptString)
		}
	}
	return params
}

func decodeGetQueryParams(args [0]string, argsEscaped bool, r *http.Request) (params GetQueryParams, _ error) {
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
	// Decode query: time.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "time",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotTimeVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotTimeVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Time.SetTo(paramsDotTimeVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "time",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}