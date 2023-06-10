// Code generated by ogen, DO NOT EDIT.

package promapi

import (
	"math/bits"
	"strconv"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/validate"
)

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

// Encode encodes InstantQueryData as json.
func (s InstantQueryData) Encode(e *jx.Encoder) {
	switch s.Type {
	case MatrixInstantQueryData:
		e.ObjStart()
		e.FieldStart("resultType")
		e.Str("matrix")
		s.Matrix.encodeFields(e)
		e.ObjEnd()
	}
}

// Decode decodes InstantQueryData from json.
func (s *InstantQueryData) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode InstantQueryData to nil")
	}
	// Sum type discriminator.
	if typ := d.Next(); typ != jx.Object {
		return errors.Errorf("unexpected json type %q", typ)
	}

	var found bool
	if err := d.Capture(func(d *jx.Decoder) error {
		return d.ObjBytes(func(d *jx.Decoder, key []byte) error {
			if found {
				return d.Skip()
			}
			switch string(key) {
			case "resultType":
				typ, err := d.Str()
				if err != nil {
					return err
				}
				switch typ {
				case "matrix":
					s.Type = MatrixInstantQueryData
					found = true
				default:
					return errors.Errorf("unknown type %s", typ)
				}
				return nil
			}
			return d.Skip()
		})
	}); err != nil {
		return errors.Wrap(err, "capture")
	}
	if !found {
		return errors.New("unable to detect sum type variant")
	}
	switch s.Type {
	case MatrixInstantQueryData:
		if err := s.Matrix.Decode(d); err != nil {
			return err
		}
	default:
		return errors.Errorf("inferred invalid type: %s", s.Type)
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s InstantQueryData) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *InstantQueryData) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *InstantQueryResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *InstantQueryResponse) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("status")
		e.Str(s.Status)
	}
	{
		e.FieldStart("data")
		s.Data.Encode(e)
	}
}

var jsonFieldsNameOfInstantQueryResponse = [2]string{
	0: "status",
	1: "data",
}

// Decode decodes InstantQueryResponse from json.
func (s *InstantQueryResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode InstantQueryResponse to nil")
	}
	var requiredBitSet [1]uint8

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
		return errors.Wrap(err, "decode InstantQueryResponse")
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
				if fieldIdx < len(jsonFieldsNameOfInstantQueryResponse) {
					name = jsonFieldsNameOfInstantQueryResponse[fieldIdx]
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
func (s *InstantQueryResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *InstantQueryResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Matrix) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Matrix) encodeFields(e *jx.Encoder) {
}

var jsonFieldsNameOfMatrix = [0]string{}

// Decode decodes Matrix from json.
func (s *Matrix) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Matrix to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Matrix")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Matrix) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Matrix) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}
