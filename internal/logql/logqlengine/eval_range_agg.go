package logqlengine

import (
	"context"
	"strconv"
	"time"

	"github.com/go-faster/errors"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"

	"github.com/go-faster/oteldb/internal/iterators"
	"github.com/go-faster/oteldb/internal/logql"
	"github.com/go-faster/oteldb/internal/lokiapi"
	"github.com/go-faster/oteldb/internal/otelstorage"
)

// fpoint is a metric fpoint.
type fpoint struct {
	Timestamp otelstorage.Timestamp
	Value     float64
}

// MilliT returns Prometheus millisecond timestamp.
func (p fpoint) MilliT() int64 {
	return p.Timestamp.AsTime().UnixMilli()
}

type sample struct {
	data float64
	set  lokiapi.LabelSet
	key  string
}

type series struct {
	data []fpoint
	set  lokiapi.LabelSet
	key  string
}

type rangeAggIterator struct {
	iter iterators.Iterator[sampledEntry]

	agg aggregator
	// step state
	current time.Time
	end     time.Time
	step    time.Duration

	grouper grouper
	// window state
	window   map[string]series
	interval time.Duration
	sampled  sampledEntry
	// buffered whether last entry is buffered
	buffered bool
}

func newRangeAggIterator(
	iter iterators.Iterator[sampledEntry],
	expr *logql.RangeAggregationExpr,
	start, end time.Time,
	step time.Duration,
) (*rangeAggIterator, error) {
	if step == 0 {
		step = time.Second
	}
	aggtr, err := buildAggregator(expr)
	if err != nil {
		return nil, errors.Wrap(err, "build aggregator")
	}

	return &rangeAggIterator{
		iter: iter,

		agg:     aggtr,
		current: start.Add(-step),
		end:     end,
		step:    step,

		grouper:  buildGrouper(expr.Grouping),
		window:   map[string]series{},
		interval: expr.Range.Range,
	}, nil
}

type rangeAgg struct {
	ts      otelstorage.Timestamp
	samples []sample
}

func (i *rangeAggIterator) Next(r *rangeAgg) bool {
	i.current = i.current.Add(i.step)
	if i.current.After(i.end) {
		return false
	}

	// Fill the window.
	windowStart := i.current.Add(-i.interval)
	windowEnd := i.current
	i.fillWindow(windowStart, windowEnd)

	// Aggregate the window.
	r.ts = otelstorage.NewTimestampFromTime(i.current)
	r.samples = r.samples[:0]
	for _, s := range i.window {
		r.samples = append(r.samples, sample{
			data: i.agg.Aggregate(s.data),
			set:  s.set,
			key:  s.key,
		})
	}

	return true
}

func (i *rangeAggIterator) clearWindow(windowStart time.Time) {
	for key, s := range i.window {
		// Filter series data in place: timestamp should be >= windowStart.
		n := 0
		for _, p := range s.data {
			if p.Timestamp.AsTime().Before(windowStart) {
				continue
			}
			s.data[n] = p
			n++
		}
		s.data = s.data[:n]

		if len(s.data) < 1 {
			// Delete empty series.
			delete(i.window, key)
		} else {
			i.window[key] = s
		}
	}
}

func (i *rangeAggIterator) fillWindow(windowStart, windowEnd time.Time) {
	i.clearWindow(windowStart)

	for {
		if !i.buffered {
			if !i.iter.Next(&i.sampled) {
				return
			}
		} else {
			// Do not read next entry, use buffered
			i.buffered = false
		}

		s := i.sampled
		e := s.entry

		switch ts := e.ts.AsTime(); {
		case ts.After(windowEnd):
			// Entry is after the end of current window: buffer for the next window.
			i.buffered = true
			return
		case ts.Before(windowStart):
			// Entry is before the start of current window: just skip it.
			continue
		}

		groupKey, metric := i.grouper.Group(e)

		ser, ok := i.window[groupKey]
		if !ok {
			ser.set = metric
			ser.key = groupKey
		}
		ser.data = append(ser.data, fpoint{
			Timestamp: e.ts,
			Value:     s.sample,
		})
		i.window[groupKey] = ser
	}
}

func (i *rangeAggIterator) Err() error {
	return i.iter.Err()
}

func (i *rangeAggIterator) Close() error {
	return i.iter.Close()
}

func (e *Engine) rangeAggIterator(ctx context.Context, expr *logql.RangeAggregationExpr, params EvalParams) (_ *rangeAggIterator, rerr error) {
	qrange := expr.Range
	if o := qrange.Offset; o != nil {
		params.Start = addDuration(params.Start, -o.Duration)
		params.End = addDuration(params.End, -o.Duration)
	}

	entries, err := e.selectLogs(ctx, qrange.Sel, qrange.Pipeline, params)
	if err != nil {
		return nil, errors.Wrap(err, "select logs")
	}
	defer func() {
		if rerr != nil {
			_ = entries.Close()
		}
	}()

	samples, err := newSampleIterator(entries, expr)
	if err != nil {
		return nil, errors.Wrap(err, "build sample iterator")
	}

	var (
		start = params.Start.AsTime()
		end   = params.End.AsTime()
		step  = params.Step
	)
	return newRangeAggIterator(samples, expr, start, end, step)
}

func (e *Engine) evalRangeAggregation(ctx context.Context, expr *logql.RangeAggregationExpr, params EvalParams) (s lokiapi.QueryResponseData, _ error) {
	iter, err := e.rangeAggIterator(ctx, expr, params)
	if err != nil {
		return s, err
	}
	defer func() {
		_ = iter.Close()
	}()

	return readRangeAggregation(iter, params.IsInstant())
}

func readRangeAggregation(iter iterators.Iterator[rangeAgg], instant bool) (s lokiapi.QueryResponseData, _ error) {
	var (
		agg          rangeAgg
		matrixSeries map[string]lokiapi.Series
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
			for _, s := range agg.samples {
				vector = append(vector, lokiapi.Sample{
					Metric: lokiapi.NewOptLabelSet(s.set),
					Value: lokiapi.FPoint{
						T: getPrometheusTimestamp(agg.ts.AsTime()),
						V: strconv.FormatFloat(s.data, 'f', -1, 64),
					},
				})
			}

			s.SetVectorResult(lokiapi.VectorResult{
				Result: vector,
			})
			return s, nil
		}

		if matrixSeries == nil {
			matrixSeries = map[string]lokiapi.Series{}
		}
		for _, s := range agg.samples {
			ser, ok := matrixSeries[s.key]
			if !ok {
				ser.Metric.SetTo(s.set)
			}

			ser.Values = append(ser.Values, lokiapi.FPoint{
				T: getPrometheusTimestamp(agg.ts.AsTime()),
				V: strconv.FormatFloat(s.data, 'f', -1, 64),
			})
			matrixSeries[s.key] = ser
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
