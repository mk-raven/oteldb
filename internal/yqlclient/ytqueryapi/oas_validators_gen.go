// Code generated by ogen, DO NOT EDIT.

package ytqueryapi

import (
	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/validate"
)

func (s Engine) Validate() error {
	switch s {
	case "yql":
		return nil
	case "ql":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s OperationState) Validate() error {
	switch s {
	case "running":
		return nil
	case "pending":
		return nil
	case "completed":
		return nil
	case "failed":
		return nil
	case "aborted":
		return nil
	case "reviving":
		return nil
	case "initializing":
		return nil
	case "preparing":
		return nil
	case "materializing":
		return nil
	case "completing":
		return nil
	case "aborting":
		return nil
	case "failing":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s OutputFormat) Validate() error {
	switch s {
	case "json":
		return nil
	case "yson":
		return nil
	case "dsv":
		return nil
	case "schemaful_dsv":
		return nil
	case "protobuf":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s *QueryStatus) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.Engine.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "engine",
			Error: err,
		})
	}
	if err := func() error {
		if err := s.State.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "state",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}