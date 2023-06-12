// Package tracestorage defines storage structure for trace storage.
package tracestorage

import (
	"context"
	"time"
)

// Querier is a trace storage query interface.
type Querier interface {
	// SearchTags performs search by given tags.
	SearchTags(ctx context.Context, tags map[string]string, opts SearchTagsOptions) (Iterator[Span], error)

	// TagNames returns all available tag names.
	TagNames(ctx context.Context) ([]string, error)
	// TagValues returns all available tag values for given tag.
	TagValues(ctx context.Context, tagName string) (Iterator[Tag], error)

	// TraceByID returns spans of given trace.
	TraceByID(ctx context.Context, id TraceID, opts TraceByIDOptions) (Iterator[Span], error)
}

// SearchTagsOptions defines options for SearchTags method.
type SearchTagsOptions struct {
	MinDuration time.Duration
	MaxDuration time.Duration

	// Start defines time range for search.
	//
	// Querier ignores parameter, if it is zero.
	Start Timestamp
	// End defines time range for search.
	//
	// Querier ignores parameter, if it is zero.
	End Timestamp
}

// TraceByIDOptions defines options for TraceByID method.
type TraceByIDOptions struct {
	// Start defines time range for search.
	//
	// Querier ignores parameter, if it is zero.
	Start Timestamp
	// End defines time range for search.
	//
	// Querier ignores parameter, if it is zero.
	End Timestamp
}

// Inserter is a trace storage insert interface.
type Inserter interface {
	// InsertSpans inserts given spans.
	//
	// FIXME(tdakkota): probably, it's better to return some kind of batch writer.
	InsertSpans(ctx context.Context, spans []Span) error
	// InsertTags insert given set of tags to the storage.
	//
	// FIXME(tdakkota): probably, storage should do tag extraction by itself.
	InsertTags(ctx context.Context, tags map[Tag]struct{}) error
}