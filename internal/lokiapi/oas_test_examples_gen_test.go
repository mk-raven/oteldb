// Code generated by ogen, DO NOT EDIT.

package lokiapi

import (
	"github.com/go-faster/jx"

	std "encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestError_EncodeDecode(t *testing.T) {
	var typ Error
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Error
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestLabelSet_EncodeDecode(t *testing.T) {
	var typ LabelSet
	typ = make(LabelSet)
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 LabelSet
	typ2 = make(LabelSet)
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestLabels_EncodeDecode(t *testing.T) {
	var typ Labels
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Labels
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestLogEntry_EncodeDecode(t *testing.T) {
	var typ LogEntry
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 LogEntry
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestMaps_EncodeDecode(t *testing.T) {
	var typ Maps
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Maps
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestMapsDataItem_EncodeDecode(t *testing.T) {
	var typ MapsDataItem
	typ = make(MapsDataItem)
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 MapsDataItem
	typ2 = make(MapsDataItem)
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestMatrix_EncodeDecode(t *testing.T) {
	var typ Matrix
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Matrix
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestMatrixResult_EncodeDecode(t *testing.T) {
	var typ MatrixResult
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 MatrixResult
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestPrometheusSamplePair_EncodeDecode(t *testing.T) {
	var typ PrometheusSamplePair
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 PrometheusSamplePair
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestPush_EncodeDecode(t *testing.T) {
	var typ Push
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Push
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestQueryResponse_EncodeDecode(t *testing.T) {
	var typ QueryResponse
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 QueryResponse
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestQueryResponseData_EncodeDecode(t *testing.T) {
	var typ QueryResponseData
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 QueryResponseData
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestScalarResult_EncodeDecode(t *testing.T) {
	var typ ScalarResult
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 ScalarResult
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestSeries_EncodeDecode(t *testing.T) {
	var typ Series
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Series
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestStats_EncodeDecode(t *testing.T) {
	var typ Stats
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Stats
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestStream_EncodeDecode(t *testing.T) {
	var typ Stream
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Stream
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestStreams_EncodeDecode(t *testing.T) {
	var typ Streams
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Streams
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestStreamsResult_EncodeDecode(t *testing.T) {
	var typ StreamsResult
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 StreamsResult
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestValues_EncodeDecode(t *testing.T) {
	var typ Values
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Values
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestVector_EncodeDecode(t *testing.T) {
	var typ Vector
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Vector
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestVectorResult_EncodeDecode(t *testing.T) {
	var typ VectorResult
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 VectorResult
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
