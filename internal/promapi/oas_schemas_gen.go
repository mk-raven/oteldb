// Code generated by ogen, DO NOT EDIT.

package promapi

import (
	"fmt"

	"github.com/go-faster/errors"
)

func (s *FailStatusCode) Error() string {
	return fmt.Sprintf("code %d: %+v", s.StatusCode, s.Response)
}

// Ref: #/components/schemas/Data
// Data represents sum type.
type Data struct {
	Type   DataType // switch on this field
	Matrix Matrix
	Vector Vector
	Scalar Scalar
	String String
}

// DataType is oneOf type of Data.
type DataType string

// Possible values for DataType.
const (
	MatrixData DataType = "Matrix"
	VectorData DataType = "Vector"
	ScalarData DataType = "Scalar"
	StringData DataType = "String"
)

// IsMatrix reports whether Data is Matrix.
func (s Data) IsMatrix() bool { return s.Type == MatrixData }

// IsVector reports whether Data is Vector.
func (s Data) IsVector() bool { return s.Type == VectorData }

// IsScalar reports whether Data is Scalar.
func (s Data) IsScalar() bool { return s.Type == ScalarData }

// IsString reports whether Data is String.
func (s Data) IsString() bool { return s.Type == StringData }

// SetMatrix sets Data to Matrix.
func (s *Data) SetMatrix(v Matrix) {
	s.Type = MatrixData
	s.Matrix = v
}

// GetMatrix returns Matrix and true boolean if Data is Matrix.
func (s Data) GetMatrix() (v Matrix, ok bool) {
	if !s.IsMatrix() {
		return v, false
	}
	return s.Matrix, true
}

// NewMatrixData returns new Data from Matrix.
func NewMatrixData(v Matrix) Data {
	var s Data
	s.SetMatrix(v)
	return s
}

// SetVector sets Data to Vector.
func (s *Data) SetVector(v Vector) {
	s.Type = VectorData
	s.Vector = v
}

// GetVector returns Vector and true boolean if Data is Vector.
func (s Data) GetVector() (v Vector, ok bool) {
	if !s.IsVector() {
		return v, false
	}
	return s.Vector, true
}

// NewVectorData returns new Data from Vector.
func NewVectorData(v Vector) Data {
	var s Data
	s.SetVector(v)
	return s
}

// SetScalar sets Data to Scalar.
func (s *Data) SetScalar(v Scalar) {
	s.Type = ScalarData
	s.Scalar = v
}

// GetScalar returns Scalar and true boolean if Data is Scalar.
func (s Data) GetScalar() (v Scalar, ok bool) {
	if !s.IsScalar() {
		return v, false
	}
	return s.Scalar, true
}

// NewScalarData returns new Data from Scalar.
func NewScalarData(v Scalar) Data {
	var s Data
	s.SetScalar(v)
	return s
}

// SetString sets Data to String.
func (s *Data) SetString(v String) {
	s.Type = StringData
	s.String = v
}

// GetString returns String and true boolean if Data is String.
func (s Data) GetString() (v String, ok bool) {
	if !s.IsString() {
		return v, false
	}
	return s.String, true
}

// NewStringData returns new Data from String.
func NewStringData(v String) Data {
	var s Data
	s.SetString(v)
	return s
}

// May still contain data.
// Ref: #/components/schemas/Fail
type Fail struct {
	Status    string        `json:"status"`
	Error     string        `json:"error"`
	ErrorType FailErrorType `json:"errorType"`
	Data      OptData       `json:"data"`
}

// GetStatus returns the value of Status.
func (s *Fail) GetStatus() string {
	return s.Status
}

// GetError returns the value of Error.
func (s *Fail) GetError() string {
	return s.Error
}

// GetErrorType returns the value of ErrorType.
func (s *Fail) GetErrorType() FailErrorType {
	return s.ErrorType
}

// GetData returns the value of Data.
func (s *Fail) GetData() OptData {
	return s.Data
}

// SetStatus sets the value of Status.
func (s *Fail) SetStatus(val string) {
	s.Status = val
}

// SetError sets the value of Error.
func (s *Fail) SetError(val string) {
	s.Error = val
}

// SetErrorType sets the value of ErrorType.
func (s *Fail) SetErrorType(val FailErrorType) {
	s.ErrorType = val
}

// SetData sets the value of Data.
func (s *Fail) SetData(val OptData) {
	s.Data = val
}

type FailErrorType string

const (
	FailErrorTypeTimeout     FailErrorType = "timeout"
	FailErrorTypeCanceled    FailErrorType = "canceled"
	FailErrorTypeExecution   FailErrorType = "execution"
	FailErrorTypeBadData     FailErrorType = "bad_data"
	FailErrorTypeInternal    FailErrorType = "internal"
	FailErrorTypeUnavailable FailErrorType = "unavailable"
	FailErrorTypeNotFound    FailErrorType = "not_found"
)

// MarshalText implements encoding.TextMarshaler.
func (s FailErrorType) MarshalText() ([]byte, error) {
	switch s {
	case FailErrorTypeTimeout:
		return []byte(s), nil
	case FailErrorTypeCanceled:
		return []byte(s), nil
	case FailErrorTypeExecution:
		return []byte(s), nil
	case FailErrorTypeBadData:
		return []byte(s), nil
	case FailErrorTypeInternal:
		return []byte(s), nil
	case FailErrorTypeUnavailable:
		return []byte(s), nil
	case FailErrorTypeNotFound:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *FailErrorType) UnmarshalText(data []byte) error {
	switch FailErrorType(data) {
	case FailErrorTypeTimeout:
		*s = FailErrorTypeTimeout
		return nil
	case FailErrorTypeCanceled:
		*s = FailErrorTypeCanceled
		return nil
	case FailErrorTypeExecution:
		*s = FailErrorTypeExecution
		return nil
	case FailErrorTypeBadData:
		*s = FailErrorTypeBadData
		return nil
	case FailErrorTypeInternal:
		*s = FailErrorTypeInternal
		return nil
	case FailErrorTypeUnavailable:
		*s = FailErrorTypeUnavailable
		return nil
	case FailErrorTypeNotFound:
		*s = FailErrorTypeNotFound
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// FailStatusCode wraps Fail with StatusCode.
type FailStatusCode struct {
	StatusCode int
	Response   Fail
}

// GetStatusCode returns the value of StatusCode.
func (s *FailStatusCode) GetStatusCode() int {
	return s.StatusCode
}

// GetResponse returns the value of Response.
func (s *FailStatusCode) GetResponse() Fail {
	return s.Response
}

// SetStatusCode sets the value of StatusCode.
func (s *FailStatusCode) SetStatusCode(val int) {
	s.StatusCode = val
}

// SetResponse sets the value of Response.
func (s *FailStatusCode) SetResponse(val Fail) {
	s.Response = val
}

// Ref: #/components/schemas/Matrix
type Matrix struct {
	Result []MatrixResultItem `json:"result"`
}

// GetResult returns the value of Result.
func (s *Matrix) GetResult() []MatrixResultItem {
	return s.Result
}

// SetResult sets the value of Result.
func (s *Matrix) SetResult(val []MatrixResultItem) {
	s.Result = val
}

type MatrixResultItem struct {
	Metric MatrixResultItemMetric `json:"metric"`
	Values []Value                `json:"values"`
}

// GetMetric returns the value of Metric.
func (s *MatrixResultItem) GetMetric() MatrixResultItemMetric {
	return s.Metric
}

// GetValues returns the value of Values.
func (s *MatrixResultItem) GetValues() []Value {
	return s.Values
}

// SetMetric sets the value of Metric.
func (s *MatrixResultItem) SetMetric(val MatrixResultItemMetric) {
	s.Metric = val
}

// SetValues sets the value of Values.
func (s *MatrixResultItem) SetValues(val []Value) {
	s.Values = val
}

type MatrixResultItemMetric map[string]string

func (s *MatrixResultItemMetric) init() MatrixResultItemMetric {
	m := *s
	if m == nil {
		m = map[string]string{}
		*s = m
	}
	return m
}

// NewOptData returns new OptData with value set to v.
func NewOptData(v Data) OptData {
	return OptData{
		Value: v,
		Set:   true,
	}
}

// OptData is optional Data.
type OptData struct {
	Value Data
	Set   bool
}

// IsSet returns true if OptData was set.
func (o OptData) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptData) Reset() {
	var v Data
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptData) SetTo(v Data) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptData) Get() (v Data, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptData) Or(d Data) Data {
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

// Ref: #/components/schemas/Scalar
type Scalar struct {
	Result Value `json:"result"`
}

// GetResult returns the value of Result.
func (s *Scalar) GetResult() Value {
	return s.Result
}

// SetResult sets the value of Result.
func (s *Scalar) SetResult(val Value) {
	s.Result = val
}

// Ref: #/components/schemas/String
type String struct {
	Result string `json:"result"`
}

// GetResult returns the value of Result.
func (s *String) GetResult() string {
	return s.Result
}

// SetResult sets the value of Result.
func (s *String) SetResult(val string) {
	s.Result = val
}

// Ref: #/components/schemas/Success
type Success struct {
	// Always 'success'.
	Status string `json:"status"`
	// Only if there were warnings while executing the request. There will still be data in the data
	// field.
	Warnings []string `json:"warnings"`
	Data     Data     `json:"data"`
}

// GetStatus returns the value of Status.
func (s *Success) GetStatus() string {
	return s.Status
}

// GetWarnings returns the value of Warnings.
func (s *Success) GetWarnings() []string {
	return s.Warnings
}

// GetData returns the value of Data.
func (s *Success) GetData() Data {
	return s.Data
}

// SetStatus sets the value of Status.
func (s *Success) SetStatus(val string) {
	s.Status = val
}

// SetWarnings sets the value of Warnings.
func (s *Success) SetWarnings(val []string) {
	s.Warnings = val
}

// SetData sets the value of Data.
func (s *Success) SetData(val Data) {
	s.Data = val
}

type Value []ValueItem

// ValueItem represents sum type.
type ValueItem struct {
	Type    ValueItemType // switch on this field
	Float64 float64
	String  string
}

// ValueItemType is oneOf type of ValueItem.
type ValueItemType string

// Possible values for ValueItemType.
const (
	Float64ValueItem ValueItemType = "float64"
	StringValueItem  ValueItemType = "string"
)

// IsFloat64 reports whether ValueItem is float64.
func (s ValueItem) IsFloat64() bool { return s.Type == Float64ValueItem }

// IsString reports whether ValueItem is string.
func (s ValueItem) IsString() bool { return s.Type == StringValueItem }

// SetFloat64 sets ValueItem to float64.
func (s *ValueItem) SetFloat64(v float64) {
	s.Type = Float64ValueItem
	s.Float64 = v
}

// GetFloat64 returns float64 and true boolean if ValueItem is float64.
func (s ValueItem) GetFloat64() (v float64, ok bool) {
	if !s.IsFloat64() {
		return v, false
	}
	return s.Float64, true
}

// NewFloat64ValueItem returns new ValueItem from float64.
func NewFloat64ValueItem(v float64) ValueItem {
	var s ValueItem
	s.SetFloat64(v)
	return s
}

// SetString sets ValueItem to string.
func (s *ValueItem) SetString(v string) {
	s.Type = StringValueItem
	s.String = v
}

// GetString returns string and true boolean if ValueItem is string.
func (s ValueItem) GetString() (v string, ok bool) {
	if !s.IsString() {
		return v, false
	}
	return s.String, true
}

// NewStringValueItem returns new ValueItem from string.
func NewStringValueItem(v string) ValueItem {
	var s ValueItem
	s.SetString(v)
	return s
}

// Ref: #/components/schemas/Vector
type Vector struct {
	Result []VectorResultItem `json:"result"`
}

// GetResult returns the value of Result.
func (s *Vector) GetResult() []VectorResultItem {
	return s.Result
}

// SetResult sets the value of Result.
func (s *Vector) SetResult(val []VectorResultItem) {
	s.Result = val
}

type VectorResultItem struct {
	Metric VectorResultItemMetric `json:"metric"`
	Value  Value                  `json:"value"`
}

// GetMetric returns the value of Metric.
func (s *VectorResultItem) GetMetric() VectorResultItemMetric {
	return s.Metric
}

// GetValue returns the value of Value.
func (s *VectorResultItem) GetValue() Value {
	return s.Value
}

// SetMetric sets the value of Metric.
func (s *VectorResultItem) SetMetric(val VectorResultItemMetric) {
	s.Metric = val
}

// SetValue sets the value of Value.
func (s *VectorResultItem) SetValue(val Value) {
	s.Value = val
}

type VectorResultItemMetric map[string]string

func (s *VectorResultItemMetric) init() VectorResultItemMetric {
	m := *s
	if m == nil {
		m = map[string]string{}
		*s = m
	}
	return m
}