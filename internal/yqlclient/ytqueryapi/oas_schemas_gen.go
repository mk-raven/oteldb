// Code generated by ogen, DO NOT EDIT.

package ytqueryapi

import (
	"fmt"
	"io"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
)

func (s *ErrorStatusCode) Error() string {
	return fmt.Sprintf("code %d: %+v", s.StatusCode, s.Response)
}

// Ref: #/components/schemas/AbortedQuery
type AbortedQuery struct{}

// Ref: #/components/schemas/Attributes
type Attributes struct {
	Host         OptString   `json:"host"`
	Pid          OptUint64   `json:"pid"`
	Tid          OptUint64   `json:"tid"`
	Thread       OptString   `json:"thread"`
	Fid          OptUint64   `json:"fid"`
	Datetime     OptDateTime `json:"datetime"`
	TraceID      OptString   `json:"trace_id"`
	SpanID       OptUint64   `json:"span_id"`
	ConnectionID OptString   `json:"connection_id"`
	RealmID      OptString   `json:"realm_id"`
	Method       OptString   `json:"method"`
	RequestID    OptString   `json:"request_id"`
	Timeout      OptInt      `json:"timeout"`
	Address      OptString   `json:"address"`
	Service      OptString   `json:"service"`
}

// GetHost returns the value of Host.
func (s *Attributes) GetHost() OptString {
	return s.Host
}

// GetPid returns the value of Pid.
func (s *Attributes) GetPid() OptUint64 {
	return s.Pid
}

// GetTid returns the value of Tid.
func (s *Attributes) GetTid() OptUint64 {
	return s.Tid
}

// GetThread returns the value of Thread.
func (s *Attributes) GetThread() OptString {
	return s.Thread
}

// GetFid returns the value of Fid.
func (s *Attributes) GetFid() OptUint64 {
	return s.Fid
}

// GetDatetime returns the value of Datetime.
func (s *Attributes) GetDatetime() OptDateTime {
	return s.Datetime
}

// GetTraceID returns the value of TraceID.
func (s *Attributes) GetTraceID() OptString {
	return s.TraceID
}

// GetSpanID returns the value of SpanID.
func (s *Attributes) GetSpanID() OptUint64 {
	return s.SpanID
}

// GetConnectionID returns the value of ConnectionID.
func (s *Attributes) GetConnectionID() OptString {
	return s.ConnectionID
}

// GetRealmID returns the value of RealmID.
func (s *Attributes) GetRealmID() OptString {
	return s.RealmID
}

// GetMethod returns the value of Method.
func (s *Attributes) GetMethod() OptString {
	return s.Method
}

// GetRequestID returns the value of RequestID.
func (s *Attributes) GetRequestID() OptString {
	return s.RequestID
}

// GetTimeout returns the value of Timeout.
func (s *Attributes) GetTimeout() OptInt {
	return s.Timeout
}

// GetAddress returns the value of Address.
func (s *Attributes) GetAddress() OptString {
	return s.Address
}

// GetService returns the value of Service.
func (s *Attributes) GetService() OptString {
	return s.Service
}

// SetHost sets the value of Host.
func (s *Attributes) SetHost(val OptString) {
	s.Host = val
}

// SetPid sets the value of Pid.
func (s *Attributes) SetPid(val OptUint64) {
	s.Pid = val
}

// SetTid sets the value of Tid.
func (s *Attributes) SetTid(val OptUint64) {
	s.Tid = val
}

// SetThread sets the value of Thread.
func (s *Attributes) SetThread(val OptString) {
	s.Thread = val
}

// SetFid sets the value of Fid.
func (s *Attributes) SetFid(val OptUint64) {
	s.Fid = val
}

// SetDatetime sets the value of Datetime.
func (s *Attributes) SetDatetime(val OptDateTime) {
	s.Datetime = val
}

// SetTraceID sets the value of TraceID.
func (s *Attributes) SetTraceID(val OptString) {
	s.TraceID = val
}

// SetSpanID sets the value of SpanID.
func (s *Attributes) SetSpanID(val OptUint64) {
	s.SpanID = val
}

// SetConnectionID sets the value of ConnectionID.
func (s *Attributes) SetConnectionID(val OptString) {
	s.ConnectionID = val
}

// SetRealmID sets the value of RealmID.
func (s *Attributes) SetRealmID(val OptString) {
	s.RealmID = val
}

// SetMethod sets the value of Method.
func (s *Attributes) SetMethod(val OptString) {
	s.Method = val
}

// SetRequestID sets the value of RequestID.
func (s *Attributes) SetRequestID(val OptString) {
	s.RequestID = val
}

// SetTimeout sets the value of Timeout.
func (s *Attributes) SetTimeout(val OptInt) {
	s.Timeout = val
}

// SetAddress sets the value of Address.
func (s *Attributes) SetAddress(val OptString) {
	s.Address = val
}

// SetService sets the value of Service.
func (s *Attributes) SetService(val OptString) {
	s.Service = val
}

// Ref: #/components/schemas/Engine
type Engine string

const (
	EngineYql Engine = "yql"
	EngineQl  Engine = "ql"
)

// MarshalText implements encoding.TextMarshaler.
func (s Engine) MarshalText() ([]byte, error) {
	switch s {
	case EngineYql:
		return []byte(s), nil
	case EngineQl:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *Engine) UnmarshalText(data []byte) error {
	switch Engine(data) {
	case EngineYql:
		*s = EngineYql
		return nil
	case EngineQl:
		*s = EngineQl
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// Ref: #/components/schemas/Error
type Error struct {
	Code        int           `json:"code"`
	Message     string        `json:"message"`
	Attributes  OptAttributes `json:"attributes"`
	InnerErrors []Error       `json:"inner_errors"`
}

// GetCode returns the value of Code.
func (s *Error) GetCode() int {
	return s.Code
}

// GetMessage returns the value of Message.
func (s *Error) GetMessage() string {
	return s.Message
}

// GetAttributes returns the value of Attributes.
func (s *Error) GetAttributes() OptAttributes {
	return s.Attributes
}

// GetInnerErrors returns the value of InnerErrors.
func (s *Error) GetInnerErrors() []Error {
	return s.InnerErrors
}

// SetCode sets the value of Code.
func (s *Error) SetCode(val int) {
	s.Code = val
}

// SetMessage sets the value of Message.
func (s *Error) SetMessage(val string) {
	s.Message = val
}

// SetAttributes sets the value of Attributes.
func (s *Error) SetAttributes(val OptAttributes) {
	s.Attributes = val
}

// SetInnerErrors sets the value of InnerErrors.
func (s *Error) SetInnerErrors(val []Error) {
	s.InnerErrors = val
}

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

// Ref: #/components/schemas/OperationState
type OperationState string

const (
	OperationStateRunning       OperationState = "running"
	OperationStatePending       OperationState = "pending"
	OperationStateCompleted     OperationState = "completed"
	OperationStateFailed        OperationState = "failed"
	OperationStateAborted       OperationState = "aborted"
	OperationStateReviving      OperationState = "reviving"
	OperationStateInitializing  OperationState = "initializing"
	OperationStatePreparing     OperationState = "preparing"
	OperationStateMaterializing OperationState = "materializing"
	OperationStateCompleting    OperationState = "completing"
	OperationStateAborting      OperationState = "aborting"
	OperationStateFailing       OperationState = "failing"
)

// MarshalText implements encoding.TextMarshaler.
func (s OperationState) MarshalText() ([]byte, error) {
	switch s {
	case OperationStateRunning:
		return []byte(s), nil
	case OperationStatePending:
		return []byte(s), nil
	case OperationStateCompleted:
		return []byte(s), nil
	case OperationStateFailed:
		return []byte(s), nil
	case OperationStateAborted:
		return []byte(s), nil
	case OperationStateReviving:
		return []byte(s), nil
	case OperationStateInitializing:
		return []byte(s), nil
	case OperationStatePreparing:
		return []byte(s), nil
	case OperationStateMaterializing:
		return []byte(s), nil
	case OperationStateCompleting:
		return []byte(s), nil
	case OperationStateAborting:
		return []byte(s), nil
	case OperationStateFailing:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *OperationState) UnmarshalText(data []byte) error {
	switch OperationState(data) {
	case OperationStateRunning:
		*s = OperationStateRunning
		return nil
	case OperationStatePending:
		*s = OperationStatePending
		return nil
	case OperationStateCompleted:
		*s = OperationStateCompleted
		return nil
	case OperationStateFailed:
		*s = OperationStateFailed
		return nil
	case OperationStateAborted:
		*s = OperationStateAborted
		return nil
	case OperationStateReviving:
		*s = OperationStateReviving
		return nil
	case OperationStateInitializing:
		*s = OperationStateInitializing
		return nil
	case OperationStatePreparing:
		*s = OperationStatePreparing
		return nil
	case OperationStateMaterializing:
		*s = OperationStateMaterializing
		return nil
	case OperationStateCompleting:
		*s = OperationStateCompleting
		return nil
	case OperationStateAborting:
		*s = OperationStateAborting
		return nil
	case OperationStateFailing:
		*s = OperationStateFailing
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// NewOptAttributes returns new OptAttributes with value set to v.
func NewOptAttributes(v Attributes) OptAttributes {
	return OptAttributes{
		Value: v,
		Set:   true,
	}
}

// OptAttributes is optional Attributes.
type OptAttributes struct {
	Value Attributes
	Set   bool
}

// IsSet returns true if OptAttributes was set.
func (o OptAttributes) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptAttributes) Reset() {
	var v Attributes
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptAttributes) SetTo(v Attributes) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptAttributes) Get() (v Attributes, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptAttributes) Or(d Attributes) Attributes {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptDateTime returns new OptDateTime with value set to v.
func NewOptDateTime(v time.Time) OptDateTime {
	return OptDateTime{
		Value: v,
		Set:   true,
	}
}

// OptDateTime is optional time.Time.
type OptDateTime struct {
	Value time.Time
	Set   bool
}

// IsSet returns true if OptDateTime was set.
func (o OptDateTime) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptDateTime) Reset() {
	var v time.Time
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptDateTime) SetTo(v time.Time) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptDateTime) Get() (v time.Time, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptDateTime) Or(d time.Time) time.Time {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptError returns new OptError with value set to v.
func NewOptError(v Error) OptError {
	return OptError{
		Value: v,
		Set:   true,
	}
}

// OptError is optional Error.
type OptError struct {
	Value Error
	Set   bool
}

// IsSet returns true if OptError was set.
func (o OptError) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptError) Reset() {
	var v Error
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptError) SetTo(v Error) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptError) Get() (v Error, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptError) Or(d Error) Error {
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

// NewOptQueryStatusAnnotations returns new OptQueryStatusAnnotations with value set to v.
func NewOptQueryStatusAnnotations(v QueryStatusAnnotations) OptQueryStatusAnnotations {
	return OptQueryStatusAnnotations{
		Value: v,
		Set:   true,
	}
}

// OptQueryStatusAnnotations is optional QueryStatusAnnotations.
type OptQueryStatusAnnotations struct {
	Value QueryStatusAnnotations
	Set   bool
}

// IsSet returns true if OptQueryStatusAnnotations was set.
func (o OptQueryStatusAnnotations) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptQueryStatusAnnotations) Reset() {
	var v QueryStatusAnnotations
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptQueryStatusAnnotations) SetTo(v QueryStatusAnnotations) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptQueryStatusAnnotations) Get() (v QueryStatusAnnotations, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptQueryStatusAnnotations) Or(d QueryStatusAnnotations) QueryStatusAnnotations {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptQueryStatusProgress returns new OptQueryStatusProgress with value set to v.
func NewOptQueryStatusProgress(v QueryStatusProgress) OptQueryStatusProgress {
	return OptQueryStatusProgress{
		Value: v,
		Set:   true,
	}
}

// OptQueryStatusProgress is optional QueryStatusProgress.
type OptQueryStatusProgress struct {
	Value QueryStatusProgress
	Set   bool
}

// IsSet returns true if OptQueryStatusProgress was set.
func (o OptQueryStatusProgress) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptQueryStatusProgress) Reset() {
	var v QueryStatusProgress
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptQueryStatusProgress) SetTo(v QueryStatusProgress) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptQueryStatusProgress) Get() (v QueryStatusProgress, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptQueryStatusProgress) Or(d QueryStatusProgress) QueryStatusProgress {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptQueryStatusSettings returns new OptQueryStatusSettings with value set to v.
func NewOptQueryStatusSettings(v QueryStatusSettings) OptQueryStatusSettings {
	return OptQueryStatusSettings{
		Value: v,
		Set:   true,
	}
}

// OptQueryStatusSettings is optional QueryStatusSettings.
type OptQueryStatusSettings struct {
	Value QueryStatusSettings
	Set   bool
}

// IsSet returns true if OptQueryStatusSettings was set.
func (o OptQueryStatusSettings) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptQueryStatusSettings) Reset() {
	var v QueryStatusSettings
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptQueryStatusSettings) SetTo(v QueryStatusSettings) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptQueryStatusSettings) Get() (v QueryStatusSettings, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptQueryStatusSettings) Or(d QueryStatusSettings) QueryStatusSettings {
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

// NewOptUint64 returns new OptUint64 with value set to v.
func NewOptUint64(v uint64) OptUint64 {
	return OptUint64{
		Value: v,
		Set:   true,
	}
}

// OptUint64 is optional uint64.
type OptUint64 struct {
	Value uint64
	Set   bool
}

// IsSet returns true if OptUint64 was set.
func (o OptUint64) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptUint64) Reset() {
	var v uint64
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptUint64) SetTo(v uint64) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptUint64) Get() (v uint64, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptUint64) Or(d uint64) uint64 {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Ref: #/components/schemas/OutputFormat
type OutputFormat string

const (
	OutputFormatJSON         OutputFormat = "json"
	OutputFormatYson         OutputFormat = "yson"
	OutputFormatDsv          OutputFormat = "dsv"
	OutputFormatSchemafulDsv OutputFormat = "schemaful_dsv"
	OutputFormatProtobuf     OutputFormat = "protobuf"
)

// MarshalText implements encoding.TextMarshaler.
func (s OutputFormat) MarshalText() ([]byte, error) {
	switch s {
	case OutputFormatJSON:
		return []byte(s), nil
	case OutputFormatYson:
		return []byte(s), nil
	case OutputFormatDsv:
		return []byte(s), nil
	case OutputFormatSchemafulDsv:
		return []byte(s), nil
	case OutputFormatProtobuf:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *OutputFormat) UnmarshalText(data []byte) error {
	switch OutputFormat(data) {
	case OutputFormatJSON:
		*s = OutputFormatJSON
		return nil
	case OutputFormatYson:
		*s = OutputFormatYson
		return nil
	case OutputFormatDsv:
		*s = OutputFormatDsv
		return nil
	case OutputFormatSchemafulDsv:
		*s = OutputFormatSchemafulDsv
		return nil
	case OutputFormatProtobuf:
		*s = OutputFormatProtobuf
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

type QueryID string

type QueryResult struct {
	Data io.Reader
}

// Read reads data from the Data reader.
//
// Kept to satisfy the io.Reader interface.
func (s QueryResult) Read(p []byte) (n int, err error) {
	if s.Data == nil {
		return 0, io.EOF
	}
	return s.Data.Read(p)
}

// Note that error is could be present even if query is successful.
// Successful query error code is 0.
// Ref: #/components/schemas/QueryStatus
type QueryStatus struct {
	ID          QueryID                   `json:"id"`
	Engine      Engine                    `json:"engine"`
	StartTime   time.Time                 `json:"start_time"`
	FinishTime  OptDateTime               `json:"finish_time"`
	PingTime    OptDateTime               `json:"ping_time"`
	Settings    OptQueryStatusSettings    `json:"settings"`
	User        OptString                 `json:"user"`
	State       OperationState            `json:"state"`
	ResultCount OptInt                    `json:"result_count"`
	Progress    OptQueryStatusProgress    `json:"progress"`
	Annotations OptQueryStatusAnnotations `json:"annotations"`
	Incarnation OptInt                    `json:"incarnation"`
	Error       OptError                  `json:"error"`
}

// GetID returns the value of ID.
func (s *QueryStatus) GetID() QueryID {
	return s.ID
}

// GetEngine returns the value of Engine.
func (s *QueryStatus) GetEngine() Engine {
	return s.Engine
}

// GetStartTime returns the value of StartTime.
func (s *QueryStatus) GetStartTime() time.Time {
	return s.StartTime
}

// GetFinishTime returns the value of FinishTime.
func (s *QueryStatus) GetFinishTime() OptDateTime {
	return s.FinishTime
}

// GetPingTime returns the value of PingTime.
func (s *QueryStatus) GetPingTime() OptDateTime {
	return s.PingTime
}

// GetSettings returns the value of Settings.
func (s *QueryStatus) GetSettings() OptQueryStatusSettings {
	return s.Settings
}

// GetUser returns the value of User.
func (s *QueryStatus) GetUser() OptString {
	return s.User
}

// GetState returns the value of State.
func (s *QueryStatus) GetState() OperationState {
	return s.State
}

// GetResultCount returns the value of ResultCount.
func (s *QueryStatus) GetResultCount() OptInt {
	return s.ResultCount
}

// GetProgress returns the value of Progress.
func (s *QueryStatus) GetProgress() OptQueryStatusProgress {
	return s.Progress
}

// GetAnnotations returns the value of Annotations.
func (s *QueryStatus) GetAnnotations() OptQueryStatusAnnotations {
	return s.Annotations
}

// GetIncarnation returns the value of Incarnation.
func (s *QueryStatus) GetIncarnation() OptInt {
	return s.Incarnation
}

// GetError returns the value of Error.
func (s *QueryStatus) GetError() OptError {
	return s.Error
}

// SetID sets the value of ID.
func (s *QueryStatus) SetID(val QueryID) {
	s.ID = val
}

// SetEngine sets the value of Engine.
func (s *QueryStatus) SetEngine(val Engine) {
	s.Engine = val
}

// SetStartTime sets the value of StartTime.
func (s *QueryStatus) SetStartTime(val time.Time) {
	s.StartTime = val
}

// SetFinishTime sets the value of FinishTime.
func (s *QueryStatus) SetFinishTime(val OptDateTime) {
	s.FinishTime = val
}

// SetPingTime sets the value of PingTime.
func (s *QueryStatus) SetPingTime(val OptDateTime) {
	s.PingTime = val
}

// SetSettings sets the value of Settings.
func (s *QueryStatus) SetSettings(val OptQueryStatusSettings) {
	s.Settings = val
}

// SetUser sets the value of User.
func (s *QueryStatus) SetUser(val OptString) {
	s.User = val
}

// SetState sets the value of State.
func (s *QueryStatus) SetState(val OperationState) {
	s.State = val
}

// SetResultCount sets the value of ResultCount.
func (s *QueryStatus) SetResultCount(val OptInt) {
	s.ResultCount = val
}

// SetProgress sets the value of Progress.
func (s *QueryStatus) SetProgress(val OptQueryStatusProgress) {
	s.Progress = val
}

// SetAnnotations sets the value of Annotations.
func (s *QueryStatus) SetAnnotations(val OptQueryStatusAnnotations) {
	s.Annotations = val
}

// SetIncarnation sets the value of Incarnation.
func (s *QueryStatus) SetIncarnation(val OptInt) {
	s.Incarnation = val
}

// SetError sets the value of Error.
func (s *QueryStatus) SetError(val OptError) {
	s.Error = val
}

type QueryStatusAnnotations map[string]jx.Raw

func (s *QueryStatusAnnotations) init() QueryStatusAnnotations {
	m := *s
	if m == nil {
		m = map[string]jx.Raw{}
		*s = m
	}
	return m
}

type QueryStatusProgress map[string]jx.Raw

func (s *QueryStatusProgress) init() QueryStatusProgress {
	m := *s
	if m == nil {
		m = map[string]jx.Raw{}
		*s = m
	}
	return m
}

type QueryStatusSettings map[string]jx.Raw

func (s *QueryStatusSettings) init() QueryStatusSettings {
	m := *s
	if m == nil {
		m = map[string]jx.Raw{}
		*s = m
	}
	return m
}

// Ref: #/components/schemas/StartedQuery
type StartedQuery struct {
	QueryID QueryID `json:"query_id"`
}

// GetQueryID returns the value of QueryID.
func (s *StartedQuery) GetQueryID() QueryID {
	return s.QueryID
}

// SetQueryID sets the value of QueryID.
func (s *StartedQuery) SetQueryID(val QueryID) {
	s.QueryID = val
}

type YTToken struct {
	APIKey string
}

// GetAPIKey returns the value of APIKey.
func (s *YTToken) GetAPIKey() string {
	return s.APIKey
}

// SetAPIKey sets the value of APIKey.
func (s *YTToken) SetAPIKey(val string) {
	s.APIKey = val
}
