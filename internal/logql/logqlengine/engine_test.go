// Package logqlengine implements LogQL evaluation engine.
package logqlengine

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"golang.org/x/exp/slices"

	"github.com/go-faster/oteldb/internal/iterators"
	"github.com/go-faster/oteldb/internal/logql"
	"github.com/go-faster/oteldb/internal/logstorage"
	"github.com/go-faster/oteldb/internal/lokiapi"
	"github.com/go-faster/oteldb/internal/otelstorage"
)

type inputLine struct {
	line  string
	attrs pcommon.Map
}

type resultLine struct {
	line   string
	labels map[string]string
}

type mockQuerier struct {
	lines []inputLine
	step  time.Duration
}

func (m *mockQuerier) Сapabilities() (caps QuerierСapabilities) {
	return caps
}

func (m *mockQuerier) SelectLogs(_ context.Context, start, _ otelstorage.Timestamp, _ SelectLogsParams) (iterators.Iterator[logstorage.Record], error) {
	step := m.step
	if step == 0 {
		step = time.Millisecond
	}
	ts := start.AsTime()

	var (
		records    []logstorage.Record
		scopeAttrs = pcommon.NewMap()
		resAttrs   = pcommon.NewMap()
	)
	scopeAttrs.PutStr("scope", "test")
	resAttrs.PutStr("resource", "test")

	for _, l := range m.lines {
		ts = ts.Add(step)
		records = append(records, logstorage.Record{
			Timestamp:     otelstorage.NewTimestampFromTime(ts),
			Body:          l.line,
			Attrs:         otelstorage.Attrs(l.attrs),
			ScopeAttrs:    otelstorage.Attrs(scopeAttrs),
			ResourceAttrs: otelstorage.Attrs(resAttrs),
		})
	}

	return iterators.Slice(records), nil
}

func TestEngineEvalStream(t *testing.T) {
	justLines := func(lines ...string) []inputLine {
		r := make([]inputLine, len(lines))
		for i, line := range lines {
			r[i] = inputLine{
				line: line,
			}
		}
		return r
	}

	startTime := otelstorage.Timestamp(1688833731000000000)
	inputLines := justLines(
		`{"id": 1, "foo": "4m", "bar": "1s", "baz": "1kb"}`,
		`{"id": 2, "foo": "5m", "bar": "2s", "baz": "1mb"}`,
		`{"id": 3, "foo": "6m", "bar": "3s", "baz": "1gb"}`,
	)
	resultLines := []resultLine{
		{
			`{"id": 1, "foo": "4m", "bar": "1s", "baz": "1kb"}`,
			map[string]string{
				"id":  "1",
				"foo": "4m",
				"bar": "1s",
				"baz": "1kb",
			},
		},
		{
			`{"id": 2, "foo": "5m", "bar": "2s", "baz": "1mb"}`,
			map[string]string{
				"id":  "2",
				"foo": "5m",
				"bar": "2s",
				"baz": "1mb",
			},
		},
		{
			`{"id": 3, "foo": "6m", "bar": "3s", "baz": "1gb"}`,
			map[string]string{
				"id":  "3",
				"foo": "6m",
				"bar": "3s",
				"baz": "1gb",
			},
		},
	}

	tests := []struct {
		query    string
		input    []inputLine
		wantData []resultLine
		wantErr  bool
	}{
		// Line filter.
		{
			`{resource="test"} |= "5m" | json`,
			inputLines,
			[]resultLine{resultLines[1]},
			false,
		},
		{
			`{resource="test"} |~ ".*5m.*" | json`,
			inputLines,
			[]resultLine{resultLines[1]},
			false,
		},
		{
			`{resource="test"} != "5m" | json`,
			inputLines,
			[]resultLine{
				resultLines[0],
				resultLines[2],
			},
			false,
		},
		{
			`{resource="test"} !~ "5m" | json`,
			inputLines,
			[]resultLine{
				resultLines[0],
				resultLines[2],
			},
			false,
		},
		// Label matcher.
		{
			`{resource="test"} | json | foo = "5"`,
			inputLines,
			[]resultLine{},
			false,
		},
		{
			`{resource="test"} | json | foo = "5m"`,
			inputLines,
			[]resultLine{resultLines[1]},
			false,
		},
		{
			`{resource="test"} | json | foo =~ ".*5m.*"`,
			inputLines,
			[]resultLine{resultLines[1]},
			false,
		},
		{
			`{resource="test"} | json | foo != "5m"`,
			inputLines,
			[]resultLine{
				resultLines[0],
				resultLines[2],
			},
			false,
		},
		{
			`{resource="test"} | json | foo !~ "5m"`,
			inputLines,
			[]resultLine{
				resultLines[0],
				resultLines[2],
			},
			false,
		},

		// Drop expression.
		{
			`{resource="test"} | json | drop id`,
			justLines(
				`{"id": 1, "foo": "bar"}`,
				`{"id": 2, "foo": "baz"}`,
			),
			[]resultLine{
				{
					`{"id": 1, "foo": "bar"}`,
					map[string]string{"foo": "bar"},
				},
				{
					`{"id": 2, "foo": "baz"}`,
					map[string]string{"foo": "baz"},
				},
			},
			false,
		},
		{
			`{resource="test"} | json | drop clearly_not_exists`,
			justLines(
				`{"id": 1, "foo": "bar"}`,
			),
			[]resultLine{
				{
					`{"id": 1, "foo": "bar"}`,
					map[string]string{
						"id":  "1",
						"foo": "bar",
					},
				},
			},
			false,
		},

		// Distinct filter.
		{
			`{resource="test"} | json | distinct id`,
			justLines(
				`{"id": 1, "foo": "bar"}`,
				`{"id": 1, "foo": "baz"}`,
				`{"id": 2, "foo": "bar"}`,
			),
			[]resultLine{
				{
					`{"id": 1, "foo": "bar"}`,
					map[string]string{
						"id":  "1",
						"foo": "bar",
					},
				},
				{
					`{"id": 2, "foo": "bar"}`,
					map[string]string{
						"id":  "2",
						"foo": "bar",
					},
				},
			},
			false,
		},
		{
			`{resource="test"} | json | distinct id, foo`,
			justLines(
				`{"id": 1, "foo": "bar"}`,
				`{"id": 1, "foo": "baz"}`,
				`{"id": 2, "foo": "bar"}`,
			),
			[]resultLine{
				{
					`{"id": 1, "foo": "bar"}`,
					map[string]string{
						"id":  "1",
						"foo": "bar",
					},
				},
			},
			false,
		},
		{
			`{resource="test"} | json | distinct clearly_not_exists`,
			inputLines,
			resultLines,
			false,
		},

		// Complex queries.
		{
			`{resource="test"} | json | foo <= 5m, bar == 2s or baz == 1MB`,
			inputLines,
			[]resultLine{resultLines[1]},
			false,
		},
	}
	for i, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("Test%d", i+1), func(t *testing.T) {
			ctx := context.Background()

			opts := Options{
				ParseOptions: logql.ParseOptions{AllowDots: true},
			}
			e := NewEngine(&mockQuerier{lines: tt.input}, opts)

			gotData, err := e.Eval(ctx, tt.query, EvalParams{
				Start: startTime,
				End:   startTime,
				Limit: 1000,
			})
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			result, ok := gotData.GetStreamsResult()
			require.True(t, ok)

			type entry struct {
				ts     uint64
				line   string
				labels map[string]string
			}
			entries := make([]entry, 0, len(tt.wantData))
			for _, s := range result.Result {
				set := s.Stream.Value

				// Ensure engine passes resource and scope attrs.
				for k, v := range map[string]string{
					"scope":    "test",
					"resource": "test",
				} {
					assert.Contains(t, set, k)
					assert.Equal(t, v, set[k])
					// Do add labels to the test.
					delete(set, k)
				}

				for _, e := range s.Values {
					entries = append(entries, entry{
						ts:     e.T,
						line:   e.V,
						labels: set,
					})
				}
			}
			slices.SortFunc(entries, func(a, b entry) bool { return a.ts < b.ts })

			assert.Len(t, entries, len(tt.wantData))
			for i, e := range entries {
				result := resultLine{
					line:   e.line,
					labels: e.labels,
				}
				assert.Equal(t, result, tt.wantData[i])
			}
		})
	}
}

func TestEngineEvalLiteral(t *testing.T) {
	type timeRange struct {
		start uint64
		end   uint64
		step  time.Duration
	}
	tests := []struct {
		query   string
		tsRange timeRange

		wantData lokiapi.QueryResponseData
		wantErr  bool
	}{
		// Literal eval.
		{
			`3.14`,
			timeRange{
				start: 1700000001_000000000,
				end:   1700000001_000000000,
			},
			lokiapi.QueryResponseData{
				Type: lokiapi.ScalarResultQueryResponseData,
				ScalarResult: lokiapi.ScalarResult{
					Result: lokiapi.FPoint{
						T: 1700000001,
						V: "3.14",
					},
				},
			},
			false,
		},
		{
			`3.14`,
			timeRange{
				start: 1700000001_000000000,
				end:   1700000003_000000000,
				step:  time.Second,
			},
			lokiapi.QueryResponseData{
				Type: lokiapi.MatrixResultQueryResponseData,
				MatrixResult: lokiapi.MatrixResult{
					Result: lokiapi.Matrix{
						{
							Values: []lokiapi.FPoint{
								{T: 1700000001, V: "3.14"},
								{T: 1700000002, V: "3.14"},
								{T: 1700000003, V: "3.14"},
							},
						},
					},
				},
			},
			false,
		},

		// Vector eval.
		{
			`vector(3.14)`,
			timeRange{
				start: 1700000001_000000000,
				end:   1700000001_000000000,
			},
			lokiapi.QueryResponseData{
				Type: lokiapi.VectorResultQueryResponseData,
				VectorResult: lokiapi.VectorResult{
					Result: lokiapi.Vector{
						{
							Value: lokiapi.FPoint{
								T: 1700000001,
								V: "3.14",
							},
						},
					},
				},
			},
			false,
		},
		{
			`vector(3.14)`,
			timeRange{
				start: 1700000001_000000000,
				end:   1700000003_000000000,
				step:  time.Second,
			},
			lokiapi.QueryResponseData{
				Type: lokiapi.MatrixResultQueryResponseData,
				MatrixResult: lokiapi.MatrixResult{
					Result: lokiapi.Matrix{
						{
							Values: []lokiapi.FPoint{
								{T: 1700000001, V: "3.14"},
								{T: 1700000002, V: "3.14"},
								{T: 1700000003, V: "3.14"},
							},
						},
					},
				},
			},
			false,
		},

		// Binop reduce.
		{
			`2*2`,
			timeRange{
				start: 1700000001_000000000,
				end:   1700000001_000000000,
			},
			lokiapi.QueryResponseData{
				Type: lokiapi.ScalarResultQueryResponseData,
				ScalarResult: lokiapi.ScalarResult{
					Result: lokiapi.FPoint{
						T: 1700000001,
						V: "4",
					},
				},
			},
			false,
		},
	}
	for i, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("Test%d", i+1), func(t *testing.T) {
			ctx := context.Background()

			opts := Options{
				ParseOptions: logql.ParseOptions{AllowDots: true},
			}
			e := NewEngine(&mockQuerier{}, opts)

			gotData, err := e.Eval(ctx, tt.query, EvalParams{
				Start: otelstorage.Timestamp(tt.tsRange.start),
				End:   otelstorage.Timestamp(tt.tsRange.end),
				Step:  tt.tsRange.step,
				Limit: 1000,
			})
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.wantData, gotData)
		})
	}
}
