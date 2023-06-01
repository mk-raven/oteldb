package ytstore

import (
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"

	"github.com/go-faster/oteldb/internal/tempoapi"
)

// FillTraceMetadata files TraceSearchMetadata fields using span.
//
// The span should be a parent span.
func (span Span) FillTraceMetadata(m *tempoapi.TraceSearchMetadata) {
	ss := &m.SpanSet

	m.RootTraceName = span.Name
	if attr, ok := pcommon.Map(span.ResourceAttrs).Get("service.name"); ok {
		m.RootServiceName = attr.AsString()
	}
	var (
		start = time.Unix(0, int64(span.Start))
		end   = time.Unix(0, int64(span.End))
	)

	m.StartTimeUnixNano = start
	m.DurationMs = int(end.Sub(start).Milliseconds())
	if ss.Attributes == nil {
		ss.Attributes = new(tempoapi.Attributes)
	}
	ytToTempoAttrs(ss.Attributes, span.ScopeAttrs)
	ytToTempoAttrs(ss.Attributes, span.ResourceAttrs)
}

// AsTempoSpan converts span to TempoSpan.
func (span Span) AsTempoSpan() (s tempoapi.TempoSpan) {
	s = tempoapi.TempoSpan{
		SpanID:            span.SpanID.Hex(),
		Name:              span.Name,
		StartTimeUnixNano: time.Unix(0, int64(span.Start)),
		DurationNanos:     int64(span.End - span.Start),
		Attributes:        &tempoapi.Attributes{},
	}
	ytToTempoAttrs(s.Attributes, span.Attrs)
	return s
}

func ytToTempoAttrs(to *tempoapi.Attributes, from Attrs) {
	var convertValue func(val pcommon.Value) (r tempoapi.AnyValue)
	convertValue = func(val pcommon.Value) (r tempoapi.AnyValue) {
		switch val.Type() {
		case pcommon.ValueTypeStr:
			r.SetStringValue(tempoapi.StringValue{StringValue: val.Str()})
		case pcommon.ValueTypeBool:
			r.SetBoolValue(tempoapi.BoolValue{BoolValue: val.Bool()})
		case pcommon.ValueTypeInt:
			r.SetIntValue(tempoapi.IntValue{IntValue: val.Int()})
		case pcommon.ValueTypeDouble:
			r.SetDoubleValue(tempoapi.DoubleValue{DoubleValue: val.Double()})
		case pcommon.ValueTypeMap:
			m := tempoapi.KvlistValue{}
			val.Map().Range(func(k string, v pcommon.Value) bool {
				m.KvlistValue = append(m.KvlistValue, tempoapi.KeyValue{
					Key:   k,
					Value: convertValue(v),
				})
				return true
			})
			r.SetKvlistValue(m)
		case pcommon.ValueTypeSlice:
			a := tempoapi.ArrayValue{}
			ss := val.Slice()
			for i := 0; i < ss.Len(); i++ {
				v := ss.At(i)
				a.ArrayValue = append(a.ArrayValue, convertValue(v))
			}
			r.SetArrayValue(a)
		case pcommon.ValueTypeBytes:
			r.SetBytesValue(tempoapi.BytesValue{BytesValue: val.Bytes().AsRaw()})
		default:
			r.Type = tempoapi.StringValueAnyValue
		}
		return r
	}

	pcommon.Map(from).Range(func(k string, v pcommon.Value) bool {
		*to = append(*to, tempoapi.KeyValue{
			Key:   k,
			Value: convertValue(v),
		})
		return true
	})
}

type metadataCollector struct {
	metadatas map[TraceID]tempoapi.TraceSearchMetadata
}

func (b *metadataCollector) init() {
	if b.metadatas == nil {
		b.metadatas = make(map[TraceID]tempoapi.TraceSearchMetadata)
	}
}

func (b *metadataCollector) AddSpan(span Span) error {
	b.init()

	traceID := span.TraceID
	m, ok := b.metadatas[traceID]
	if !ok {
		m = tempoapi.TraceSearchMetadata{
			TraceID: traceID.Hex(),
		}
	}
	ss := &m.SpanSet

	if span.ParentSpanID.IsEmpty() {
		span.FillTraceMetadata(&m)
	}
	ss.Spans = append(ss.Spans, span.AsTempoSpan())

	// Put modified struct back to map.
	b.metadatas[traceID] = m
	return nil
}

func (b *metadataCollector) Result() map[TraceID]tempoapi.TraceSearchMetadata {
	b.init()
	return b.metadatas
}