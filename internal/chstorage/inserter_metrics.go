package chstorage

import (
	"context"
	"strconv"
	"time"

	"github.com/ClickHouse/ch-go"
	"github.com/ClickHouse/ch-go/chpool"
	"github.com/ClickHouse/ch-go/proto"
	"github.com/go-faster/errors"
	"github.com/go-faster/sdk/zctx"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"golang.org/x/sync/errgroup"
)

// ConsumeMetrics inserts given metrics.
func (i *Inserter) ConsumeMetrics(ctx context.Context, metrics pmetric.Metrics) error {
	b := newMetricBatch()
	if err := i.mapMetrics(b, metrics); err != nil {
		return errors.Wrap(err, "map metrics")
	}
	if err := b.Insert(ctx, i.tables, i.ch); err != nil {
		return errors.Wrap(err, "send batch")
	}
	return nil
}

type metricsBatch struct {
	points        *pointColumns
	histograms    *histogramColumns
	expHistograms *expHistogramColumns
	summaries     *summaryColumns
	labels        map[[2]string]struct{}
}

func newMetricBatch() *metricsBatch {
	return &metricsBatch{
		points:        newPointColumns(),
		histograms:    newHistogramColumns(),
		expHistograms: newExpHistogramColumns(),
		summaries:     newSummaryColumns(),
		labels:        map[[2]string]struct{}{},
	}
}

func (b *metricsBatch) Insert(ctx context.Context, tables Tables, client *chpool.Pool) error {
	labels := newLabelsColumns()
	for pair := range b.labels {
		labels.name.Append(pair[0])
		labels.value.Append(pair[1])
	}

	grp, grpCtx := errgroup.WithContext(ctx)
	type columns interface {
		Input() proto.Input
	}
	for _, table := range []struct {
		name    string
		columns columns
	}{
		{tables.Points, b.points},
		{tables.Histograms, b.histograms},
		{tables.ExpHistograms, b.expHistograms},
		{tables.Summaries, b.summaries},
		{tables.Labels, labels},
	} {
		table := table
		grp.Go(func() error {
			ctx := grpCtx

			input := table.columns.Input()
			if err := client.Do(ctx, ch.Query{
				Logger: zctx.From(ctx),
				Body:   input.Into(table.name),
				Input:  input,
			}); err != nil {
				return errors.Wrapf(err, "insert %q", table.name)
			}
			return nil
		})
	}
	if err := grp.Wait(); err != nil {
		return errors.Wrap(err, "insert")
	}

	return nil
}

func (b *metricsBatch) addPoints(name string, res pcommon.Map, slice pmetric.NumberDataPointSlice) error {
	c := b.points

	for i := 0; i < slice.Len(); i++ {
		point := slice.At(i)
		ts := point.Timestamp().AsTime()
		flags := point.Flags()
		attrs := point.Attributes()

		var val float64
		switch typ := point.ValueType(); typ {
		case pmetric.NumberDataPointValueTypeEmpty:
			// Just ignore it.
			continue
		case pmetric.NumberDataPointValueTypeInt:
			// TODO(tdakkota): check for overflow
			val = float64(point.IntValue())
		case pmetric.NumberDataPointValueTypeDouble:
			val = point.DoubleValue()
		default:
			return errors.Errorf("unexpected metric %q value type: %v", name, typ)
		}

		b.addLabels(attrs)
		c.name.Append(name)
		c.timestamp.Append(ts)
		c.mapping.Append(proto.Enum8(noMapping))
		c.value.Append(val)
		c.flags.Append(uint32(flags))
		c.attributes.Append(encodeAttributes(attrs))
		c.resource.Append(encodeAttributes(res))
	}
	return nil
}

func (b *metricsBatch) addHistogramPoints(name string, res pcommon.Map, slice pmetric.HistogramDataPointSlice) error {
	c := b.histograms
	for i := 0; i < slice.Len(); i++ {
		point := slice.At(i)
		ts := point.Timestamp().AsTime()
		flags := point.Flags()
		attrs := point.Attributes()
		count := point.Count()
		sum := proto.Nullable[float64]{
			Set:   point.HasSum(),
			Value: point.Sum(),
		}
		_min := proto.Nullable[float64]{
			Set:   point.HasMin(),
			Value: point.Min(),
		}
		_max := proto.Nullable[float64]{
			Set:   point.HasMax(),
			Value: point.Max(),
		}
		bucketCounts := point.BucketCounts().AsRaw()
		explicitBounds := point.ExplicitBounds().AsRaw()

		b.addLabels(attrs)
		// Save original histogram.
		c.name.Append(name)
		c.timestamp.Append(ts)
		c.count.Append(count)
		c.sum.Append(sum)
		c.min.Append(_min)
		c.max.Append(_max)
		c.bucketCounts.Append(bucketCounts)
		c.explicitBounds.Append(explicitBounds)
		c.flags.Append(uint32(flags))
		c.attributes.Append(encodeAttributes(attrs))
		c.resource.Append(encodeAttributes(res))

		// Map histogram as set of series for Prometheus compatibility.
		series := mappedSeries{
			ts:    ts,
			flags: flags,
			attrs: attrs,
			res:   res,
		}
		if sum.Set {
			b.addMappedSample(
				series,
				name+"_sum",
				histogramSum,
				sum.Value,
				[2]string{},
			)
		}
		if _min.Set {
			b.addMappedSample(
				series,
				name+"_min",
				histogramMin,
				_min.Value,
				[2]string{},
			)
		}
		if _max.Set {
			b.addMappedSample(
				series,
				name+"_max",
				histogramMax,
				_max.Value,
				[2]string{},
			)
		}
		b.addMappedSample(
			series,
			name+"_count",
			histogramCount,
			float64(count),
			[2]string{},
		)

		var (
			cumCount   uint64
			bucketName = name + "_bucket"
		)
		for i := 0; i < min(len(bucketCounts), len(explicitBounds)); i++ {
			bound := explicitBounds[i]
			cumCount += bucketCounts[i]

			// Generate series with "_bucket" suffix and "le" label.
			b.addMappedSample(
				series,
				bucketName,
				histogramBucket,
				float64(cumCount),
				[2]string{
					"le",
					strconv.FormatFloat(bound, 'f', -1, 64),
				},
			)
		}
		// Generate series with "_bucket" suffix and "le" label.
		b.addMappedSample(
			series,
			bucketName,
			histogramBucket,
			float64(cumCount),
			[2]string{
				"le",
				"+Inf",
			},
		)
	}
	return nil
}

func (b *metricsBatch) addExpHistogramPoints(name string, res pcommon.Map, slice pmetric.ExponentialHistogramDataPointSlice) error {
	var (
		c          = b.expHistograms
		mapBuckets = func(b pmetric.ExponentialHistogramDataPointBuckets) (offset int32, counts []uint64) {
			offset = b.Offset()
			counts = b.BucketCounts().AsRaw()
			return offset, counts
		}
	)
	for i := 0; i < slice.Len(); i++ {
		point := slice.At(i)
		ts := point.Timestamp().AsTime()
		flags := point.Flags()
		attrs := point.Attributes()
		count := point.Count()
		sum := proto.Nullable[float64]{
			Set:   point.HasSum(),
			Value: point.Sum(),
		}
		min := proto.Nullable[float64]{
			Set:   point.HasMin(),
			Value: point.Min(),
		}
		max := proto.Nullable[float64]{
			Set:   point.HasMax(),
			Value: point.Max(),
		}
		scale := point.Scale()
		zerocount := point.ZeroCount()

		positiveOffset, positiveBucketCounts := mapBuckets(point.Positive())
		negativeOffset, negativeBucketCounts := mapBuckets(point.Negative())

		b.addLabels(attrs)
		c.name.Append(name)
		c.timestamp.Append(ts)
		c.count.Append(count)
		c.sum.Append(sum)
		c.min.Append(min)
		c.max.Append(max)
		c.scale.Append(scale)
		c.zerocount.Append(zerocount)
		c.positiveOffset.Append(positiveOffset)
		c.positiveBucketCounts.Append(positiveBucketCounts)
		c.negativeOffset.Append(negativeOffset)
		c.negativeBucketCounts.Append(negativeBucketCounts)
		c.flags.Append(uint32(flags))
		c.attributes.Append(encodeAttributes(attrs))
		c.resource.Append(encodeAttributes(res))
	}
	return nil
}

func (b *metricsBatch) addSummaryPoints(name string, res pcommon.Map, slice pmetric.SummaryDataPointSlice) error {
	c := b.summaries
	for i := 0; i < slice.Len(); i++ {
		point := slice.At(i)
		ts := point.Timestamp().AsTime()
		flags := point.Flags()
		attrs := point.Attributes()
		count := point.Count()
		sum := point.Sum()
		var (
			qv = point.QuantileValues()

			quantiles = make([]float64, qv.Len())
			values    = make([]float64, qv.Len())
		)
		for i := 0; i < qv.Len(); i++ {
			p := qv.At(i)

			quantiles[i] = p.Quantile()
			values[i] = p.Value()
		}

		b.addLabels(attrs)
		c.name.Append(name)
		c.timestamp.Append(ts)
		c.count.Append(count)
		c.sum.Append(sum)
		c.quantiles.Append(quantiles)
		c.values.Append(values)
		c.flags.Append(uint32(flags))
		c.attributes.Append(encodeAttributes(attrs))
		c.resource.Append(encodeAttributes(res))

		series := mappedSeries{
			ts:    ts,
			flags: flags,
			attrs: attrs,
			res:   res,
		}
		b.addMappedSample(series, name+"_count", summaryCount, float64(count), [2]string{})
		b.addMappedSample(series, name+"_sum", summarySum, sum, [2]string{})

		for i := 0; i < min(len(quantiles), len(values)); i++ {
			quantile := quantiles[i]
			value := values[i]

			// Generate series with "quantile" label.
			b.addMappedSample(series, name, summaryQuantile, value, [2]string{
				"quantile",
				strconv.FormatFloat(quantile, 'f', -1, 64),
			})
		}
	}
	return nil
}

type mappedSeries struct {
	ts         time.Time
	flags      pmetric.DataPointFlags
	attrs, res pcommon.Map
}

func (b *metricsBatch) addMappedSample(
	series mappedSeries,
	name string,
	mapping metricMapping,
	val float64,
	bucketKey [2]string,
) {
	c := b.points
	c.name.Append(name)
	c.timestamp.Append(series.ts)
	c.mapping.Append(proto.Enum8(mapping))
	c.value.Append(val)
	c.flags.Append(uint32(series.flags))
	c.attributes.Append(encodeAttributes(series.attrs, bucketKey))
	c.resource.Append(encodeAttributes(series.res, bucketKey))
}

func (b *metricsBatch) addLabels(m pcommon.Map) {
	m.Range(func(name string, value pcommon.Value) bool {
		key := [2]string{
			name,
			// FIXME(tdakkota): annoying allocations
			value.AsString(),
		}
		b.labels[key] = struct{}{}
		return true
	})
}

func (i *Inserter) mapMetrics(b *metricsBatch, metrics pmetric.Metrics) error {
	resMetrics := metrics.ResourceMetrics()
	for i := 0; i < resMetrics.Len(); i++ {
		resMetric := resMetrics.At(i)
		resAttrs := resMetric.Resource().Attributes()
		b.addLabels(resAttrs)

		scopeMetrics := resMetric.ScopeMetrics()
		for i := 0; i < scopeMetrics.Len(); i++ {
			scopeLog := scopeMetrics.At(i)

			records := scopeLog.Metrics()
			for i := 0; i < records.Len(); i++ {
				record := records.At(i)
				name := record.Name()

				switch typ := record.Type(); typ {
				case pmetric.MetricTypeGauge:
					gauge := record.Gauge()
					if err := b.addPoints(name, resAttrs, gauge.DataPoints()); err != nil {
						return err
					}
				case pmetric.MetricTypeSum:
					sum := record.Sum()
					if err := b.addPoints(name, resAttrs, sum.DataPoints()); err != nil {
						return err
					}
				case pmetric.MetricTypeHistogram:
					hist := record.Histogram()
					if err := b.addHistogramPoints(name, resAttrs, hist.DataPoints()); err != nil {
						return err
					}
				case pmetric.MetricTypeExponentialHistogram:
					hist := record.ExponentialHistogram()
					if err := b.addExpHistogramPoints(name, resAttrs, hist.DataPoints()); err != nil {
						return err
					}
				case pmetric.MetricTypeSummary:
					summary := record.Summary()
					if err := b.addSummaryPoints(name, resAttrs, summary.DataPoints()); err != nil {
						return err
					}
				default:
					return errors.Errorf("unexpected metric %q type %v", name, typ)
				}
			}
		}
	}

	return nil
}