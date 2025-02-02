package prometheusremotewrite

import (
	"fmt"
	"testing"
	"time"

	"github.com/prometheus/prometheus/prompb"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
)

var (
	now       = time.Now()
	nowMillis = now.UnixNano() / int64(time.Millisecond)
)

func TestPrwConfig_FromTimeSeries(t *testing.T) {
	type args struct {
		ts []prompb.TimeSeries
	}
	tests := []struct {
		name     string
		settings Settings
		args     args
		want     pmetric.Metrics
		wantErr  assert.ErrorAssertionFunc
	}{
		{
			name: "sum",
			settings: Settings{
				Logger:        *zap.NewNop(),
				TimeThreshold: 24,
			},
			args: args{
				[]prompb.TimeSeries{
					{
						Labels: []prompb.Label{
							{
								Name:  nameStr,
								Value: value71,
							},
							{
								Name:  label12,
								Value: value12,
							},
						},
						Samples: []prompb.Sample{
							{Value: 1.0, Timestamp: nowMillis},
							{Value: 2.0, Timestamp: nowMillis},
						},
					},
				},
			},
			want: getMetrics(getSumMetric(
				value71, "", true, getAttributes(label12, value12),
				numberPoint{1., uint64(time.Now().UnixNano())},
				numberPoint{2., uint64(time.Now().UnixNano())},
			)),
			wantErr: assert.NoError,
		},
		{
			name: "count",
			settings: Settings{
				Logger:        *zap.NewNop(),
				TimeThreshold: 24,
			},
			args: args{
				[]prompb.TimeSeries{
					{
						Labels: []prompb.Label{
							{
								Name:  nameStr,
								Value: value61,
							},
							{
								Name:  label12,
								Value: value12,
							},
						},
						Samples: []prompb.Sample{
							{Value: 2.0, Timestamp: nowMillis},
							{Value: 3.0, Timestamp: nowMillis},
						},
					},
				},
			},
			want: getMetrics(getSumMetric(
				value61, "", true, getAttributes(label12, value12),
				numberPoint{2., uint64(time.Now().UnixNano())},
				numberPoint{3., uint64(time.Now().UnixNano())},
			)),
			wantErr: assert.NoError,
		},
		{
			name: "bytes",
			settings: Settings{
				Logger:        *zap.NewNop(),
				TimeThreshold: 24,
			},
			args: args{
				[]prompb.TimeSeries{
					{
						Labels: []prompb.Label{
							{
								Name:  nameStr,
								Value: value81,
							},
							{
								Name:  label12,
								Value: value12,
							},
						},
						Samples: []prompb.Sample{{Value: 2.0, Timestamp: nowMillis}},
					},
				},
			},
			want:    getMetrics(getDoubleGaugeMetric(value81, "bytes", getAttributes(label12, value12), 2., uint64(time.Now().UnixNano()))),
			wantErr: assert.NoError,
		},
		{
			name: "bytes_total",
			settings: Settings{
				Logger:        *zap.NewNop(),
				TimeThreshold: 24,
			},
			args: args{
				[]prompb.TimeSeries{
					{
						Labels: []prompb.Label{
							{
								Name:  nameStr,
								Value: value91,
							},
							{
								Name:  label12,
								Value: value12,
							},
						},
						Samples: []prompb.Sample{{Value: 2.0, Timestamp: nowMillis}},
					},
				},
			},
			want: getMetrics(getSumMetric(
				value91, "bytes", true, getAttributes(label12, value12),
				numberPoint{2., uint64(time.Now().UnixNano())},
			)),
			wantErr: assert.NoError,
		},
		{
			name: "count - old",
			settings: Settings{
				Logger:        *zap.NewNop(),
				TimeThreshold: 24,
			},
			args: args{
				[]prompb.TimeSeries{
					{
						Labels: []prompb.Label{
							{
								Name:  nameStr,
								Value: value61,
							},
							{
								Name:  label12,
								Value: value12,
							},
						},
						Samples: []prompb.Sample{{Value: 0.0, Timestamp: now.Add(-time.Hour*24).UnixNano() / int64(time.Millisecond)}},
					},
				},
			},
			want:    getMetrics(getNoneMetric(value61)),
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMetrics, err := FromTimeSeries(tt.args.ts, tt.settings)
			for i := 0; i < gotMetrics.ResourceMetrics().Len(); i++ {
				rm := gotMetrics.ResourceMetrics().At(i)
				for j := 0; j < rm.ScopeMetrics().Len(); j++ {
					sm := rm.ScopeMetrics().At(j)
					for k := 0; k < sm.Metrics().Len(); k++ {
						gotMetric := sm.Metrics().At(k)
						testMetric := tt.want.ResourceMetrics().At(i).ScopeMetrics().At(j).Metrics().At(k)
						if !tt.wantErr(t, err, fmt.Sprintf("FromTimeSeries(%v)", tt.args.ts)) {
							return
						}
						if (gotMetric.Sum() != pmetric.Sum{}) {
							for i := 0; i < gotMetric.Sum().DataPoints().Len(); i++ {
								gotMetric.Sum().DataPoints().At(i).SetTimestamp(pcommon.NewTimestampFromTime(testMetric.Sum().DataPoints().At(i).Timestamp().AsTime()))
							}
						}
						if (gotMetric.Summary() != pmetric.Summary{}) {
							for i := 0; i < gotMetric.Summary().DataPoints().Len(); i++ {
								gotMetric.Summary().DataPoints().At(i).SetTimestamp(pcommon.NewTimestampFromTime(testMetric.Summary().DataPoints().At(i).Timestamp().AsTime()))
							}
						}
						if (gotMetric.Gauge() != pmetric.Gauge{}) {
							for i := 0; i < gotMetric.Gauge().DataPoints().Len(); i++ {
								gotMetric.Gauge().DataPoints().At(i).SetTimestamp(pcommon.NewTimestampFromTime(testMetric.Gauge().DataPoints().At(i).Timestamp().AsTime()))
							}
						}
						if (gotMetric.Histogram() != pmetric.Histogram{}) {
							for i := 0; i < gotMetric.Histogram().DataPoints().Len(); i++ {
								gotMetric.Histogram().DataPoints().At(i).SetTimestamp(pcommon.NewTimestampFromTime(testMetric.Histogram().DataPoints().At(i).Timestamp().AsTime()))
							}
						}
						assert.Equalf(t, testMetric, gotMetric, "FromTimeSeries(%v)", tt.args.ts)
					}
				}
			}
			assert.Equalf(t, tt.want, gotMetrics, "FromTimeSeries(%v)", tt.args.ts)
		})
	}
}

func getMetrics(metric pmetric.Metric) pmetric.Metrics {
	metrics := pmetric.NewMetrics()
	empty := metrics.ResourceMetrics().AppendEmpty().ScopeMetrics().AppendEmpty().Metrics().AppendEmpty()
	pm := pmetric.NewMetric()
	pm.MoveTo(empty)
	metric.MoveTo(empty)
	return metrics
}

func Test_finalName(t *testing.T) {
	type args struct {
		labels []prompb.Label
	}
	tests := []struct {
		name    string
		args    args
		wantRet string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "test if __name__ label is set",
			args: args{
				labels: []prompb.Label{
					{
						Name:  nameStr,
						Value: "foo",
					},
				},
			},
			wantRet: "foo",
			wantErr: assert.NoError,
		},
		{
			name: "test if __name__ label is not set",
			args: args{
				labels: []prompb.Label{
					{
						Name:  "foo",
						Value: "bar",
					},
				},
			},
			wantRet: "",
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRet, err := finalName(tt.args.labels)
			if !tt.wantErr(t, err, fmt.Sprintf("finalName(%v)", tt.args.labels)) {
				return
			}
			assert.Equalf(t, tt.wantRet, gotRet, "finalName(%v)", tt.args.labels)
		})
	}
}
