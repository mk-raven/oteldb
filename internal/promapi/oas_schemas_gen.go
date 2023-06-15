// Code generated by ogen, DO NOT EDIT.

package promapi

import (
	"fmt"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
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

// Ref: #/components/schemas/Exemplar
type Exemplar struct {
	Labels    OptLabelSet `json:"labels"`
	Value     OptFloat64  `json:"value"`
	Timestamp OptInt64    `json:"timestamp"`
}

// GetLabels returns the value of Labels.
func (s *Exemplar) GetLabels() OptLabelSet {
	return s.Labels
}

// GetValue returns the value of Value.
func (s *Exemplar) GetValue() OptFloat64 {
	return s.Value
}

// GetTimestamp returns the value of Timestamp.
func (s *Exemplar) GetTimestamp() OptInt64 {
	return s.Timestamp
}

// SetLabels sets the value of Labels.
func (s *Exemplar) SetLabels(val OptLabelSet) {
	s.Labels = val
}

// SetValue sets the value of Value.
func (s *Exemplar) SetValue(val OptFloat64) {
	s.Value = val
}

// SetTimestamp sets the value of Timestamp.
func (s *Exemplar) SetTimestamp(val OptInt64) {
	s.Timestamp = val
}

type Exemplars []ExemplarsSet

// Ref: #/components/schemas/ExemplarsSet
type ExemplarsSet struct {
	SeriesLabels OptLabelSet `json:"seriesLabels"`
	Exemplars    []Exemplar  `json:"exemplars"`
}

// GetSeriesLabels returns the value of SeriesLabels.
func (s *ExemplarsSet) GetSeriesLabels() OptLabelSet {
	return s.SeriesLabels
}

// GetExemplars returns the value of Exemplars.
func (s *ExemplarsSet) GetExemplars() []Exemplar {
	return s.Exemplars
}

// SetSeriesLabels sets the value of SeriesLabels.
func (s *ExemplarsSet) SetSeriesLabels(val OptLabelSet) {
	s.SeriesLabels = val
}

// SetExemplars sets the value of Exemplars.
func (s *ExemplarsSet) SetExemplars(val []Exemplar) {
	s.Exemplars = val
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

type GetRulesType string

const (
	GetRulesTypeAlert  GetRulesType = "alert"
	GetRulesTypeRecord GetRulesType = "record"
)

// MarshalText implements encoding.TextMarshaler.
func (s GetRulesType) MarshalText() ([]byte, error) {
	switch s {
	case GetRulesTypeAlert:
		return []byte(s), nil
	case GetRulesTypeRecord:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *GetRulesType) UnmarshalText(data []byte) error {
	switch GetRulesType(data) {
	case GetRulesTypeAlert:
		*s = GetRulesTypeAlert
		return nil
	case GetRulesTypeRecord:
		*s = GetRulesTypeRecord
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// Ref: #/components/schemas/LabelSet
type LabelSet map[string]string

func (s *LabelSet) init() LabelSet {
	m := *s
	if m == nil {
		m = map[string]string{}
		*s = m
	}
	return m
}

type LabelValues []string

type LabelValuesResponse struct {
	// Always 'success'.
	Status string `json:"status"`
	// Only if there were warnings while executing the request. There will still be data in the data
	// field.
	Warnings []string    `json:"warnings"`
	Data     LabelValues `json:"data"`
}

// GetStatus returns the value of Status.
func (s *LabelValuesResponse) GetStatus() string {
	return s.Status
}

// GetWarnings returns the value of Warnings.
func (s *LabelValuesResponse) GetWarnings() []string {
	return s.Warnings
}

// GetData returns the value of Data.
func (s *LabelValuesResponse) GetData() LabelValues {
	return s.Data
}

// SetStatus sets the value of Status.
func (s *LabelValuesResponse) SetStatus(val string) {
	s.Status = val
}

// SetWarnings sets the value of Warnings.
func (s *LabelValuesResponse) SetWarnings(val []string) {
	s.Warnings = val
}

// SetData sets the value of Data.
func (s *LabelValuesResponse) SetData(val LabelValues) {
	s.Data = val
}

type Labels []string

type LabelsResponse struct {
	// Always 'success'.
	Status string `json:"status"`
	// Only if there were warnings while executing the request. There will still be data in the data
	// field.
	Warnings []string `json:"warnings"`
	Data     Labels   `json:"data"`
}

// GetStatus returns the value of Status.
func (s *LabelsResponse) GetStatus() string {
	return s.Status
}

// GetWarnings returns the value of Warnings.
func (s *LabelsResponse) GetWarnings() []string {
	return s.Warnings
}

// GetData returns the value of Data.
func (s *LabelsResponse) GetData() Labels {
	return s.Data
}

// SetStatus sets the value of Status.
func (s *LabelsResponse) SetStatus(val string) {
	s.Status = val
}

// SetWarnings sets the value of Warnings.
func (s *LabelsResponse) SetWarnings(val []string) {
	s.Warnings = val
}

// SetData sets the value of Data.
func (s *LabelsResponse) SetData(val Labels) {
	s.Data = val
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

// Ref: #/components/schemas/Metadata
type Metadata map[string][]MetadataItemItem

func (s *Metadata) init() Metadata {
	m := *s
	if m == nil {
		m = map[string][]MetadataItemItem{}
		*s = m
	}
	return m
}

type MetadataItemItem struct {
	Type OptMetadataItemItemType `json:"type"`
	Help OptString               `json:"help"`
	Unit OptString               `json:"unit"`
}

// GetType returns the value of Type.
func (s *MetadataItemItem) GetType() OptMetadataItemItemType {
	return s.Type
}

// GetHelp returns the value of Help.
func (s *MetadataItemItem) GetHelp() OptString {
	return s.Help
}

// GetUnit returns the value of Unit.
func (s *MetadataItemItem) GetUnit() OptString {
	return s.Unit
}

// SetType sets the value of Type.
func (s *MetadataItemItem) SetType(val OptMetadataItemItemType) {
	s.Type = val
}

// SetHelp sets the value of Help.
func (s *MetadataItemItem) SetHelp(val OptString) {
	s.Help = val
}

// SetUnit sets the value of Unit.
func (s *MetadataItemItem) SetUnit(val OptString) {
	s.Unit = val
}

type MetadataItemItemType string

const (
	MetadataItemItemTypeCounter        MetadataItemItemType = "counter"
	MetadataItemItemTypeGauge          MetadataItemItemType = "gauge"
	MetadataItemItemTypeHistogram      MetadataItemItemType = "histogram"
	MetadataItemItemTypeGaugehistogram MetadataItemItemType = "gaugehistogram"
	MetadataItemItemTypeSummary        MetadataItemItemType = "summary"
	MetadataItemItemTypeInfo           MetadataItemItemType = "info"
	MetadataItemItemTypeStateset       MetadataItemItemType = "stateset"
	MetadataItemItemTypeUnknown        MetadataItemItemType = "unknown"
)

// MarshalText implements encoding.TextMarshaler.
func (s MetadataItemItemType) MarshalText() ([]byte, error) {
	switch s {
	case MetadataItemItemTypeCounter:
		return []byte(s), nil
	case MetadataItemItemTypeGauge:
		return []byte(s), nil
	case MetadataItemItemTypeHistogram:
		return []byte(s), nil
	case MetadataItemItemTypeGaugehistogram:
		return []byte(s), nil
	case MetadataItemItemTypeSummary:
		return []byte(s), nil
	case MetadataItemItemTypeInfo:
		return []byte(s), nil
	case MetadataItemItemTypeStateset:
		return []byte(s), nil
	case MetadataItemItemTypeUnknown:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *MetadataItemItemType) UnmarshalText(data []byte) error {
	switch MetadataItemItemType(data) {
	case MetadataItemItemTypeCounter:
		*s = MetadataItemItemTypeCounter
		return nil
	case MetadataItemItemTypeGauge:
		*s = MetadataItemItemTypeGauge
		return nil
	case MetadataItemItemTypeHistogram:
		*s = MetadataItemItemTypeHistogram
		return nil
	case MetadataItemItemTypeGaugehistogram:
		*s = MetadataItemItemTypeGaugehistogram
		return nil
	case MetadataItemItemTypeSummary:
		*s = MetadataItemItemTypeSummary
		return nil
	case MetadataItemItemTypeInfo:
		*s = MetadataItemItemTypeInfo
		return nil
	case MetadataItemItemTypeStateset:
		*s = MetadataItemItemTypeStateset
		return nil
	case MetadataItemItemTypeUnknown:
		*s = MetadataItemItemTypeUnknown
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

type MetadataResponse struct {
	// Always 'success'.
	Status string `json:"status"`
	// Only if there were warnings while executing the request. There will still be data in the data
	// field.
	Warnings []string `json:"warnings"`
	Data     Metadata `json:"data"`
}

// GetStatus returns the value of Status.
func (s *MetadataResponse) GetStatus() string {
	return s.Status
}

// GetWarnings returns the value of Warnings.
func (s *MetadataResponse) GetWarnings() []string {
	return s.Warnings
}

// GetData returns the value of Data.
func (s *MetadataResponse) GetData() Metadata {
	return s.Data
}

// SetStatus sets the value of Status.
func (s *MetadataResponse) SetStatus(val string) {
	s.Status = val
}

// SetWarnings sets the value of Warnings.
func (s *MetadataResponse) SetWarnings(val []string) {
	s.Warnings = val
}

// SetData sets the value of Data.
func (s *MetadataResponse) SetData(val Metadata) {
	s.Data = val
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

// NewOptFloat64 returns new OptFloat64 with value set to v.
func NewOptFloat64(v float64) OptFloat64 {
	return OptFloat64{
		Value: v,
		Set:   true,
	}
}

// OptFloat64 is optional float64.
type OptFloat64 struct {
	Value float64
	Set   bool
}

// IsSet returns true if OptFloat64 was set.
func (o OptFloat64) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptFloat64) Reset() {
	var v float64
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptFloat64) SetTo(v float64) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptFloat64) Get() (v float64, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptFloat64) Or(d float64) float64 {
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

// NewOptInt64 returns new OptInt64 with value set to v.
func NewOptInt64(v int64) OptInt64 {
	return OptInt64{
		Value: v,
		Set:   true,
	}
}

// OptInt64 is optional int64.
type OptInt64 struct {
	Value int64
	Set   bool
}

// IsSet returns true if OptInt64 was set.
func (o OptInt64) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt64) Reset() {
	var v int64
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt64) SetTo(v int64) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt64) Get() (v int64, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt64) Or(d int64) int64 {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptLabelSet returns new OptLabelSet with value set to v.
func NewOptLabelSet(v LabelSet) OptLabelSet {
	return OptLabelSet{
		Value: v,
		Set:   true,
	}
}

// OptLabelSet is optional LabelSet.
type OptLabelSet struct {
	Value LabelSet
	Set   bool
}

// IsSet returns true if OptLabelSet was set.
func (o OptLabelSet) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptLabelSet) Reset() {
	var v LabelSet
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptLabelSet) SetTo(v LabelSet) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptLabelSet) Get() (v LabelSet, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptLabelSet) Or(d LabelSet) LabelSet {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptMetadataItemItemType returns new OptMetadataItemItemType with value set to v.
func NewOptMetadataItemItemType(v MetadataItemItemType) OptMetadataItemItemType {
	return OptMetadataItemItemType{
		Value: v,
		Set:   true,
	}
}

// OptMetadataItemItemType is optional MetadataItemItemType.
type OptMetadataItemItemType struct {
	Value MetadataItemItemType
	Set   bool
}

// IsSet returns true if OptMetadataItemItemType was set.
func (o OptMetadataItemItemType) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptMetadataItemItemType) Reset() {
	var v MetadataItemItemType
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptMetadataItemItemType) SetTo(v MetadataItemItemType) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptMetadataItemItemType) Get() (v MetadataItemItemType, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptMetadataItemItemType) Or(d MetadataItemItemType) MetadataItemItemType {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptPrometheusTimestamp returns new OptPrometheusTimestamp with value set to v.
func NewOptPrometheusTimestamp(v PrometheusTimestamp) OptPrometheusTimestamp {
	return OptPrometheusTimestamp{
		Value: v,
		Set:   true,
	}
}

// OptPrometheusTimestamp is optional PrometheusTimestamp.
type OptPrometheusTimestamp struct {
	Value PrometheusTimestamp
	Set   bool
}

// IsSet returns true if OptPrometheusTimestamp was set.
func (o OptPrometheusTimestamp) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptPrometheusTimestamp) Reset() {
	var v PrometheusTimestamp
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptPrometheusTimestamp) SetTo(v PrometheusTimestamp) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptPrometheusTimestamp) Get() (v PrometheusTimestamp, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptPrometheusTimestamp) Or(d PrometheusTimestamp) PrometheusTimestamp {
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

type PrometheusDuration string

type PrometheusTimestamp string

type QueryExemplarsResponse struct {
	// Always 'success'.
	Status string `json:"status"`
	// Only if there were warnings while executing the request. There will still be data in the data
	// field.
	Warnings []string  `json:"warnings"`
	Data     Exemplars `json:"data"`
}

// GetStatus returns the value of Status.
func (s *QueryExemplarsResponse) GetStatus() string {
	return s.Status
}

// GetWarnings returns the value of Warnings.
func (s *QueryExemplarsResponse) GetWarnings() []string {
	return s.Warnings
}

// GetData returns the value of Data.
func (s *QueryExemplarsResponse) GetData() Exemplars {
	return s.Data
}

// SetStatus sets the value of Status.
func (s *QueryExemplarsResponse) SetStatus(val string) {
	s.Status = val
}

// SetWarnings sets the value of Warnings.
func (s *QueryExemplarsResponse) SetWarnings(val []string) {
	s.Warnings = val
}

// SetData sets the value of Data.
func (s *QueryExemplarsResponse) SetData(val Exemplars) {
	s.Data = val
}

type QueryResponse struct {
	// Always 'success'.
	Status string `json:"status"`
	// Only if there were warnings while executing the request. There will still be data in the data
	// field.
	Warnings []string `json:"warnings"`
	Data     Data     `json:"data"`
}

// GetStatus returns the value of Status.
func (s *QueryResponse) GetStatus() string {
	return s.Status
}

// GetWarnings returns the value of Warnings.
func (s *QueryResponse) GetWarnings() []string {
	return s.Warnings
}

// GetData returns the value of Data.
func (s *QueryResponse) GetData() Data {
	return s.Data
}

// SetStatus sets the value of Status.
func (s *QueryResponse) SetStatus(val string) {
	s.Status = val
}

// SetWarnings sets the value of Warnings.
func (s *QueryResponse) SetWarnings(val []string) {
	s.Warnings = val
}

// SetData sets the value of Data.
func (s *QueryResponse) SetData(val Data) {
	s.Data = val
}

type Rule jx.Raw

// Ref: #/components/schemas/RuleGroup
type RuleGroup struct {
	Name OptString `json:"name"`
	File OptString `json:"file"`
	// In order to preserve rule ordering, while exposing type (alerting or recording)
	// specific properties, both alerting and recording rules are exposed in the
	// same array.
	Rules          []Rule     `json:"rules"`
	Internal       OptFloat64 `json:"internal"`
	Limit          OptInt     `json:"limit"`
	EvaluationTime OptFloat64 `json:"evaluationTime"`
	LastEvaluation OptString  `json:"lastEvaluation"`
}

// GetName returns the value of Name.
func (s *RuleGroup) GetName() OptString {
	return s.Name
}

// GetFile returns the value of File.
func (s *RuleGroup) GetFile() OptString {
	return s.File
}

// GetRules returns the value of Rules.
func (s *RuleGroup) GetRules() []Rule {
	return s.Rules
}

// GetInternal returns the value of Internal.
func (s *RuleGroup) GetInternal() OptFloat64 {
	return s.Internal
}

// GetLimit returns the value of Limit.
func (s *RuleGroup) GetLimit() OptInt {
	return s.Limit
}

// GetEvaluationTime returns the value of EvaluationTime.
func (s *RuleGroup) GetEvaluationTime() OptFloat64 {
	return s.EvaluationTime
}

// GetLastEvaluation returns the value of LastEvaluation.
func (s *RuleGroup) GetLastEvaluation() OptString {
	return s.LastEvaluation
}

// SetName sets the value of Name.
func (s *RuleGroup) SetName(val OptString) {
	s.Name = val
}

// SetFile sets the value of File.
func (s *RuleGroup) SetFile(val OptString) {
	s.File = val
}

// SetRules sets the value of Rules.
func (s *RuleGroup) SetRules(val []Rule) {
	s.Rules = val
}

// SetInternal sets the value of Internal.
func (s *RuleGroup) SetInternal(val OptFloat64) {
	s.Internal = val
}

// SetLimit sets the value of Limit.
func (s *RuleGroup) SetLimit(val OptInt) {
	s.Limit = val
}

// SetEvaluationTime sets the value of EvaluationTime.
func (s *RuleGroup) SetEvaluationTime(val OptFloat64) {
	s.EvaluationTime = val
}

// SetLastEvaluation sets the value of LastEvaluation.
func (s *RuleGroup) SetLastEvaluation(val OptString) {
	s.LastEvaluation = val
}

// Ref: #/components/schemas/Rules
type Rules struct {
	Groups RuleGroup `json:"groups"`
}

// GetGroups returns the value of Groups.
func (s *Rules) GetGroups() RuleGroup {
	return s.Groups
}

// SetGroups sets the value of Groups.
func (s *Rules) SetGroups(val RuleGroup) {
	s.Groups = val
}

type RulesResponse struct {
	// Always 'success'.
	Status string `json:"status"`
	// Only if there were warnings while executing the request. There will still be data in the data
	// field.
	Warnings []string `json:"warnings"`
	Data     Rules    `json:"data"`
}

// GetStatus returns the value of Status.
func (s *RulesResponse) GetStatus() string {
	return s.Status
}

// GetWarnings returns the value of Warnings.
func (s *RulesResponse) GetWarnings() []string {
	return s.Warnings
}

// GetData returns the value of Data.
func (s *RulesResponse) GetData() Rules {
	return s.Data
}

// SetStatus sets the value of Status.
func (s *RulesResponse) SetStatus(val string) {
	s.Status = val
}

// SetWarnings sets the value of Warnings.
func (s *RulesResponse) SetWarnings(val []string) {
	s.Warnings = val
}

// SetData sets the value of Data.
func (s *RulesResponse) SetData(val Rules) {
	s.Data = val
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

// Ref: #/components/schemas/Value
type Value struct {
	T float64
	V string
}

// GetT returns the value of T.
func (s *Value) GetT() float64 {
	return s.T
}

// GetV returns the value of V.
func (s *Value) GetV() string {
	return s.V
}

// SetT sets the value of T.
func (s *Value) SetT(val float64) {
	s.T = val
}

// SetV sets the value of V.
func (s *Value) SetV(val string) {
	s.V = val
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
