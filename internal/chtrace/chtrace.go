// Package chtrace implements reading ClickHouse traces.
package chtrace

import (
	"encoding/binary"
	"time"

	"github.com/ClickHouse/ch-go/proto"
	"go.opentelemetry.io/otel/trace"
)

// Table represents otel span table.
type Table struct {
	TraceID         proto.ColUUID                    // trace_id
	SpanID          proto.ColUInt64                  // span_id
	ParentSpanID    proto.ColUInt64                  // parent_span_id
	OperationName   *proto.ColLowCardinality[string] // operation_name
	StartTimeMicro  proto.ColUInt64                  // start_time_us
	FinishTimeMicro proto.ColUInt64                  // finish_time_us
	FinishDate      proto.ColDate                    // finish_date
	Attributes      proto.ColMap[string, string]     // attribute
}

// Rows returns Trace per row.
func (t Table) Rows() []Trace {
	var out []Trace
	for i := 0; i < t.TraceID.Rows(); i++ {
		tt := Trace{
			TraceID:       trace.TraceID(t.TraceID.Row(i)),
			StartTime:     time.UnixMicro(int64(t.StartTimeMicro.Row(i))),
			FinishTime:    time.UnixMicro(int64(t.FinishTimeMicro.Row(i))),
			Attributes:    t.Attributes.Row(i),
			OperationName: t.OperationName.Row(i),
		}
		binary.BigEndian.PutUint64(tt.SpanID[:], t.SpanID.Row(i))
		binary.BigEndian.PutUint64(tt.ParentSpanID[:], t.ParentSpanID.Row(i))

		out = append(out, tt)
	}

	return out
}

// Trace is a single trace span.
type Trace struct {
	TraceID       trace.TraceID
	SpanID        trace.SpanID
	ParentSpanID  trace.SpanID
	OperationName string
	StartTime     time.Time
	FinishTime    time.Time
	Attributes    map[string]string
}

// Result returns proto.Results for Table.
func (t *Table) Result() proto.Results {
	return proto.Results{
		{Name: "trace_id", Data: &t.TraceID},
		{Name: "span_id", Data: &t.SpanID},
		{Name: "parent_span_id", Data: &t.ParentSpanID},
		{Name: "operation_name", Data: t.OperationName},
		{Name: "start_time_us", Data: &t.StartTimeMicro},
		{Name: "finish_time_us", Data: &t.FinishTimeMicro},
		{Name: "finish_date", Data: &t.FinishDate},
		{Name: "attribute", Data: &t.Attributes},
	}
}

// Columns returns column names (and mappings) for doing SELECTS.
func (t *Table) Columns() []string {
	var out []string
	for _, v := range t.Result() {
		switch v.Name {
		case "attribute_names":
			out = append(out, "mapKeys(attribute) as attribute_names")
		case "attribute_values":
			out = append(out, "mapValues(attribute) as attribute_values")
		default:
			out = append(out, v.Name)
		}
	}
	return out
}

// NewTable creates and initializes new Table.
func NewTable() *Table {
	return &Table{
		OperationName: new(proto.ColStr).LowCardinality(),
		Attributes: proto.ColMap[string, string]{
			Keys:   new(proto.ColStr).LowCardinality(),
			Values: new(proto.ColStr),
		},
	}
}
