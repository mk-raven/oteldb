// Code generated by ogen, DO NOT EDIT.

package lokiapi

import (
	"fmt"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/validate"
)

func (s Direction) Validate() error {
	switch s {
	case "backward":
		return nil
	case "forward":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s *Labels) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.Data == nil {
			return errors.New("nil is invalid value")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "data",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *Maps) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.Data == nil {
			return errors.New("nil is invalid value")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "data",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s Matrix) Validate() error {
	alias := ([]Series)(s)
	if alias == nil {
		return errors.New("nil is invalid value")
	}
	var failures []validate.FieldError
	for i, elem := range alias {
		if err := func() error {
			if err := elem.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			failures = append(failures, validate.FieldError{
				Name:  fmt.Sprintf("[%d]", i),
				Error: err,
			})
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *MatrixResult) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.Result.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "result",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s PrometheusDuration) Validate() error {
	alias := (string)(s)
	if err := (validate.String{
		MinLength:    0,
		MinLengthSet: false,
		MaxLength:    0,
		MaxLengthSet: false,
		Email:        false,
		Hostname:     false,
		Regex:        regexMap["^[0-9smhdwy]+$"],
	}).Validate(string(alias)); err != nil {
		return errors.Wrap(err, "string")
	}
	return nil
}
func (s *PrometheusSamplePair) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.Float{}).Validate(float64(s.T)); err != nil {
			return errors.Wrap(err, "float")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "T",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *Push) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		var failures []validate.FieldError
		for i, elem := range s.Streams {
			if err := func() error {
				if err := elem.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
		}
		if len(failures) > 0 {
			return &validate.Error{Fields: failures}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "streams",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *QueryResponse) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.Data.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "data",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s QueryResponseData) Validate() error {
	switch s.Type {
	case StreamsResultQueryResponseData:
		if err := s.StreamsResult.Validate(); err != nil {
			return err
		}
		return nil
	case ScalarResultQueryResponseData:
		if err := s.ScalarResult.Validate(); err != nil {
			return err
		}
		return nil
	case VectorResultQueryResponseData:
		if err := s.VectorResult.Validate(); err != nil {
			return err
		}
		return nil
	case MatrixResultQueryResponseData:
		if err := s.MatrixResult.Validate(); err != nil {
			return err
		}
		return nil
	default:
		return errors.Errorf("invalid type %q", s.Type)
	}
}

func (s *ScalarResult) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.Result.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "result",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *Series) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		var failures []validate.FieldError
		for i, elem := range s.Values {
			if err := func() error {
				if err := elem.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
		}
		if len(failures) > 0 {
			return &validate.Error{Fields: failures}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "values",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *Stream) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.Values == nil {
			return errors.New("nil is invalid value")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "values",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s Streams) Validate() error {
	alias := ([]Stream)(s)
	if alias == nil {
		return errors.New("nil is invalid value")
	}
	var failures []validate.FieldError
	for i, elem := range alias {
		if err := func() error {
			if err := elem.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			failures = append(failures, validate.FieldError{
				Name:  fmt.Sprintf("[%d]", i),
				Error: err,
			})
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *StreamsResult) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.Result.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "result",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *Values) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.Data == nil {
			return errors.New("nil is invalid value")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "data",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *Vector) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.Value.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "value",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *VectorResult) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.Result == nil {
			return errors.New("nil is invalid value")
		}
		var failures []validate.FieldError
		for i, elem := range s.Result {
			if err := func() error {
				if err := elem.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
		}
		if len(failures) > 0 {
			return &validate.Error{Fields: failures}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "result",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
