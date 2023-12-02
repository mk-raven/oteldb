// Package zapotel provides OpenTelemetry logs exporter zap core implementation.
package zapotel

import (
	"context"
	"encoding/hex"
	"math"
	"strings"

	"github.com/go-faster/errors"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/plog/plogotlp"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.uber.org/zap/zapcore"
)

// New initializes new zapcore.Core from grpc client and resource.
func New(enab zapcore.LevelEnabler, res *resource.Resource, client plogotlp.GRPCClient) zapcore.Core {
	return &exporter{
		LevelEnabler: enab,
		client:       client,
		res:          res,
	}
}

type exporter struct {
	zapcore.LevelEnabler
	context []zapcore.Field
	client  plogotlp.GRPCClient
	res     *resource.Resource
}

var (
	_ zapcore.Core         = (*exporter)(nil)
	_ zapcore.LevelEnabler = (*exporter)(nil)
)

func (e *exporter) Level() zapcore.Level {
	return zapcore.LevelOf(e.LevelEnabler)
}

func (e *exporter) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if e.Enabled(ent.Level) {
		return ce.AddCore(ent, e)
	}
	return ce
}

func (e *exporter) With(fields []zapcore.Field) zapcore.Core {
	return &exporter{
		LevelEnabler: e.LevelEnabler,
		client:       e.client,
		res:          e.res,
		context:      append(e.context[:len(e.context):len(e.context)], fields...),
	}
}

func (e *exporter) toLogs(ent zapcore.Entry, fields []zapcore.Field) plog.Logs {
	var (
		ld = plog.NewLogs()
		rl = ld.ResourceLogs().AppendEmpty()
	)
	{
		a := rl.Resource().Attributes()
		for _, kv := range e.res.Attributes() {
			k := string(kv.Key)
			switch kv.Value.Type() {
			case attribute.STRING:
				a.PutStr(k, kv.Value.AsString())
			case attribute.BOOL:
				a.PutBool(k, kv.Value.AsBool())
			default:
				a.PutStr(k, kv.Value.AsString())
			}
		}
	}

	il := rl.ScopeLogs().AppendEmpty()

	scope := il.Scope()
	scope.SetName("zapotel")
	scope.SetVersion("v0.1")

	lg := il.LogRecords().AppendEmpty()
	lg.Body().SetStr(ent.Message)
	// TODO: update mapping from spec
	switch ent.Level {
	case zapcore.DebugLevel:
		lg.SetSeverityNumber(plog.SeverityNumberDebug)
	case zapcore.InfoLevel:
		lg.SetSeverityNumber(plog.SeverityNumberInfo)
	case zapcore.WarnLevel:
		lg.SetSeverityNumber(plog.SeverityNumberWarn)
	case zapcore.ErrorLevel:
		lg.SetSeverityNumber(plog.SeverityNumberError)
	case zapcore.DPanicLevel:
		lg.SetSeverityNumber(plog.SeverityNumberFatal)
	case zapcore.PanicLevel:
		lg.SetSeverityNumber(plog.SeverityNumberFatal)
	case zapcore.FatalLevel:
		lg.SetSeverityNumber(plog.SeverityNumberFatal)
	}
	lg.SetSeverityText(ent.Level.String())
	lg.SetTimestamp(pcommon.NewTimestampFromTime(ent.Time))
	lg.SetObservedTimestamp(pcommon.NewTimestampFromTime(ent.Time))
	{
		a := lg.Attributes()
		var skipped uint32
		for _, f := range fields {
			k := f.Key
			switch f.Type {
			case zapcore.BoolType:
				a.PutBool(k, f.Integer == 1)
			case zapcore.StringType:
				l := len(f.String)
				if (k == "trace_id" && l == 32) || (k == "span_id" && l == 16) {
					// Checking for tracing.
					var (
						traceID pcommon.TraceID
						spanID  pcommon.SpanID
					)
					v, err := hex.DecodeString(strings.ToLower(f.String))
					if err == nil {
						switch k {
						case "trace_id":
							copy(traceID[:], v)
							lg.SetTraceID(traceID)
						case "span_id":
							copy(spanID[:], v)
							lg.SetSpanID(spanID)
						}
						// Don't add as regular string.
						continue
					}
				}
				a.PutStr(k, f.String)
			case zapcore.Int8Type, zapcore.Int16Type, zapcore.Int32Type, zapcore.Int64Type,
				zapcore.Uint8Type, zapcore.Uint16Type, zapcore.Uint32Type, zapcore.Uint64Type:
				a.PutInt(k, f.Integer)
			case zapcore.Float32Type:
				a.PutDouble(k, float64(math.Float32frombits(uint32(f.Integer))))
			case zapcore.Float64Type:
				a.PutDouble(k, math.Float64frombits(uint64(f.Integer)))
			default:
				// Time, duration, "any", ...
				// TODO(ernado): support
				skipped++
			}
		}
		if skipped > 0 {
			scope.SetDroppedAttributesCount(skipped)
		}
	}
	return ld
}

func (e *exporter) Write(ent zapcore.Entry, fields []zapcore.Field) error {
	all := make([]zapcore.Field, 0, len(fields)+len(e.context))
	all = append(all, e.context...)
	all = append(all, fields...)

	ctx := context.TODO()
	logs := e.toLogs(ent, all)
	req := plogotlp.NewExportRequestFromLogs(logs)
	if _, err := e.client.Export(ctx, req); err != nil {
		return errors.Wrap(err, "send logs")
	}

	return nil
}

func (e *exporter) Sync() error { return nil }