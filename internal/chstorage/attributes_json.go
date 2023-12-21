package chstorage

import (
	"github.com/go-faster/jx"
	"go.opentelemetry.io/collector/pdata/pcommon"

	"github.com/go-faster/oteldb/internal/otelstorage"
)

type lazyAttributes struct {
	orig pcommon.Map

	encoded string
}

func (l *lazyAttributes) Encode(additional ...[2]string) string {
	switch {
	case len(additional) > 0:
		return encodeAttributes(l.orig, additional...)
	case l.encoded != "":
		return l.encoded
	default:
		l.encoded = encodeAttributes(l.orig)
		return l.encoded
	}
}

func encodeAttributes(attrs pcommon.Map, additional ...[2]string) string {
	e := jx.GetEncoder()
	defer jx.PutEncoder(e)

	encodeMap(e, attrs, additional...)
	return e.String()
}

func encodeValue(e *jx.Encoder, v pcommon.Value) {
	switch v.Type() {
	case pcommon.ValueTypeStr:
		e.Str(v.Str())
	case pcommon.ValueTypeInt:
		e.Int64(v.Int())
	case pcommon.ValueTypeDouble:
		e.Float64(v.Double())
	case pcommon.ValueTypeBool:
		e.Bool(v.Bool())
	case pcommon.ValueTypeMap:
		m := v.Map()
		encodeMap(e, m)
	case pcommon.ValueTypeSlice:
		s := v.Slice()
		e.ArrStart()
		for i := 0; i < s.Len(); i++ {
			encodeValue(e, s.At(i))
		}
		e.ArrEnd()
	case pcommon.ValueTypeBytes:
		e.ByteStr(v.Bytes().AsRaw())
	default:
		e.Null()
	}
}

func encodeMap(e *jx.Encoder, m pcommon.Map, additional ...[2]string) {
	if otelstorage.Attrs(m).IsZero() && len(additional) == 0 {
		e.ObjEmpty()
		return
	}
	e.ObjStart()
	m.Range(func(k string, v pcommon.Value) bool {
		e.FieldStart(k)
		encodeValue(e, v)
		return true
	})
	for _, pair := range additional {
		e.FieldStart(pair[0])
		e.Str(pair[1])
	}
	e.ObjEnd()
}

func decodeAttributes(s string) (otelstorage.Attrs, error) {
	result := pcommon.NewMap()
	err := decodeMap(jx.DecodeStr(s), result)
	return otelstorage.Attrs(result), err
}

func decodeValue(d *jx.Decoder, val pcommon.Value) error {
	switch d.Next() {
	case jx.String:
		v, err := d.Str()
		if err != nil {
			return err
		}
		val.SetStr(v)
		return nil
	case jx.Number:
		n, err := d.Num()
		if err != nil {
			return err
		}
		if n.IsInt() {
			v, err := n.Int64()
			if err != nil {
				return err
			}
			val.SetInt(v)
		} else {
			v, err := n.Float64()
			if err != nil {
				return err
			}
			val.SetDouble(v)
		}
		return nil
	case jx.Null:
		if err := d.Null(); err != nil {
			return err
		}
		// Do nothing, keep value empty.
		return nil
	case jx.Bool:
		v, err := d.Bool()
		if err != nil {
			return err
		}
		val.SetBool(v)
		return nil
	case jx.Array:
		s := val.SetEmptySlice()
		return d.Arr(func(d *jx.Decoder) error {
			rval := s.AppendEmpty()
			return decodeValue(d, rval)
		})
	case jx.Object:
		m := val.SetEmptyMap()
		return decodeMap(d, m)
	default:
		return d.Skip()
	}
}

func decodeMap(d *jx.Decoder, m pcommon.Map) error {
	return d.Obj(func(d *jx.Decoder, key string) error {
		rval := m.PutEmpty(key)
		return decodeValue(d, rval)
	})
}