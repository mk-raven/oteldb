package traceqlengine

import (
	"fmt"
	"testing"

	"golang.org/x/exp/slices"

	"github.com/go-faster/oteldb/internal/otelstorage"
	"github.com/go-faster/oteldb/internal/tracestorage"

	"github.com/stretchr/testify/require"
)

type spanIDs struct {
	id     uint64
	parent uint64
}

func generateSpans(ids []spanIDs) []tracestorage.Span {
	result := make([]tracestorage.Span, 0, len(ids))
	for _, span := range ids {
		result = append(result, tracestorage.Span{
			SpanID:       otelstorage.SpanIDFromUint64(span.id),
			ParentSpanID: otelstorage.SpanIDFromUint64(span.parent),
			Name:         fmt.Sprintf("Span #%d", span.id),
		})
	}
	return result
}

func TestChildSpans(t *testing.T) {
	tests := []struct {
		left, right []spanIDs
		wantResult  []uint64
	}{
		{
			nil,
			nil,
			nil,
		},
		{
			[]spanIDs{
				{id: 1},
				{id: 2},
				{id: 3},
			},
			[]spanIDs{
				{id: 1},
				{id: 4, parent: 2},
				{id: 5, parent: 3},
				{id: 6, parent: 4},
				{id: 7, parent: 3},
			},
			[]uint64{
				4, 5, 7,
			},
		},
		{
			[]spanIDs{
				{id: 1},
				{id: 2},
				{id: 3},
			},
			[]spanIDs{
				{id: 4, parent: 11},
				{id: 5, parent: 12},
				{id: 6, parent: 13},
			},
			nil,
		},
	}
	for i, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("Test%d", i+1), func(t *testing.T) {
			result := childSpans(
				Spanset{
					Spans: generateSpans(tt.left),
				},
				Spanset{
					Spans: generateSpans(tt.right),
				},
			)

			var got []uint64
			for _, span := range result {
				got = append(got, span.SpanID.AsUint64())
			}
			slices.Sort(got)
			require.Equal(t, tt.wantResult, got)
		})
	}
}

func TestSiblingSpans(t *testing.T) {
	tests := []struct {
		left, right []spanIDs
		wantResult  []uint64
	}{
		{
			nil,
			nil,
			nil,
		},
		{
			[]spanIDs{
				{id: 1, parent: 11},
				{id: 2, parent: 12},
				{id: 3, parent: 13},
				{id: 10},
			},
			[]spanIDs{
				{id: 4, parent: 11},
				{id: 5, parent: 12},
				{id: 6, parent: 13},
				{id: 7, parent: 14},
				{id: 8, parent: 11},
				{id: 9}, // matches id = 10, since they both have no parent.
			},
			[]uint64{
				4, 5, 6, 8, 9,
			},
		},
		// Have no common parents.
		{
			[]spanIDs{
				{id: 1, parent: 11},
				{id: 2, parent: 12},
				{id: 3, parent: 13},
			},
			[]spanIDs{
				{id: 4, parent: 21},
				{id: 5, parent: 22},
				{id: 6, parent: 23},
			},
			nil,
		},
		// Parent spans are not sibling spans.
		{
			[]spanIDs{
				{id: 1, parent: 11},
				{id: 2, parent: 12},
				{id: 3, parent: 13},
			},
			[]spanIDs{
				{id: 4, parent: 1},
				{id: 5, parent: 2},
				{id: 6, parent: 3},
			},
			nil,
		},
		// Do not match root span and non-root span.
		{
			[]spanIDs{
				{id: 1},
			},
			[]spanIDs{
				{id: 2, parent: 11},
			},
			nil,
		},
		// Duplcate spans.
		{
			[]spanIDs{
				{id: 1, parent: 11},
				{id: 2, parent: 12},
				{id: 3, parent: 13},
			},
			[]spanIDs{
				{id: 2, parent: 12},
				{id: 4, parent: 13},
				{id: 2, parent: 12},
				{id: 5, parent: 14},
			},
			[]uint64{2, 4},
		},
	}
	for i, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("Test%d", i+1), func(t *testing.T) {
			result := siblingSpans(
				Spanset{
					Spans: generateSpans(tt.left),
				},
				Spanset{
					Spans: generateSpans(tt.right),
				},
			)

			var got []uint64
			for _, span := range result {
				got = append(got, span.SpanID.AsUint64())
			}
			slices.Sort(got)
			require.Equal(t, tt.wantResult, got)
		})
	}
}
