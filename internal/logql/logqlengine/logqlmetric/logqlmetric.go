// Package logqlmetric provides metric queries implementation.
package logqlmetric

import (
	"strconv"
	"time"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"

	"github.com/go-faster/oteldb/internal/iterators"
	"github.com/go-faster/oteldb/internal/lokiapi"
	"github.com/go-faster/oteldb/internal/otelstorage"
)

// Step represents a one query range step i
type Step struct {
	Timestamp otelstorage.Timestamp
	Samples   []Sample
}

// StepIterator is a one query range step iterator.
type StepIterator = iterators.Iterator[Step]

// ReadStepResponse reads aggregation result into API structure.
func ReadStepResponse(iter iterators.Iterator[Step], instant bool) (s lokiapi.QueryResponseData, _ error) {
	var (
		agg          Step
		matrixSeries map[GroupingKey]lokiapi.Series
	)
	for {
		if !iter.Next(&agg) {
			break
		}

		if instant {
			if err := iter.Err(); err != nil {
				return s, err
			}

			var vector lokiapi.Vector
			for _, s := range agg.Samples {
				vector = append(vector, lokiapi.Sample{
					Metric: lokiapi.NewOptLabelSet(s.Set.AsLokiAPI()),
					Value: lokiapi.FPoint{
						T: getPrometheusTimestamp(agg.Timestamp.AsTime()),
						V: strconv.FormatFloat(s.Data, 'f', -1, 64),
					},
				})
			}

			s.SetVectorResult(lokiapi.VectorResult{
				Result: vector,
			})
			return s, nil
		}

		if matrixSeries == nil {
			matrixSeries = map[GroupingKey]lokiapi.Series{}
		}
		for _, s := range agg.Samples {
			key := s.Set.Key()
			ser, ok := matrixSeries[key]
			if !ok {
				ser.Metric.SetTo(s.Set.AsLokiAPI())
			}

			ser.Values = append(ser.Values, lokiapi.FPoint{
				T: getPrometheusTimestamp(agg.Timestamp.AsTime()),
				V: strconv.FormatFloat(s.Data, 'f', -1, 64),
			})
			matrixSeries[key] = ser
		}
	}
	if err := iter.Err(); err != nil {
		return s, err
	}

	// Sort points inside series.
	for k, s := range matrixSeries {
		slices.SortFunc(s.Values, func(a, b lokiapi.FPoint) bool {
			return a.T < b.T
		})
		matrixSeries[k] = s
	}
	result := maps.Values(matrixSeries)
	slices.SortFunc(result, func(a, b lokiapi.Series) bool {
		if len(a.Values) < 1 || len(b.Values) < 1 {
			return len(a.Values) < len(b.Values)
		}
		return a.Values[0].T < b.Values[0].T
	})

	s.SetMatrixResult(lokiapi.MatrixResult{
		Result: result,
	})
	return s, iter.Err()
}

func getPrometheusTimestamp(t time.Time) float64 {
	// Pass milliseconds as fraction part.
	return float64(t.UnixMilli()) / 1000
}
