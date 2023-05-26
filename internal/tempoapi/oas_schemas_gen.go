// Code generated by ogen, DO NOT EDIT.

package tempoapi

import (
	"fmt"
	"io"
	"time"
)

func (s *ErrorStatusCode) Error() string {
	return fmt.Sprintf("code %d: %+v", s.StatusCode, s.Response)
}

// Ref: #/components/schemas/AnyValue
// AnyValue represents sum type.
type AnyValue struct {
	Type        AnyValueType // switch on this field
	StringValue StringValue
	BoolValue   BoolValue
	IntValue    IntValue
	DoubleValue DoubleValue
	ArrayValue  ArrayValue
	KvlistValue KvlistValue
	BytesValue  BytesValue
}

// AnyValueType is oneOf type of AnyValue.
type AnyValueType string

// Possible values for AnyValueType.
const (
	StringValueAnyValue AnyValueType = "StringValue"
	BoolValueAnyValue   AnyValueType = "BoolValue"
	IntValueAnyValue    AnyValueType = "IntValue"
	DoubleValueAnyValue AnyValueType = "DoubleValue"
	ArrayValueAnyValue  AnyValueType = "ArrayValue"
	KvlistValueAnyValue AnyValueType = "KvlistValue"
	BytesValueAnyValue  AnyValueType = "BytesValue"
)

// IsStringValue reports whether AnyValue is StringValue.
func (s AnyValue) IsStringValue() bool { return s.Type == StringValueAnyValue }

// IsBoolValue reports whether AnyValue is BoolValue.
func (s AnyValue) IsBoolValue() bool { return s.Type == BoolValueAnyValue }

// IsIntValue reports whether AnyValue is IntValue.
func (s AnyValue) IsIntValue() bool { return s.Type == IntValueAnyValue }

// IsDoubleValue reports whether AnyValue is DoubleValue.
func (s AnyValue) IsDoubleValue() bool { return s.Type == DoubleValueAnyValue }

// IsArrayValue reports whether AnyValue is ArrayValue.
func (s AnyValue) IsArrayValue() bool { return s.Type == ArrayValueAnyValue }

// IsKvlistValue reports whether AnyValue is KvlistValue.
func (s AnyValue) IsKvlistValue() bool { return s.Type == KvlistValueAnyValue }

// IsBytesValue reports whether AnyValue is BytesValue.
func (s AnyValue) IsBytesValue() bool { return s.Type == BytesValueAnyValue }

// SetStringValue sets AnyValue to StringValue.
func (s *AnyValue) SetStringValue(v StringValue) {
	s.Type = StringValueAnyValue
	s.StringValue = v
}

// GetStringValue returns StringValue and true boolean if AnyValue is StringValue.
func (s AnyValue) GetStringValue() (v StringValue, ok bool) {
	if !s.IsStringValue() {
		return v, false
	}
	return s.StringValue, true
}

// NewStringValueAnyValue returns new AnyValue from StringValue.
func NewStringValueAnyValue(v StringValue) AnyValue {
	var s AnyValue
	s.SetStringValue(v)
	return s
}

// SetBoolValue sets AnyValue to BoolValue.
func (s *AnyValue) SetBoolValue(v BoolValue) {
	s.Type = BoolValueAnyValue
	s.BoolValue = v
}

// GetBoolValue returns BoolValue and true boolean if AnyValue is BoolValue.
func (s AnyValue) GetBoolValue() (v BoolValue, ok bool) {
	if !s.IsBoolValue() {
		return v, false
	}
	return s.BoolValue, true
}

// NewBoolValueAnyValue returns new AnyValue from BoolValue.
func NewBoolValueAnyValue(v BoolValue) AnyValue {
	var s AnyValue
	s.SetBoolValue(v)
	return s
}

// SetIntValue sets AnyValue to IntValue.
func (s *AnyValue) SetIntValue(v IntValue) {
	s.Type = IntValueAnyValue
	s.IntValue = v
}

// GetIntValue returns IntValue and true boolean if AnyValue is IntValue.
func (s AnyValue) GetIntValue() (v IntValue, ok bool) {
	if !s.IsIntValue() {
		return v, false
	}
	return s.IntValue, true
}

// NewIntValueAnyValue returns new AnyValue from IntValue.
func NewIntValueAnyValue(v IntValue) AnyValue {
	var s AnyValue
	s.SetIntValue(v)
	return s
}

// SetDoubleValue sets AnyValue to DoubleValue.
func (s *AnyValue) SetDoubleValue(v DoubleValue) {
	s.Type = DoubleValueAnyValue
	s.DoubleValue = v
}

// GetDoubleValue returns DoubleValue and true boolean if AnyValue is DoubleValue.
func (s AnyValue) GetDoubleValue() (v DoubleValue, ok bool) {
	if !s.IsDoubleValue() {
		return v, false
	}
	return s.DoubleValue, true
}

// NewDoubleValueAnyValue returns new AnyValue from DoubleValue.
func NewDoubleValueAnyValue(v DoubleValue) AnyValue {
	var s AnyValue
	s.SetDoubleValue(v)
	return s
}

// SetArrayValue sets AnyValue to ArrayValue.
func (s *AnyValue) SetArrayValue(v ArrayValue) {
	s.Type = ArrayValueAnyValue
	s.ArrayValue = v
}

// GetArrayValue returns ArrayValue and true boolean if AnyValue is ArrayValue.
func (s AnyValue) GetArrayValue() (v ArrayValue, ok bool) {
	if !s.IsArrayValue() {
		return v, false
	}
	return s.ArrayValue, true
}

// NewArrayValueAnyValue returns new AnyValue from ArrayValue.
func NewArrayValueAnyValue(v ArrayValue) AnyValue {
	var s AnyValue
	s.SetArrayValue(v)
	return s
}

// SetKvlistValue sets AnyValue to KvlistValue.
func (s *AnyValue) SetKvlistValue(v KvlistValue) {
	s.Type = KvlistValueAnyValue
	s.KvlistValue = v
}

// GetKvlistValue returns KvlistValue and true boolean if AnyValue is KvlistValue.
func (s AnyValue) GetKvlistValue() (v KvlistValue, ok bool) {
	if !s.IsKvlistValue() {
		return v, false
	}
	return s.KvlistValue, true
}

// NewKvlistValueAnyValue returns new AnyValue from KvlistValue.
func NewKvlistValueAnyValue(v KvlistValue) AnyValue {
	var s AnyValue
	s.SetKvlistValue(v)
	return s
}

// SetBytesValue sets AnyValue to BytesValue.
func (s *AnyValue) SetBytesValue(v BytesValue) {
	s.Type = BytesValueAnyValue
	s.BytesValue = v
}

// GetBytesValue returns BytesValue and true boolean if AnyValue is BytesValue.
func (s AnyValue) GetBytesValue() (v BytesValue, ok bool) {
	if !s.IsBytesValue() {
		return v, false
	}
	return s.BytesValue, true
}

// NewBytesValueAnyValue returns new AnyValue from BytesValue.
func NewBytesValueAnyValue(v BytesValue) AnyValue {
	var s AnyValue
	s.SetBytesValue(v)
	return s
}

// Ref: #/components/schemas/ArrayValue
type ArrayValue struct {
	ArrayValue []AnyValue `json:"arrayValue"`
}

// GetArrayValue returns the value of ArrayValue.
func (s *ArrayValue) GetArrayValue() []AnyValue {
	return s.ArrayValue
}

// SetArrayValue sets the value of ArrayValue.
func (s *ArrayValue) SetArrayValue(val []AnyValue) {
	s.ArrayValue = val
}

type Attributes []KeyValue

// Ref: #/components/schemas/BoolValue
type BoolValue struct {
	BoolValue bool `json:"boolValue"`
}

// GetBoolValue returns the value of BoolValue.
func (s *BoolValue) GetBoolValue() bool {
	return s.BoolValue
}

// SetBoolValue sets the value of BoolValue.
func (s *BoolValue) SetBoolValue(val bool) {
	s.BoolValue = val
}

// Ref: #/components/schemas/BytesValue
type BytesValue struct {
	BytesValue []byte `json:"bytesValue"`
}

// GetBytesValue returns the value of BytesValue.
func (s *BytesValue) GetBytesValue() []byte {
	return s.BytesValue
}

// SetBytesValue sets the value of BytesValue.
func (s *BytesValue) SetBytesValue(val []byte) {
	s.BytesValue = val
}

// Ref: #/components/schemas/DoubleValue
type DoubleValue struct {
	DoubleValue float64 `json:"doubleValue"`
}

// GetDoubleValue returns the value of DoubleValue.
func (s *DoubleValue) GetDoubleValue() float64 {
	return s.DoubleValue
}

// SetDoubleValue sets the value of DoubleValue.
func (s *DoubleValue) SetDoubleValue(val float64) {
	s.DoubleValue = val
}

type Error string

// ErrorStatusCode wraps Error with StatusCode.
type ErrorStatusCode struct {
	StatusCode int
	Response   Error
}

// GetStatusCode returns the value of StatusCode.
func (s *ErrorStatusCode) GetStatusCode() int {
	return s.StatusCode
}

// GetResponse returns the value of Response.
func (s *ErrorStatusCode) GetResponse() Error {
	return s.Response
}

// SetStatusCode sets the value of StatusCode.
func (s *ErrorStatusCode) SetStatusCode(val int) {
	s.StatusCode = val
}

// SetResponse sets the value of Response.
func (s *ErrorStatusCode) SetResponse(val Error) {
	s.Response = val
}

// Ref: #/components/schemas/IntValue
type IntValue struct {
	IntValue int64 `json:"intValue"`
}

// GetIntValue returns the value of IntValue.
func (s *IntValue) GetIntValue() int64 {
	return s.IntValue
}

// SetIntValue sets the value of IntValue.
func (s *IntValue) SetIntValue(val int64) {
	s.IntValue = val
}

// Ref: #/components/schemas/KeyValue
type KeyValue struct {
	Key   string   `json:"key"`
	Value AnyValue `json:"value"`
}

// GetKey returns the value of Key.
func (s *KeyValue) GetKey() string {
	return s.Key
}

// GetValue returns the value of Value.
func (s *KeyValue) GetValue() AnyValue {
	return s.Value
}

// SetKey sets the value of Key.
func (s *KeyValue) SetKey(val string) {
	s.Key = val
}

// SetValue sets the value of Value.
func (s *KeyValue) SetValue(val AnyValue) {
	s.Value = val
}

// Ref: #/components/schemas/KvlistValue
type KvlistValue struct {
	KvlistValue []KeyValue `json:"kvlistValue"`
}

// GetKvlistValue returns the value of KvlistValue.
func (s *KvlistValue) GetKvlistValue() []KeyValue {
	return s.KvlistValue
}

// SetKvlistValue sets the value of KvlistValue.
func (s *KvlistValue) SetKvlistValue(val []KeyValue) {
	s.KvlistValue = val
}

// NewOptDuration returns new OptDuration with value set to v.
func NewOptDuration(v time.Duration) OptDuration {
	return OptDuration{
		Value: v,
		Set:   true,
	}
}

// OptDuration is optional time.Duration.
type OptDuration struct {
	Value time.Duration
	Set   bool
}

// IsSet returns true if OptDuration was set.
func (o OptDuration) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptDuration) Reset() {
	var v time.Duration
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptDuration) SetTo(v time.Duration) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptDuration) Get() (v time.Duration, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptDuration) Or(d time.Duration) time.Duration {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptInt returns new OptInt with value set to v.
func NewOptInt(v int) OptInt {
	return OptInt{
		Value: v,
		Set:   true,
	}
}

// OptInt is optional int.
type OptInt struct {
	Value int
	Set   bool
}

// IsSet returns true if OptInt was set.
func (o OptInt) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt) SetTo(v int) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt) Get() (v int, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptUnixSeconds returns new OptUnixSeconds with value set to v.
func NewOptUnixSeconds(v time.Time) OptUnixSeconds {
	return OptUnixSeconds{
		Value: v,
		Set:   true,
	}
}

// OptUnixSeconds is optional time.Time.
type OptUnixSeconds struct {
	Value time.Time
	Set   bool
}

// IsSet returns true if OptUnixSeconds was set.
func (o OptUnixSeconds) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptUnixSeconds) Reset() {
	var v time.Time
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptUnixSeconds) SetTo(v time.Time) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptUnixSeconds) Get() (v time.Time, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptUnixSeconds) Or(d time.Time) time.Time {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Ref: #/components/schemas/StringValue
type StringValue struct {
	StringValue string `json:"stringValue"`
}

// GetStringValue returns the value of StringValue.
func (s *StringValue) GetStringValue() string {
	return s.StringValue
}

// SetStringValue sets the value of StringValue.
func (s *StringValue) SetStringValue(val string) {
	s.StringValue = val
}

// Ref: #/components/schemas/TagNames
type TagNames struct {
	TagNames []string `json:"tagNames"`
}

// GetTagNames returns the value of TagNames.
func (s *TagNames) GetTagNames() []string {
	return s.TagNames
}

// SetTagNames sets the value of TagNames.
func (s *TagNames) SetTagNames(val []string) {
	s.TagNames = val
}

// Ref: #/components/schemas/TagValue
type TagValue struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// GetType returns the value of Type.
func (s *TagValue) GetType() string {
	return s.Type
}

// GetValue returns the value of Value.
func (s *TagValue) GetValue() string {
	return s.Value
}

// SetType sets the value of Type.
func (s *TagValue) SetType(val string) {
	s.Type = val
}

// SetValue sets the value of Value.
func (s *TagValue) SetValue(val string) {
	s.Value = val
}

// Ref: #/components/schemas/TagValues
type TagValues struct {
	TagValues []string `json:"tagValues"`
}

// GetTagValues returns the value of TagValues.
func (s *TagValues) GetTagValues() []string {
	return s.TagValues
}

// SetTagValues sets the value of TagValues.
func (s *TagValues) SetTagValues(val []string) {
	s.TagValues = val
}

// Ref: #/components/schemas/TagValuesV2
type TagValuesV2 struct {
	TagValues []TagValue `json:"tagValues"`
}

// GetTagValues returns the value of TagValues.
func (s *TagValuesV2) GetTagValues() []TagValue {
	return s.TagValues
}

// SetTagValues sets the value of TagValues.
func (s *TagValuesV2) SetTagValues(val []TagValue) {
	s.TagValues = val
}

// Ref: #/components/schemas/TempoSpan
type TempoSpan struct {
	SpanID            string      `json:"spanID"`
	Name              string      `json:"name"`
	StartTimeUnixNano time.Time   `json:"startTimeUnixNano"`
	DurationNanos     int64       `json:"durationNanos"`
	Attributes        *Attributes `json:"attributes"`
}

// GetSpanID returns the value of SpanID.
func (s *TempoSpan) GetSpanID() string {
	return s.SpanID
}

// GetName returns the value of Name.
func (s *TempoSpan) GetName() string {
	return s.Name
}

// GetStartTimeUnixNano returns the value of StartTimeUnixNano.
func (s *TempoSpan) GetStartTimeUnixNano() time.Time {
	return s.StartTimeUnixNano
}

// GetDurationNanos returns the value of DurationNanos.
func (s *TempoSpan) GetDurationNanos() int64 {
	return s.DurationNanos
}

// GetAttributes returns the value of Attributes.
func (s *TempoSpan) GetAttributes() *Attributes {
	return s.Attributes
}

// SetSpanID sets the value of SpanID.
func (s *TempoSpan) SetSpanID(val string) {
	s.SpanID = val
}

// SetName sets the value of Name.
func (s *TempoSpan) SetName(val string) {
	s.Name = val
}

// SetStartTimeUnixNano sets the value of StartTimeUnixNano.
func (s *TempoSpan) SetStartTimeUnixNano(val time.Time) {
	s.StartTimeUnixNano = val
}

// SetDurationNanos sets the value of DurationNanos.
func (s *TempoSpan) SetDurationNanos(val int64) {
	s.DurationNanos = val
}

// SetAttributes sets the value of Attributes.
func (s *TempoSpan) SetAttributes(val *Attributes) {
	s.Attributes = val
}

// Ref: #/components/schemas/TempoSpanSet
type TempoSpanSet struct {
	Spans      []TempoSpan `json:"spans"`
	Matched    OptInt      `json:"matched"`
	Attributes *Attributes `json:"attributes"`
}

// GetSpans returns the value of Spans.
func (s *TempoSpanSet) GetSpans() []TempoSpan {
	return s.Spans
}

// GetMatched returns the value of Matched.
func (s *TempoSpanSet) GetMatched() OptInt {
	return s.Matched
}

// GetAttributes returns the value of Attributes.
func (s *TempoSpanSet) GetAttributes() *Attributes {
	return s.Attributes
}

// SetSpans sets the value of Spans.
func (s *TempoSpanSet) SetSpans(val []TempoSpan) {
	s.Spans = val
}

// SetMatched sets the value of Matched.
func (s *TempoSpanSet) SetMatched(val OptInt) {
	s.Matched = val
}

// SetAttributes sets the value of Attributes.
func (s *TempoSpanSet) SetAttributes(val *Attributes) {
	s.Attributes = val
}

type TraceByID struct {
	Data io.Reader
}

// Read reads data from the Data reader.
//
// Kept to satisfy the io.Reader interface.
func (s TraceByID) Read(p []byte) (n int, err error) {
	if s.Data == nil {
		return 0, nil
	}
	return s.Data.Read(p)
}

func (*TraceByID) traceByIDRes() {}

// Ref: #/components/responses/TraceByIDNotFound
type TraceByIDNotFound struct{}

func (*TraceByIDNotFound) traceByIDRes() {}

// Ref: #/components/schemas/TraceSearchMetadata
type TraceSearchMetadata struct {
	TraceID           string       `json:"traceID"`
	RootServiceName   string       `json:"rootServiceName"`
	RootTraceName     string       `json:"rootTraceName"`
	StartTimeUnixNano time.Time    `json:"startTimeUnixNano"`
	DurationMs        int          `json:"durationMs"`
	SpanSet           TempoSpanSet `json:"spanSet"`
}

// GetTraceID returns the value of TraceID.
func (s *TraceSearchMetadata) GetTraceID() string {
	return s.TraceID
}

// GetRootServiceName returns the value of RootServiceName.
func (s *TraceSearchMetadata) GetRootServiceName() string {
	return s.RootServiceName
}

// GetRootTraceName returns the value of RootTraceName.
func (s *TraceSearchMetadata) GetRootTraceName() string {
	return s.RootTraceName
}

// GetStartTimeUnixNano returns the value of StartTimeUnixNano.
func (s *TraceSearchMetadata) GetStartTimeUnixNano() time.Time {
	return s.StartTimeUnixNano
}

// GetDurationMs returns the value of DurationMs.
func (s *TraceSearchMetadata) GetDurationMs() int {
	return s.DurationMs
}

// GetSpanSet returns the value of SpanSet.
func (s *TraceSearchMetadata) GetSpanSet() TempoSpanSet {
	return s.SpanSet
}

// SetTraceID sets the value of TraceID.
func (s *TraceSearchMetadata) SetTraceID(val string) {
	s.TraceID = val
}

// SetRootServiceName sets the value of RootServiceName.
func (s *TraceSearchMetadata) SetRootServiceName(val string) {
	s.RootServiceName = val
}

// SetRootTraceName sets the value of RootTraceName.
func (s *TraceSearchMetadata) SetRootTraceName(val string) {
	s.RootTraceName = val
}

// SetStartTimeUnixNano sets the value of StartTimeUnixNano.
func (s *TraceSearchMetadata) SetStartTimeUnixNano(val time.Time) {
	s.StartTimeUnixNano = val
}

// SetDurationMs sets the value of DurationMs.
func (s *TraceSearchMetadata) SetDurationMs(val int) {
	s.DurationMs = val
}

// SetSpanSet sets the value of SpanSet.
func (s *TraceSearchMetadata) SetSpanSet(val TempoSpanSet) {
	s.SpanSet = val
}

// Ref: #/components/schemas/Traces
type Traces struct {
	Traces []TraceSearchMetadata `json:"traces"`
}

// GetTraces returns the value of Traces.
func (s *Traces) GetTraces() []TraceSearchMetadata {
	return s.Traces
}

// SetTraces sets the value of Traces.
func (s *Traces) SetTraces(val []TraceSearchMetadata) {
	s.Traces = val
}
