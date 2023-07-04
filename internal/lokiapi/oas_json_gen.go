// Code generated by ogen, DO NOT EDIT.

package lokiapi

import (
	"math/bits"
	"strconv"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/validate"
)

// Encode encodes Entry as json.
func (s Entry) Encode(e *jx.Encoder) {
	unwrapped := []Value(s)

	e.ArrStart()
	for _, elem := range unwrapped {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes Entry from json.
func (s *Entry) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Entry to nil")
	}
	var unwrapped []Value
	if err := func() error {
		unwrapped = make([]Value, 0)
		if err := d.Arr(func(d *jx.Decoder) error {
			var elem Value
			if err := elem.Decode(d); err != nil {
				return err
			}
			unwrapped = append(unwrapped, elem)
			return nil
		}); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = Entry(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s Entry) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Entry) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes Error as json.
func (s Error) Encode(e *jx.Encoder) {
	unwrapped := string(s)

	e.Str(unwrapped)
}

// Decode decodes Error from json.
func (s *Error) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Error to nil")
	}
	var unwrapped string
	if err := func() error {
		v, err := d.Str()
		unwrapped = string(v)
		if err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = Error(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s Error) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Error) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s LabelSet) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields implements json.Marshaler.
func (s LabelSet) encodeFields(e *jx.Encoder) {
	for k, elem := range s {
		e.FieldStart(k)

		e.Str(elem)
	}
}

// Decode decodes LabelSet from json.
func (s *LabelSet) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode LabelSet to nil")
	}
	m := s.init()
	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		var elem string
		if err := func() error {
			v, err := d.Str()
			elem = string(v)
			if err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return errors.Wrapf(err, "decode field %q", k)
		}
		m[string(k)] = elem
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode LabelSet")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s LabelSet) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *LabelSet) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Labels) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Labels) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("data")
		e.ArrStart()
		for _, elem := range s.Data {
			e.Str(elem)
		}
		e.ArrEnd()
	}
	{
		e.FieldStart("status")
		e.Str(s.Status)
	}
}

var jsonFieldsNameOfLabels = [2]string{
	0: "data",
	1: "status",
}

// Decode decodes Labels from json.
func (s *Labels) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Labels to nil")
	}
	var requiredBitSet [1]uint8
	s.setDefaults()

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "data":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				s.Data = make([]string, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem string
					v, err := d.Str()
					elem = string(v)
					if err != nil {
						return err
					}
					s.Data = append(s.Data, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"data\"")
			}
		case "status":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Status = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"status\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Labels")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfLabels) {
					name = jsonFieldsNameOfLabels[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Labels) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Labels) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Maps) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Maps) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("data")
		e.ArrStart()
		for _, elem := range s.Data {
			elem.Encode(e)
		}
		e.ArrEnd()
	}
	{
		e.FieldStart("status")
		e.Str(s.Status)
	}
}

var jsonFieldsNameOfMaps = [2]string{
	0: "data",
	1: "status",
}

// Decode decodes Maps from json.
func (s *Maps) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Maps to nil")
	}
	var requiredBitSet [1]uint8
	s.setDefaults()

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "data":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				s.Data = make([]MapsDataItem, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem MapsDataItem
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Data = append(s.Data, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"data\"")
			}
		case "status":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Status = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"status\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Maps")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfMaps) {
					name = jsonFieldsNameOfMaps[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Maps) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Maps) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s MapsDataItem) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields implements json.Marshaler.
func (s MapsDataItem) encodeFields(e *jx.Encoder) {
	for k, elem := range s {
		e.FieldStart(k)

		e.Str(elem)
	}
}

// Decode decodes MapsDataItem from json.
func (s *MapsDataItem) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode MapsDataItem to nil")
	}
	m := s.init()
	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		var elem string
		if err := func() error {
			v, err := d.Str()
			elem = string(v)
			if err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return errors.Wrapf(err, "decode field %q", k)
		}
		m[string(k)] = elem
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode MapsDataItem")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s MapsDataItem) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *MapsDataItem) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes LabelSet as json.
func (o OptLabelSet) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes LabelSet from json.
func (o *OptLabelSet) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptLabelSet to nil")
	}
	o.Set = true
	o.Value = make(LabelSet)
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptLabelSet) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptLabelSet) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Push) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Push) encodeFields(e *jx.Encoder) {
	{
		if s.Streams != nil {
			e.FieldStart("streams")
			e.ArrStart()
			for _, elem := range s.Streams {
				elem.Encode(e)
			}
			e.ArrEnd()
		}
	}
}

var jsonFieldsNameOfPush = [1]string{
	0: "streams",
}

// Decode decodes Push from json.
func (s *Push) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Push to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "streams":
			if err := func() error {
				s.Streams = make([]Stream, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem Stream
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Streams = append(s.Streams, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"streams\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Push")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Push) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Push) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *QueryResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *QueryResponse) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("status")
		e.Str(s.Status)
	}
	{
		e.FieldStart("data")
		s.Data.Encode(e)
	}
}

var jsonFieldsNameOfQueryResponse = [2]string{
	0: "status",
	1: "data",
}

// Decode decodes QueryResponse from json.
func (s *QueryResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode QueryResponse to nil")
	}
	var requiredBitSet [1]uint8
	s.setDefaults()

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "status":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.Status = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"status\"")
			}
		case "data":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				if err := s.Data.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"data\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode QueryResponse")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfQueryResponse) {
					name = jsonFieldsNameOfQueryResponse[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *QueryResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *QueryResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *QueryResponseData) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *QueryResponseData) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("resultType")
		s.ResultType.Encode(e)
	}
	{
		e.FieldStart("result")
		s.Result.Encode(e)
	}
	{
		e.FieldStart("stats")
		s.Stats.Encode(e)
	}
}

var jsonFieldsNameOfQueryResponseData = [3]string{
	0: "resultType",
	1: "result",
	2: "stats",
}

// Decode decodes QueryResponseData from json.
func (s *QueryResponseData) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode QueryResponseData to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "resultType":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				if err := s.ResultType.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"resultType\"")
			}
		case "result":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				if err := s.Result.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"result\"")
			}
		case "stats":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				if err := s.Stats.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"stats\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode QueryResponseData")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfQueryResponseData) {
					name = jsonFieldsNameOfQueryResponseData[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *QueryResponseData) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *QueryResponseData) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes QueryResponseDataResultType as json.
func (s QueryResponseDataResultType) Encode(e *jx.Encoder) {
	e.Str(string(s))
}

// Decode decodes QueryResponseDataResultType from json.
func (s *QueryResponseDataResultType) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode QueryResponseDataResultType to nil")
	}
	v, err := d.StrBytes()
	if err != nil {
		return err
	}
	// Try to use constant string.
	switch QueryResponseDataResultType(v) {
	case QueryResponseDataResultTypeStreams:
		*s = QueryResponseDataResultTypeStreams
	case QueryResponseDataResultTypeMatrix:
		*s = QueryResponseDataResultTypeMatrix
	default:
		*s = QueryResponseDataResultType(v)
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s QueryResponseDataResultType) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *QueryResponseDataResultType) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Stats) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Stats) encodeFields(e *jx.Encoder) {
}

var jsonFieldsNameOfStats = [0]string{}

// Decode decodes Stats from json.
func (s *Stats) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Stats to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Stats")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Stats) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Stats) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Stream) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Stream) encodeFields(e *jx.Encoder) {
	{
		if s.Stream.Set {
			e.FieldStart("stream")
			s.Stream.Encode(e)
		}
	}
	{
		if s.Metric.Set {
			e.FieldStart("metric")
			s.Metric.Encode(e)
		}
	}
	{
		e.FieldStart("values")
		e.ArrStart()
		for _, elem := range s.Values {
			elem.Encode(e)
		}
		e.ArrEnd()
	}
}

var jsonFieldsNameOfStream = [3]string{
	0: "stream",
	1: "metric",
	2: "values",
}

// Decode decodes Stream from json.
func (s *Stream) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Stream to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "stream":
			if err := func() error {
				s.Stream.Reset()
				if err := s.Stream.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"stream\"")
			}
		case "metric":
			if err := func() error {
				s.Metric.Reset()
				if err := s.Metric.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"metric\"")
			}
		case "values":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				s.Values = make([]Entry, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem Entry
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Values = append(s.Values, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"values\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Stream")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000100,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfStream) {
					name = jsonFieldsNameOfStream[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Stream) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Stream) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes Streams as json.
func (s Streams) Encode(e *jx.Encoder) {
	unwrapped := []Stream(s)

	e.ArrStart()
	for _, elem := range unwrapped {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes Streams from json.
func (s *Streams) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Streams to nil")
	}
	var unwrapped []Stream
	if err := func() error {
		unwrapped = make([]Stream, 0)
		if err := d.Arr(func(d *jx.Decoder) error {
			var elem Stream
			if err := elem.Decode(d); err != nil {
				return err
			}
			unwrapped = append(unwrapped, elem)
			return nil
		}); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = Streams(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s Streams) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Streams) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes Value as json.
func (s Value) Encode(e *jx.Encoder) {
	switch s.Type {
	case StringValue:
		e.Str(s.String)
	case Float64Value:
		e.Float64(s.Float64)
	}
}

// Decode decodes Value from json.
func (s *Value) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Value to nil")
	}
	// Sum type type_discriminator.
	switch t := d.Next(); t {
	case jx.Number:
		v, err := d.Float64()
		s.Float64 = float64(v)
		if err != nil {
			return err
		}
		s.Type = Float64Value
	case jx.String:
		v, err := d.Str()
		s.String = string(v)
		if err != nil {
			return err
		}
		s.Type = StringValue
	default:
		return errors.Errorf("unexpected json type %q", t)
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s Value) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Value) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Values) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Values) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("data")
		e.ArrStart()
		for _, elem := range s.Data {
			e.Str(elem)
		}
		e.ArrEnd()
	}
	{
		e.FieldStart("status")
		e.Str(s.Status)
	}
}

var jsonFieldsNameOfValues = [2]string{
	0: "data",
	1: "status",
}

// Decode decodes Values from json.
func (s *Values) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Values to nil")
	}
	var requiredBitSet [1]uint8
	s.setDefaults()

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "data":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				s.Data = make([]string, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem string
					v, err := d.Str()
					elem = string(v)
					if err != nil {
						return err
					}
					s.Data = append(s.Data, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"data\"")
			}
		case "status":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Status = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"status\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Values")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfValues) {
					name = jsonFieldsNameOfValues[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Values) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Values) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}
