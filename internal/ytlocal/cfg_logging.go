package ytlocal

// LogLevel string describes possible logging levels.
type LogLevel string

const (
	LogLevelTrace LogLevel = "trace"
	LogLevelDebug LogLevel = "debug"
	LogLevelInfo  LogLevel = "info"
	LogLevelError LogLevel = "error"
)

// LogWriterType string describes types of possible log writers.
type LogWriterType string

// Possible log writer types.
const (
	LogWriterTypeFile   LogWriterType = "file"
	LogWriterTypeStderr LogWriterType = "stderr"
)

// LoggingRule configures logging.
type LoggingRule struct {
	ExcludeCategories []string `yson:"exclude_categories,omitempty"`
	IncludeCategories []string `yson:"include_categories,omitempty"`
	MinLevel          LogLevel `yson:"min_level,omitempty"`
	Writers           []string `yson:"writers,omitempty"`
}

// LogFormat is a string describing possible log formats.
type LogFormat string

// Possible log formats.
const (
	LogFormatPlainText LogFormat = "plaintext"
	LogFormatJSON      LogFormat = "json"
	LogFormatYSON      LogFormat = "yson"
)

// LoggingWriter configures logging writer.
type LoggingWriter struct {
	WriterType LogWriterType `yson:"type,omitempty"`
	Format     LogFormat     `yson:"format,omitempty"`
	FileName   string        `yson:"file_name,omitempty"`
}

// Logging config.
type Logging struct {
	Writers map[string]LoggingWriter `yson:"writers,omitempty"`
	Rules   []LoggingRule            `yson:"rules,omitempty"`
}
