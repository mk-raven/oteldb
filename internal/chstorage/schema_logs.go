package chstorage

const (
	logsSchema = `CREATE TABLE IF NOT EXISTS %s
(
	timestamp DateTime64(9),
	timestamp_observed DateTime64(9),
	flags UInt32,
	severity_text String,
	severity_number Int32,
	body String,
	trace_id UUID,
	span_id UInt64,
	attributes String,
	resource String,
	scope_name String,
	scope_version String,
	scope_attributes String,
)
ENGINE = MergeTree()
ORDER BY (timestamp);`
)
