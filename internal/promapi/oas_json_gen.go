// Code generated by ogen, DO NOT EDIT.

package promapi

import (
	"math/bits"
	"strconv"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/validate"
)

// Encode encodes Data as json.
func (s Data) Encode(e *jx.Encoder) {
	switch s.Type {
	case MatrixData:
		e.ObjStart()
		e.FieldStart("resultType")
		e.Str("matrix")
		s.Matrix.encodeFields(e)
		e.ObjEnd()
	case ScalarData:
		e.ObjStart()
		e.FieldStart("resultType")
		e.Str("scalar")
		s.Scalar.encodeFields(e)
		e.ObjEnd()
	case StringData:
		e.ObjStart()
		e.FieldStart("resultType")
		e.Str("string")
		s.String.encodeFields(e)
		e.ObjEnd()
	case VectorData:
		e.ObjStart()
		e.FieldStart("resultType")
		e.Str("vector")
		s.Vector.encodeFields(e)
		e.ObjEnd()
	}
}

// Decode decodes Data from json.
func (s *Data) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Data to nil")
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
					s.Type = MatrixData
					found = true
				case "scalar":
					s.Type = ScalarData
					found = true
				case "string":
					s.Type = StringData
					found = true
				case "vector":
					s.Type = VectorData
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
	case MatrixData:
		if err := s.Matrix.Decode(d); err != nil {
			return err
		}
	case VectorData:
		if err := s.Vector.Decode(d); err != nil {
			return err
		}
	case ScalarData:
		if err := s.Scalar.Decode(d); err != nil {
			return err
		}
	case StringData:
		if err := s.String.Decode(d); err != nil {
			return err
		}
	default:
		return errors.Errorf("inferred invalid type: %s", s.Type)
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s Data) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Data) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Fail) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Fail) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("status")
		e.Str(s.Status)
	}
	{
		e.FieldStart("error")
		e.Str(s.Error)
	}
	{
		e.FieldStart("errorType")
		s.ErrorType.Encode(e)
	}
	{
		if s.Data.Set {
			e.FieldStart("data")
			s.Data.Encode(e)
		}
	}
}

var jsonFieldsNameOfFail = [4]string{
	0: "status",
	1: "error",
	2: "errorType",
	3: "data",
}

// Decode decodes Fail from json.
func (s *Fail) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Fail to nil")
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
		case "error":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Error = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"error\"")
			}
		case "errorType":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				if err := s.ErrorType.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"errorType\"")
			}
		case "data":
			if err := func() error {
				s.Data.Reset()
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
		return errors.Wrap(err, "decode Fail")
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
				if fieldIdx < len(jsonFieldsNameOfFail) {
					name = jsonFieldsNameOfFail[fieldIdx]
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
func (s *Fail) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Fail) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes FailErrorType as json.
func (s FailErrorType) Encode(e *jx.Encoder) {
	e.Str(string(s))
}

// Decode decodes FailErrorType from json.
func (s *FailErrorType) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode FailErrorType to nil")
	}
	v, err := d.StrBytes()
	if err != nil {
		return err
	}
	// Try to use constant string.
	switch FailErrorType(v) {
	case FailErrorTypeTimeout:
		*s = FailErrorTypeTimeout
	case FailErrorTypeCanceled:
		*s = FailErrorTypeCanceled
	case FailErrorTypeExecution:
		*s = FailErrorTypeExecution
	case FailErrorTypeBadData:
		*s = FailErrorTypeBadData
	case FailErrorTypeInternal:
		*s = FailErrorTypeInternal
	case FailErrorTypeUnavailable:
		*s = FailErrorTypeUnavailable
	case FailErrorTypeNotFound:
		*s = FailErrorTypeNotFound
	default:
		*s = FailErrorType(v)
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s FailErrorType) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *FailErrorType) UnmarshalJSON(data []byte) error {
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
	{
		e.FieldStart("result")
		e.ArrStart()
		for _, elem := range s.Result {
			elem.Encode(e)
		}
		e.ArrEnd()
	}
}

var jsonFieldsNameOfMatrix = [1]string{
	0: "result",
}

// Decode decodes Matrix from json.
func (s *Matrix) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Matrix to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "result":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				s.Result = make([]MatrixResultItem, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem MatrixResultItem
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Result = append(s.Result, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"result\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Matrix")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
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
				if fieldIdx < len(jsonFieldsNameOfMatrix) {
					name = jsonFieldsNameOfMatrix[fieldIdx]
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

// Encode implements json.Marshaler.
func (s *MatrixResultItem) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *MatrixResultItem) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("metric")
		s.Metric.Encode(e)
	}
	{
		if s.Values != nil {
			e.FieldStart("values")
			e.ArrStart()
			for _, elem := range s.Values {
				elem.Encode(e)
			}
			e.ArrEnd()
		}
	}
}

var jsonFieldsNameOfMatrixResultItem = [2]string{
	0: "metric",
	1: "values",
}

// Decode decodes MatrixResultItem from json.
func (s *MatrixResultItem) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode MatrixResultItem to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "metric":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				if err := s.Metric.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"metric\"")
			}
		case "values":
			if err := func() error {
				s.Values = make([]Value, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem Value
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
		return errors.Wrap(err, "decode MatrixResultItem")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
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
				if fieldIdx < len(jsonFieldsNameOfMatrixResultItem) {
					name = jsonFieldsNameOfMatrixResultItem[fieldIdx]
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
func (s *MatrixResultItem) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *MatrixResultItem) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s MatrixResultItemMetric) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields implements json.Marshaler.
func (s MatrixResultItemMetric) encodeFields(e *jx.Encoder) {
	for k, elem := range s {
		e.FieldStart(k)

		e.Str(elem)
	}
}

// Decode decodes MatrixResultItemMetric from json.
func (s *MatrixResultItemMetric) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode MatrixResultItemMetric to nil")
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
		return errors.Wrap(err, "decode MatrixResultItemMetric")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s MatrixResultItemMetric) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *MatrixResultItemMetric) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes Data as json.
func (o OptData) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes Data from json.
func (o *OptData) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptData to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptData) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptData) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Scalar) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Scalar) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("result")
		s.Result.Encode(e)
	}
}

var jsonFieldsNameOfScalar = [1]string{
	0: "result",
}

// Decode decodes Scalar from json.
func (s *Scalar) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Scalar to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "result":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				if err := s.Result.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"result\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Scalar")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
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
				if fieldIdx < len(jsonFieldsNameOfScalar) {
					name = jsonFieldsNameOfScalar[fieldIdx]
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
func (s *Scalar) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Scalar) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *String) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *String) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("result")
		e.Str(s.Result)
	}
}

var jsonFieldsNameOfString = [1]string{
	0: "result",
}

// Decode decodes String from json.
func (s *String) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode String to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "result":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.Result = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"result\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode String")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
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
				if fieldIdx < len(jsonFieldsNameOfString) {
					name = jsonFieldsNameOfString[fieldIdx]
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
func (s *String) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *String) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Success) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Success) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("status")
		e.Str(s.Status)
	}
	{
		if s.Warnings != nil {
			e.FieldStart("warnings")
			e.ArrStart()
			for _, elem := range s.Warnings {
				e.Str(elem)
			}
			e.ArrEnd()
		}
	}
	{
		e.FieldStart("data")
		s.Data.Encode(e)
	}
}

var jsonFieldsNameOfSuccess = [3]string{
	0: "status",
	1: "warnings",
	2: "data",
}

// Decode decodes Success from json.
func (s *Success) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Success to nil")
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
		case "warnings":
			if err := func() error {
				s.Warnings = make([]string, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem string
					v, err := d.Str()
					elem = string(v)
					if err != nil {
						return err
					}
					s.Warnings = append(s.Warnings, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"warnings\"")
			}
		case "data":
			requiredBitSet[0] |= 1 << 2
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
		return errors.Wrap(err, "decode Success")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000101,
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
				if fieldIdx < len(jsonFieldsNameOfSuccess) {
					name = jsonFieldsNameOfSuccess[fieldIdx]
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
func (s *Success) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Success) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes Value as json.
func (s Value) Encode(e *jx.Encoder) {
	unwrapped := []ValueItem(s)

	e.ArrStart()
	for _, elem := range unwrapped {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes Value from json.
func (s *Value) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Value to nil")
	}
	var unwrapped []ValueItem
	if err := func() error {
		unwrapped = make([]ValueItem, 0)
		if err := d.Arr(func(d *jx.Decoder) error {
			var elem ValueItem
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
	*s = Value(unwrapped)
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

// Encode encodes ValueItem as json.
func (s ValueItem) Encode(e *jx.Encoder) {
	switch s.Type {
	case IntValueItem:
		e.Int(s.Int)
	case StringValueItem:
		e.Str(s.String)
	}
}

// Decode decodes ValueItem from json.
func (s *ValueItem) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ValueItem to nil")
	}
	// Sum type type_discriminator.
	switch t := d.Next(); t {
	case jx.Number:
		v, err := d.Int()
		s.Int = int(v)
		if err != nil {
			return err
		}
		s.Type = IntValueItem
	case jx.String:
		v, err := d.Str()
		s.String = string(v)
		if err != nil {
			return err
		}
		s.Type = StringValueItem
	default:
		return errors.Errorf("unexpected json type %q", t)
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s ValueItem) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ValueItem) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Vector) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Vector) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("result")
		e.ArrStart()
		for _, elem := range s.Result {
			elem.Encode(e)
		}
		e.ArrEnd()
	}
}

var jsonFieldsNameOfVector = [1]string{
	0: "result",
}

// Decode decodes Vector from json.
func (s *Vector) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Vector to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "result":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				s.Result = make([]VectorResultItem, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem VectorResultItem
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Result = append(s.Result, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"result\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Vector")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
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
				if fieldIdx < len(jsonFieldsNameOfVector) {
					name = jsonFieldsNameOfVector[fieldIdx]
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
func (s *Vector) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Vector) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *VectorResultItem) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *VectorResultItem) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("metric")
		s.Metric.Encode(e)
	}
	{
		e.FieldStart("value")
		s.Value.Encode(e)
	}
}

var jsonFieldsNameOfVectorResultItem = [2]string{
	0: "metric",
	1: "value",
}

// Decode decodes VectorResultItem from json.
func (s *VectorResultItem) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode VectorResultItem to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "metric":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				if err := s.Metric.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"metric\"")
			}
		case "value":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				if err := s.Value.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"value\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode VectorResultItem")
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
				if fieldIdx < len(jsonFieldsNameOfVectorResultItem) {
					name = jsonFieldsNameOfVectorResultItem[fieldIdx]
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
func (s *VectorResultItem) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *VectorResultItem) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s VectorResultItemMetric) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields implements json.Marshaler.
func (s VectorResultItemMetric) encodeFields(e *jx.Encoder) {
	for k, elem := range s {
		e.FieldStart(k)

		e.Str(elem)
	}
}

// Decode decodes VectorResultItemMetric from json.
func (s *VectorResultItemMetric) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode VectorResultItemMetric to nil")
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
		return errors.Wrap(err, "decode VectorResultItemMetric")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s VectorResultItemMetric) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *VectorResultItemMetric) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}
