package logql

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func ptrTo[T any](v T) *T {
	return &v
}

func TestParse(t *testing.T) {
	tests := []struct {
		input   string
		want    Expr
		wantErr bool
	}{
		{`{}`, &LogExpr{}, false},
		{`({})`, &ParenExpr{X: &LogExpr{}}, false},
		{
			`{foo="bar"}`,
			&LogExpr{
				Sel: Selector{
					Matchers: []LabelMatcher{
						{"foo", OpEq, "bar"},
					},
				},
			},
			false,
		},
		{
			`{foo = "bar"}`,
			&LogExpr{
				Sel: Selector{
					Matchers: []LabelMatcher{
						{"foo", OpEq, "bar"},
					},
				},
			},
			false,
		},
		{
			`{foo!="bar"}`,
			&LogExpr{
				Sel: Selector{
					Matchers: []LabelMatcher{
						{"foo", OpNotEq, "bar"},
					},
				},
			},
			false,
		},
		{
			`{foo != "bar"}`,
			&LogExpr{
				Sel: Selector{
					Matchers: []LabelMatcher{
						{"foo", OpNotEq, "bar"},
					},
				},
			},
			false,
		},
		{
			`{foo =~ "bar"}`,
			&LogExpr{
				Sel: Selector{
					Matchers: []LabelMatcher{
						{"foo", OpRe, "bar"},
					},
				},
			},
			false,
		},
		{
			`{foo !~ "bar"}`,
			&LogExpr{
				Sel: Selector{
					Matchers: []LabelMatcher{
						{"foo", OpNotRe, "bar"},
					},
				},
			},
			false,
		},
		{
			`{foo !~ "bar", foo2 =~ "amongus"}`,
			&LogExpr{
				Sel: Selector{
					Matchers: []LabelMatcher{
						{"foo", OpNotRe, "bar"},
						{"foo2", OpRe, "amongus"},
					},
				},
			},
			false,
		},
		{
			`( {foo = "bar"} )`,
			&ParenExpr{
				X: &LogExpr{
					Sel: Selector{
						Matchers: []LabelMatcher{
							{"foo", OpEq, "bar"},
						},
					},
				},
			},
			false,
		},
		{
			"{} |= `foo`",
			&LogExpr{
				Pipeline: []PipelineStage{
					&LineFilter{Op: OpEq, Value: "foo"},
				},
			},
			false,
		},
		{
			`{instance=~"kafka-1",name="kafka"}
				|= "bad"
				|~ "error"
				!= "good"
				!~ "exception"`,
			&LogExpr{
				Sel: Selector{
					Matchers: []LabelMatcher{
						{"instance", OpRe, "kafka-1"},
						{"name", OpEq, "kafka"},
					},
				},
				Pipeline: []PipelineStage{
					&LineFilter{Op: OpEq, Value: "bad"},
					&LineFilter{Op: OpRe, Value: "error"},
					&LineFilter{Op: OpNotEq, Value: "good"},
					&LineFilter{Op: OpNotRe, Value: "exception"},
				},
			},
			false,
		},
		{
			`( {instance=~"kafka-1",name="kafka"} |= "bad" )`,
			&ParenExpr{
				X: &LogExpr{
					Sel: Selector{
						Matchers: []LabelMatcher{
							{"instance", OpRe, "kafka-1"},
							{"name", OpEq, "kafka"},
						},
					},
					Pipeline: []PipelineStage{
						&LineFilter{Op: OpEq, Value: "bad"},
					},
				},
			},
			false,
		},
		{
			`{name="kafka"}
				|= "bad"
				| logfmt
				| json
				| regexp ".*"
				| pattern "<ip>"
				| unpack
				| line_format "{{ . }}"
				| decolorize`,
			&LogExpr{
				Sel: Selector{
					Matchers: []LabelMatcher{
						{"name", OpEq, "kafka"},
					},
				},
				Pipeline: []PipelineStage{
					&LineFilter{Op: OpEq, Value: "bad"},
					&LogfmtExpressionParser{},
					&JSONExpressionParser{},
					&RegexpLabelParser{Regexp: ".*"},
					&PatternLabelParser{Pattern: "<ip>"},
					&UnpackLabelParser{},
					&LineFormat{Template: "{{ . }}"},
					&DecolorizeExpr{},
				},
			},
			false,
		},
		{
			`{name="kafka"}
				|= "bad"
				| json
				| json foo, bar
				| json foo="10", bar="sus"
				| logfmt foo="10", bar="sus"
			`,
			&LogExpr{
				Sel: Selector{
					Matchers: []LabelMatcher{
						{"name", OpEq, "kafka"},
					},
				},
				Pipeline: []PipelineStage{
					&LineFilter{Op: OpEq, Value: "bad"},
					&JSONExpressionParser{},
					&JSONExpressionParser{
						Labels: []Label{
							"foo",
							"bar",
						},
					},
					&JSONExpressionParser{
						Exprs: []LabelExtractionExpr{
							{"foo", "10"},
							{"bar", "sus"},
						},
					},
					&LogfmtExpressionParser{
						Exprs: []LabelExtractionExpr{
							{"foo", "10"},
							{"bar", "sus"},
						},
					},
				},
			},
			false,
		},
		{
			`{name="kafka"}
				| drop foo
				| drop foo, foo2
				| keep foo=~"bar"
				| keep foo=~"bar", foo2=~"baz"
				| drop foo,foo2=~"bar",foo3
				| keep foo!~"bar",foo2,foo3
			`,
			&LogExpr{
				Sel: Selector{
					Matchers: []LabelMatcher{
						{"name", OpEq, "kafka"},
					},
				},
				Pipeline: []PipelineStage{
					&DropLabelsExpr{
						Labels: []Label{"foo"},
					},
					&DropLabelsExpr{
						Labels: []Label{"foo", "foo2"},
					},
					&KeepLabelsExpr{
						Matchers: []LabelMatcher{
							{"foo", OpRe, "bar"},
						},
					},
					&KeepLabelsExpr{
						Matchers: []LabelMatcher{
							{"foo", OpRe, "bar"},
							{"foo2", OpRe, "baz"},
						},
					},
					&DropLabelsExpr{
						Labels: []Label{"foo", "foo3"},
						Matchers: []LabelMatcher{
							{"foo2", OpRe, "bar"},
						},
					},
					&KeepLabelsExpr{
						Labels: []Label{"foo2", "foo3"},
						Matchers: []LabelMatcher{
							{"foo", OpNotRe, "bar"},
						},
					},
				},
			},
			false,
		},
		{
			`{name="kafka"}
				| label_format foo=foo
				| label_format bar="bar"
				| label_format foo=foo,bar="bar"
			`,
			&LogExpr{
				Sel: Selector{
					Matchers: []LabelMatcher{
						{"name", OpEq, "kafka"},
					},
				},
				Pipeline: []PipelineStage{
					&LabelFormatExpr{
						Labels: []LabelFormatLabel{
							{"foo", "foo"},
						},
					},
					&LabelFormatExpr{
						Values: []LabelFormatValue{
							{"bar", "bar"},
						},
					},
					&LabelFormatExpr{
						Labels: []LabelFormatLabel{
							{"foo", "foo"},
						},
						Values: []LabelFormatValue{
							{"bar", "bar"},
						},
					},
				},
			},
			false,
		},
		{
			`{name="kafka"}
				| distinct foo
				| distinct foo,bar
				| distinct foo,bar,baz
			`,
			&LogExpr{
				Sel: Selector{
					Matchers: []LabelMatcher{
						{"name", OpEq, "kafka"},
					},
				},
				Pipeline: []PipelineStage{
					&DistinctFilter{
						Labels: []Label{"foo"},
					},
					&DistinctFilter{
						Labels: []Label{"foo", "bar"},
					},

					&DistinctFilter{
						Labels: []Label{"foo", "bar", "baz"},
					},
				},
			},
			false,
		},
		{
			`{instance=~"kafka-1",name="kafka"}
				| status == 200
				| (service = "sus1")
				| service = "sus2", request != "GET"
				| service = "sus3" request != "POST"
				| ( (service = "sus4") and request != "PUT" )`,
			&LogExpr{
				Sel: Selector{
					Matchers: []LabelMatcher{
						{"instance", OpRe, "kafka-1"},
						{"name", OpEq, "kafka"},
					},
				},
				Pipeline: []PipelineStage{
					&LabelFilter{
						Pred: &NumberFilter{"status", OpEq, 200},
					},
					&LabelFilter{
						Pred: &LabelPredicateParen{
							X: &LabelMatcher{"service", OpEq, "sus1"},
						},
					},
					&LabelFilter{
						Pred: &LabelPredicateBinOp{
							Left:  &LabelMatcher{"service", OpEq, "sus2"},
							Op:    OpAnd,
							Right: &LabelMatcher{"request", OpNotEq, "GET"},
						},
					},
					&LabelFilter{
						Pred: &LabelPredicateBinOp{
							Left:  &LabelMatcher{"service", OpEq, "sus3"},
							Op:    OpAnd,
							Right: &LabelMatcher{"request", OpNotEq, "POST"},
						},
					},
					&LabelFilter{
						Pred: &LabelPredicateParen{
							X: &LabelPredicateBinOp{
								Left: &LabelPredicateParen{
									X: &LabelMatcher{"service", OpEq, "sus4"},
								},
								Op:    OpAnd,
								Right: &LabelMatcher{"request", OpNotEq, "PUT"},
							},
						},
					},
				},
			},
			false,
		},
		{
			`{instance=~"kafka-1",name="kafka"}
				| duration >= 20ms or size == 20kb and method!~"2.."
				| ip == ip("127.0.0.1")`,
			&LogExpr{
				Sel: Selector{
					Matchers: []LabelMatcher{
						{"instance", OpRe, "kafka-1"},
						{"name", OpEq, "kafka"},
					},
				},
				Pipeline: []PipelineStage{
					&LabelFilter{
						Pred: &LabelPredicateBinOp{
							Left: &DurationFilter{"duration", OpGte, 20 * time.Millisecond},
							Op:   OpOr,
							Right: &LabelPredicateBinOp{
								Left:  &BytesFilter{"size", OpEq, 20 * 1000}, // 20kb
								Op:    OpAnd,
								Right: &LabelMatcher{"method", OpNotRe, "2.."},
							},
						},
					},
					&LabelFilter{
						Pred: &IPFilter{"ip", OpEq, "127.0.0.1"},
					},
				},
			},
			false,
		},

		// Metric queries.
		// Range aggregation.
		{
			`count_over_time( ({job="mysql"})[5m] offset 15m)`,
			&RangeAggregationExpr{
				Op: RangeOpCount,
				Range: LogRangeExpr{
					Sel: Selector{
						Matchers: []LabelMatcher{
							{"job", OpEq, "mysql"},
						},
					},
					Range: 5 * time.Minute,
					Offset: &OffsetExpr{
						Duration: 15 * time.Minute,
					},
				},
			},
			false,
		},
		{
			`avg_over_time({ job = "mysql" }[5h]) without (foo)`,
			&RangeAggregationExpr{
				Op: RangeOpAvg,
				Range: LogRangeExpr{
					Sel: Selector{
						Matchers: []LabelMatcher{
							{"job", OpEq, "mysql"},
						},
					},
					Range: 5 * time.Hour,
				},
				Grouping: &Grouping{
					Op:     GroupingOpWithout,
					Labels: []Label{"foo"},
				},
			},
			false,
		},
		{
			`avg_over_time(10, { job = "mysql" }[5h] |= "error" | logfmt) by (bar,foo)`,
			&RangeAggregationExpr{
				Op: RangeOpAvg,
				Range: LogRangeExpr{
					Sel: Selector{
						Matchers: []LabelMatcher{
							{"job", OpEq, "mysql"},
						},
					},
					Pipeline: []PipelineStage{
						&LineFilter{Op: OpEq, Value: "error"},
						&LogfmtExpressionParser{},
					},
					Range: 5 * time.Hour,
				},
				Parameter: ptrTo(10.0),
				Grouping: &Grouping{
					Op:     GroupingOpBy,
					Labels: []Label{"bar", "foo"},
				},
			},
			false,
		},
		{
			`avg_over_time({ job = "mysql" }[5h] |= "error" | unwrap duration)`,
			&RangeAggregationExpr{
				Op: RangeOpAvg,
				Range: LogRangeExpr{
					Sel: Selector{
						Matchers: []LabelMatcher{
							{"job", OpEq, "mysql"},
						},
					},
					Pipeline: []PipelineStage{
						&LineFilter{Op: OpEq, Value: "error"},
					},
					Unwrap: &UnwrapExpr{
						Label: "duration",
					},
					Range: 5 * time.Hour,
				},
			},
			false,
		},
		{
			`avg_over_time({ job = "mysql" }[5h] |= "error" | unwrap duration(bytes) | foo="bar")`,
			&RangeAggregationExpr{
				Op: RangeOpAvg,
				Range: LogRangeExpr{
					Sel: Selector{
						Matchers: []LabelMatcher{
							{"job", OpEq, "mysql"},
						},
					},
					Pipeline: []PipelineStage{
						&LineFilter{Op: OpEq, Value: "error"},
					},
					Unwrap: &UnwrapExpr{
						Op:    "duration",
						Label: "bytes",
						Filters: []LabelMatcher{
							{"foo", OpEq, "bar"},
						},
					},
					Range: 5 * time.Hour,
				},
			},
			false,
		},
		// Vector aggregation.
		{
			`sum(3.14, rate({job="mysql"}[1m])) without ()`,
			&VectorAggregationExpr{
				Op: VectorOpSum,
				Expr: &RangeAggregationExpr{
					Op: RangeOpRate,
					Range: LogRangeExpr{
						Sel: Selector{
							Matchers: []LabelMatcher{
								{"job", OpEq, "mysql"},
							},
						},
						Range: 1 * time.Minute,
					},
				},
				Parameter: ptrTo(3.14),
				Grouping: &Grouping{
					Op: GroupingOpWithout,
				},
			},
			false,
		},
		{
			`sum by (host) (rate({job="mysql"} |= "error" != "timeout" | json | duration > 10s [1m]))`,
			&VectorAggregationExpr{
				Op: VectorOpSum,
				Expr: &RangeAggregationExpr{
					Op: RangeOpRate,
					Range: LogRangeExpr{
						Sel: Selector{
							Matchers: []LabelMatcher{
								{"job", OpEq, "mysql"},
							},
						},
						Pipeline: []PipelineStage{
							&LineFilter{Op: OpEq, Value: "error"},
							&LineFilter{Op: OpNotEq, Value: "timeout"},
							&JSONExpressionParser{},
							&LabelFilter{
								Pred: &DurationFilter{"duration", OpGt, 10 * time.Second},
							},
						},
						Range: 1 * time.Minute,
					},
				},
				Grouping: &Grouping{
					Op:     GroupingOpBy,
					Labels: []Label{"host"},
				},
			},
			false,
		},
		// label_replace
		{
			`label_replace(rate({}), "dst", "replacement", "src", ".*")`,
			&LabelReplaceExpr{
				Expr: &RangeAggregationExpr{
					Op: RangeOpRate,
				},
				DstLabel:    "dst",
				Replacement: "replacement",
				SrcLabel:    "src",
				Regex:       ".*",
			},
			false,
		},
		// Literal expression.
		{
			`10.0`,
			&LiteralExpr{
				Value: 10.0,
			},
			false,
		},
		{
			`+10.0`,
			&LiteralExpr{
				Value: 10.0,
			},
			false,
		},
		{
			`-10.0`,
			&LiteralExpr{
				Value: -10.0,
			},
			false,
		},
		// Vector expression.
		{
			`vector (10.0)`,
			&VectorExpr{
				Value: 10.0,
			},
			false,
		},
		// Binary op.
		{
			`10.0 + 10.0`,
			&BinOpExpr{
				Left:  &LiteralExpr{Value: 10.0},
				Op:    OpAdd,
				Right: &LiteralExpr{Value: 10.0},
			},
			false,
		},
		{
			`(2+2)*2`,
			&BinOpExpr{
				Left: &ParenExpr{
					X: &BinOpExpr{
						Left:  &LiteralExpr{Value: 2},
						Op:    OpAdd,
						Right: &LiteralExpr{Value: 2},
					},
				},
				Op:    OpMul,
				Right: &LiteralExpr{Value: 2},
			},
			false,
		},
		{
			`2+2*2`,
			&BinOpExpr{
				Left: &LiteralExpr{Value: 2},
				Op:   OpAdd,
				Right: &BinOpExpr{
					Left:  &LiteralExpr{Value: 2},
					Op:    OpMul,
					Right: &LiteralExpr{Value: 2},
				},
			},
			false,
		},
		{
			`1+2*3/4%5`,
			&BinOpExpr{
				Left: &LiteralExpr{Value: 1},
				Op:   OpAdd,
				Right: &BinOpExpr{
					Left: &LiteralExpr{Value: 2},
					Op:   OpMul,
					Right: &BinOpExpr{
						Left: &LiteralExpr{Value: 3},
						Op:   OpDiv,
						Right: &BinOpExpr{
							Left:  &LiteralExpr{Value: 4},
							Op:    OpMod,
							Right: &LiteralExpr{Value: 5},
						},
					},
				},
			},
			false,
		},
		{
			`10.0 and bool 10.0`,
			&BinOpExpr{
				Left: &LiteralExpr{Value: 10.0},
				Op:   OpAnd,
				Modifier: BinOpModifier{
					ReturnBool: true,
				},
				Right: &LiteralExpr{Value: 10.0},
			},
			false,
		},
		{
			`vector(0) and on () vector(0)`,
			&BinOpExpr{
				Left: &VectorExpr{},
				Op:   OpAnd,
				Modifier: BinOpModifier{
					Op: "on",
				},
				Right: &VectorExpr{},
			},
			false,
		},
		{
			`vector(0) and ignoring (foo) group_left vector(0)`,
			&BinOpExpr{
				Left: &VectorExpr{},
				Op:   OpAnd,
				Modifier: BinOpModifier{
					Op:       "ignoring",
					OpLabels: []Label{"foo"},
					Group:    "left",
				},
				Right: &VectorExpr{},
			},
			false,
		},
		{
			`vector(0) and bool ignoring (foo, bar) group_right () vector(0)`,
			&BinOpExpr{
				Left: &VectorExpr{},
				Op:   OpAnd,
				Modifier: BinOpModifier{
					Op:         "ignoring",
					OpLabels:   []Label{"foo", "bar"},
					Group:      "right",
					ReturnBool: true,
				},
				Right: &VectorExpr{},
			},
			false,
		},

		{"{", nil, true},
		{"{foo}", nil, true},
		{"{foo =}", nil, true},
		{`{foo == "bar"}`, nil, true},
		{`{foo = bar}`, nil, true},
		{`{foo = "bar"} | addr == ip(`, nil, true},
		{`{foo = "bar"} | addr == ip()`, nil, true},
		{`{foo = "bar"} | addr >= ip("127.0.0.1")`, nil, true},
		{`{foo = "bar"} | foo == "bar"`, nil, true},
		{`{foo = "bar"} | status = 10`, nil, true},
		{`{foo = "bar"} | status = 10s`, nil, true},
		{`{foo = "bar"} | status = 10b`, nil, true},
	}
	for i, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("Test%d", i+1), func(t *testing.T) {
			defer func() {
				if t.Failed() {
					t.Logf("Input:\n%s", tt.input)
				}
			}()

			got, err := Parse(tt.input)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}
