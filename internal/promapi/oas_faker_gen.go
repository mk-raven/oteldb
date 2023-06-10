// Code generated by ogen, DO NOT EDIT.

package promapi

import (
	"fmt"
)

// SetFake set fake values.
func (s *Data) SetFake() {
	var variant Matrix

	{
		variant.SetFake()
	}
	s.SetMatrix(variant)
}

// SetFake set fake values.
func (s *Fail) SetFake() {
	{
		{
			s.Status = "string"
		}
	}
	{
		{
			s.Error = "string"
		}
	}
	{
		{
			s.ErrorType.SetFake()
		}
	}
	{
		{
			s.Data.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *FailErrorType) SetFake() {
	*s = FailErrorTypeTimeout
}

// SetFake set fake values.
func (s *Matrix) SetFake() {
	{
		{
			s.Result = nil
			for i := 0; i < 0; i++ {
				var elem MatrixResultItem
				{
					elem.SetFake()
				}
				s.Result = append(s.Result, elem)
			}
		}
	}
}

// SetFake set fake values.
func (s *MatrixResultItem) SetFake() {
	{
		{
			s.Metric.SetFake()
		}
	}
	{
		{
			s.Values = nil
			for i := 0; i < 0; i++ {
				var elem Value
				{
					elem.SetFake()
				}
				s.Values = append(s.Values, elem)
			}
		}
	}
}

// SetFake set fake values.
func (s *MatrixResultItemMetric) SetFake() {
	var (
		elem string
		m    map[string]string = s.init()
	)
	for i := 0; i < 0; i++ {
		m[fmt.Sprintf("fake%d", i)] = elem
	}
}

// SetFake set fake values.
func (s *OptData) SetFake() {
	var elem Data
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *Scalar) SetFake() {
	{
		{
			s.Result.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *String) SetFake() {
	{
		{
			s.Result = "string"
		}
	}
}

// SetFake set fake values.
func (s *Success) SetFake() {
	{
		{
			s.Status = "string"
		}
	}
	{
		{
			s.Warnings = nil
			for i := 0; i < 0; i++ {
				var elem string
				{
					elem = "string"
				}
				s.Warnings = append(s.Warnings, elem)
			}
		}
	}
	{
		{
			s.Data.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *Value) SetFake() {
	var unwrapped []ValueItem
	{
		unwrapped = nil
		for i := 0; i < 2; i++ {
			var elem ValueItem
			{
				elem.SetFake()
			}
			unwrapped = append(unwrapped, elem)
		}
	}
	*s = Value(unwrapped)
}

// SetFake set fake values.
func (s *ValueItem) SetFake() {
	var variant int

	{
		variant = int(0)
	}
	s.SetInt(variant)
}

// SetFake set fake values.
func (s *Vector) SetFake() {
	{
		{
			s.Result = nil
			for i := 0; i < 0; i++ {
				var elem VectorResultItem
				{
					elem.SetFake()
				}
				s.Result = append(s.Result, elem)
			}
		}
	}
}

// SetFake set fake values.
func (s *VectorResultItem) SetFake() {
	{
		{
			s.Metric.SetFake()
		}
	}
	{
		{
			s.Value.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *VectorResultItemMetric) SetFake() {
	var (
		elem string
		m    map[string]string = s.init()
	)
	for i := 0; i < 0; i++ {
		m[fmt.Sprintf("fake%d", i)] = elem
	}
}
