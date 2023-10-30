// Code generated by ogen, DO NOT EDIT.

package promapi

import (
	"net/http"
	"net/url"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// GetLabelValuesParams is parameters of getLabelValues operation.
type GetLabelValuesParams struct {
	// Label to query values.
	Label string
	// Start timestamp.
	Start OptPrometheusTimestamp
	// End timestamp.
	End OptPrometheusTimestamp
	// Repeated series selector argument that selects the series from which to read the label names.
	Match []string
}

func unpackGetLabelValuesParams(packed middleware.Parameters) (params GetLabelValuesParams) {
	{
		key := middleware.ParameterKey{
			Name: "label",
			In:   "path",
		}
		params.Label = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "start",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Start = v.(OptPrometheusTimestamp)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "end",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.End = v.(OptPrometheusTimestamp)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "match[]",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Match = v.([]string)
		}
	}
	return params
}

func decodeGetLabelValuesParams(args [1]string, argsEscaped bool, r *http.Request) (params GetLabelValuesParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: label.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "label",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Label = c
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
			Name: "label",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: start.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "start",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotStartVal PrometheusTimestamp
				if err := func() error {
					var paramsDotStartValVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						paramsDotStartValVal = c
						return nil
					}(); err != nil {
						return err
					}
					paramsDotStartVal = PrometheusTimestamp(paramsDotStartValVal)
					return nil
				}(); err != nil {
					return err
				}
				params.Start.SetTo(paramsDotStartVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "start",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: end.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "end",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotEndVal PrometheusTimestamp
				if err := func() error {
					var paramsDotEndValVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						paramsDotEndValVal = c
						return nil
					}(); err != nil {
						return err
					}
					paramsDotEndVal = PrometheusTimestamp(paramsDotEndValVal)
					return nil
				}(); err != nil {
					return err
				}
				params.End.SetTo(paramsDotEndVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "end",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: match[].
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "match[]",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				return d.DecodeArray(func(d uri.Decoder) error {
					var paramsDotMatchVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						paramsDotMatchVal = c
						return nil
					}(); err != nil {
						return err
					}
					params.Match = append(params.Match, paramsDotMatchVal)
					return nil
				})
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "match[]",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetLabelsParams is parameters of getLabels operation.
type GetLabelsParams struct {
	// Start timestamp.
	Start OptPrometheusTimestamp
	// End timestamp.
	End OptPrometheusTimestamp
	// Repeated series selector argument that selects the series from which to read the label names.
	Match []string
}

func unpackGetLabelsParams(packed middleware.Parameters) (params GetLabelsParams) {
	{
		key := middleware.ParameterKey{
			Name: "start",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Start = v.(OptPrometheusTimestamp)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "end",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.End = v.(OptPrometheusTimestamp)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "match[]",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Match = v.([]string)
		}
	}
	return params
}

func decodeGetLabelsParams(args [0]string, argsEscaped bool, r *http.Request) (params GetLabelsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: start.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "start",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotStartVal PrometheusTimestamp
				if err := func() error {
					var paramsDotStartValVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						paramsDotStartValVal = c
						return nil
					}(); err != nil {
						return err
					}
					paramsDotStartVal = PrometheusTimestamp(paramsDotStartValVal)
					return nil
				}(); err != nil {
					return err
				}
				params.Start.SetTo(paramsDotStartVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "start",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: end.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "end",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotEndVal PrometheusTimestamp
				if err := func() error {
					var paramsDotEndValVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						paramsDotEndValVal = c
						return nil
					}(); err != nil {
						return err
					}
					paramsDotEndVal = PrometheusTimestamp(paramsDotEndValVal)
					return nil
				}(); err != nil {
					return err
				}
				params.End.SetTo(paramsDotEndVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "end",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: match[].
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "match[]",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				return d.DecodeArray(func(d uri.Decoder) error {
					var paramsDotMatchVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						paramsDotMatchVal = c
						return nil
					}(); err != nil {
						return err
					}
					params.Match = append(params.Match, paramsDotMatchVal)
					return nil
				})
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "match[]",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetMetadataParams is parameters of getMetadata operation.
type GetMetadataParams struct {
	// Maximum number of metrics to return.
	Limit OptInt
	// FIXME(tdakkota): undocumented.
	LimitPerMetric OptInt
	// A metric name to filter metadata for.
	// All metric metadata is retrieved if left empty.
	Metric OptString
}

func unpackGetMetadataParams(packed middleware.Parameters) (params GetMetadataParams) {
	{
		key := middleware.ParameterKey{
			Name: "limit",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Limit = v.(OptInt)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "limit_per_metric",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.LimitPerMetric = v.(OptInt)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "metric",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Metric = v.(OptString)
		}
	}
	return params
}

func decodeGetMetadataParams(args [0]string, argsEscaped bool, r *http.Request) (params GetMetadataParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: limit.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "limit",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotLimitVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotLimitVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Limit.SetTo(paramsDotLimitVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "limit",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: limit_per_metric.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "limit_per_metric",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotLimitPerMetricVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotLimitPerMetricVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.LimitPerMetric.SetTo(paramsDotLimitPerMetricVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "limit_per_metric",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: metric.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "metric",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotMetricVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotMetricVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Metric.SetTo(paramsDotMetricVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "metric",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetQueryParams is parameters of getQuery operation.
type GetQueryParams struct {
	// Prometheus expression query string.
	Query string
	// Evaluation timestamp.
	Time OptPrometheusTimestamp
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
			params.Time = v.(OptPrometheusTimestamp)
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
				var paramsDotTimeVal PrometheusTimestamp
				if err := func() error {
					var paramsDotTimeValVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						paramsDotTimeValVal = c
						return nil
					}(); err != nil {
						return err
					}
					paramsDotTimeVal = PrometheusTimestamp(paramsDotTimeValVal)
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

// GetQueryExemplarsParams is parameters of getQueryExemplars operation.
type GetQueryExemplarsParams struct {
	// Prometheus expression query string.
	Query string
	// Start timestamp.
	Start PrometheusTimestamp
	// End timestamp.
	End PrometheusTimestamp
}

func unpackGetQueryExemplarsParams(packed middleware.Parameters) (params GetQueryExemplarsParams) {
	{
		key := middleware.ParameterKey{
			Name: "query",
			In:   "query",
		}
		params.Query = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "start",
			In:   "query",
		}
		params.Start = packed[key].(PrometheusTimestamp)
	}
	{
		key := middleware.ParameterKey{
			Name: "end",
			In:   "query",
		}
		params.End = packed[key].(PrometheusTimestamp)
	}
	return params
}

func decodeGetQueryExemplarsParams(args [0]string, argsEscaped bool, r *http.Request) (params GetQueryExemplarsParams, _ error) {
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
	// Decode query: start.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "start",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotStartVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotStartVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Start = PrometheusTimestamp(paramsDotStartVal)
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
			Name: "start",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: end.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "end",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotEndVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotEndVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.End = PrometheusTimestamp(paramsDotEndVal)
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
			Name: "end",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetQueryRangeParams is parameters of getQueryRange operation.
type GetQueryRangeParams struct {
	// Prometheus expression query string.
	Query string
	// Start timestamp, inclusive.
	Start PrometheusTimestamp
	// End timestamp, inclusive.
	End PrometheusTimestamp
	// Query resolution step width in duration format or float number of seconds.
	Step string
}

func unpackGetQueryRangeParams(packed middleware.Parameters) (params GetQueryRangeParams) {
	{
		key := middleware.ParameterKey{
			Name: "query",
			In:   "query",
		}
		params.Query = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "start",
			In:   "query",
		}
		params.Start = packed[key].(PrometheusTimestamp)
	}
	{
		key := middleware.ParameterKey{
			Name: "end",
			In:   "query",
		}
		params.End = packed[key].(PrometheusTimestamp)
	}
	{
		key := middleware.ParameterKey{
			Name: "step",
			In:   "query",
		}
		params.Step = packed[key].(string)
	}
	return params
}

func decodeGetQueryRangeParams(args [0]string, argsEscaped bool, r *http.Request) (params GetQueryRangeParams, _ error) {
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
	// Decode query: start.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "start",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotStartVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotStartVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Start = PrometheusTimestamp(paramsDotStartVal)
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
			Name: "start",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: end.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "end",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotEndVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotEndVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.End = PrometheusTimestamp(paramsDotEndVal)
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
			Name: "end",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: step.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "step",
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

				params.Step = c
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
			Name: "step",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetRulesParams is parameters of getRules operation.
type GetRulesParams struct {
	// Return only the alerting rules (e.g. type=alert) or the recording rules (e.g. type=record).
	// When the parameter is absent or empty, no filtering is done.
	Type OptGetRulesType
	// Only return rules with the given rule name.
	// If the parameter is repeated, rules with any of the provided names are returned.
	// If we've filtered out all the rules of a group, the group is not returned.
	// When the parameter is absent or empty, no filtering is done.
	RuleName []string
	// Only return rules with the given rule group name.
	// If the parameter is repeated, rules with any of the provided rule group names are returned.
	// When the parameter is absent or empty, no filtering is done.
	RuleGroup []string
	// Only return rules with the given filepath.
	// If the parameter is repeated, rules with any of the provided filepaths are returned.
	// When the parameter is absent or empty, no filtering is done.
	File []string
}

func unpackGetRulesParams(packed middleware.Parameters) (params GetRulesParams) {
	{
		key := middleware.ParameterKey{
			Name: "type",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Type = v.(OptGetRulesType)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "rule_name[]",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.RuleName = v.([]string)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "rule_group[]",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.RuleGroup = v.([]string)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "file[]",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.File = v.([]string)
		}
	}
	return params
}

func decodeGetRulesParams(args [0]string, argsEscaped bool, r *http.Request) (params GetRulesParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: type.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "type",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotTypeVal GetRulesType
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotTypeVal = GetRulesType(c)
					return nil
				}(); err != nil {
					return err
				}
				params.Type.SetTo(paramsDotTypeVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if value, ok := params.Type.Get(); ok {
					if err := func() error {
						if err := value.Validate(); err != nil {
							return err
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "type",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: rule_name[].
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "rule_name[]",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				return d.DecodeArray(func(d uri.Decoder) error {
					var paramsDotRuleNameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						paramsDotRuleNameVal = c
						return nil
					}(); err != nil {
						return err
					}
					params.RuleName = append(params.RuleName, paramsDotRuleNameVal)
					return nil
				})
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "rule_name[]",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: rule_group[].
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "rule_group[]",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				return d.DecodeArray(func(d uri.Decoder) error {
					var paramsDotRuleGroupVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						paramsDotRuleGroupVal = c
						return nil
					}(); err != nil {
						return err
					}
					params.RuleGroup = append(params.RuleGroup, paramsDotRuleGroupVal)
					return nil
				})
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "rule_group[]",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: file[].
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "file[]",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				return d.DecodeArray(func(d uri.Decoder) error {
					var paramsDotFileVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						paramsDotFileVal = c
						return nil
					}(); err != nil {
						return err
					}
					params.File = append(params.File, paramsDotFileVal)
					return nil
				})
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "file[]",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetSeriesParams is parameters of getSeries operation.
type GetSeriesParams struct {
	// Start timestamp.
	Start PrometheusTimestamp
	// End timestamp.
	End PrometheusTimestamp
	// Repeated series selector argument that selects the series from which to read the label names.
	Match []string
}

func unpackGetSeriesParams(packed middleware.Parameters) (params GetSeriesParams) {
	{
		key := middleware.ParameterKey{
			Name: "start",
			In:   "query",
		}
		params.Start = packed[key].(PrometheusTimestamp)
	}
	{
		key := middleware.ParameterKey{
			Name: "end",
			In:   "query",
		}
		params.End = packed[key].(PrometheusTimestamp)
	}
	{
		key := middleware.ParameterKey{
			Name: "match[]",
			In:   "query",
		}
		params.Match = packed[key].([]string)
	}
	return params
}

func decodeGetSeriesParams(args [0]string, argsEscaped bool, r *http.Request) (params GetSeriesParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: start.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "start",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotStartVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotStartVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Start = PrometheusTimestamp(paramsDotStartVal)
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
			Name: "start",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: end.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "end",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotEndVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotEndVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.End = PrometheusTimestamp(paramsDotEndVal)
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
			Name: "end",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: match[].
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "match[]",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				return d.DecodeArray(func(d uri.Decoder) error {
					var paramsDotMatchVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						paramsDotMatchVal = c
						return nil
					}(); err != nil {
						return err
					}
					params.Match = append(params.Match, paramsDotMatchVal)
					return nil
				})
			}); err != nil {
				return err
			}
			if err := func() error {
				if params.Match == nil {
					return errors.New("nil is invalid value")
				}
				if err := (validate.Array{
					MinLength:    1,
					MinLengthSet: true,
					MaxLength:    0,
					MaxLengthSet: false,
				}).ValidateLength(len(params.Match)); err != nil {
					return errors.Wrap(err, "array")
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
			Name: "match[]",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}
