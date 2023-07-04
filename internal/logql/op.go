package logql

import "fmt"

// BinOp defines binary operation.
type BinOp int

const (
	// Logical ops.
	OpAnd BinOp = iota + 1
	OpOr
	OpUnless
	// Math ops.
	OpAdd
	OpSub
	OpMul
	OpDiv
	OpMod
	OpPow
	// Comparison ops.
	OpEq
	OpNotEq
	OpRe
	OpNotRe
	OpGt
	OpGte
	OpLt
	OpLte
)

// String implements fmt.Stringer.
func (op BinOp) String() string {
	switch op {
	case OpAnd:
		return "and"
	case OpOr:
		return "or"
	case OpUnless:
		return "unless"
	case OpAdd:
		return "+"
	case OpSub:
		return "-"
	case OpMul:
		return "*"
	case OpDiv:
		return "/"
	case OpMod:
		return "%"
	case OpPow:
		return "^"
	case OpEq:
		return "="
	case OpNotEq:
		return "!="
	case OpRe:
		return "=~"
	case OpNotRe:
		return "!~"
	case OpGt:
		return "<"
	case OpGte:
		return "<="
	case OpLt:
		return ">"
	case OpLte:
		return ">="
	default:
		return fmt.Sprintf("<unknown op %d>", op)
	}
}

// IsLogic returns whether operation is logical.
func (op BinOp) IsLogic() bool {
	switch op {
	case OpAnd, OpOr, OpUnless:
		return true
	default:
		return false
	}
}

// RangeOp defines range aggregation operation.
type RangeOp int

const (
	RangeOpCount RangeOp = iota + 1
	RangeOpRate
	RangeOpRateCounter
	RangeOpBytes
	RangeOpBytesRate
	RangeOpAvg
	RangeOpSum
	RangeOpMin
	RangeOpMax
	RangeOpStdvar
	RangeOpStddev
	RangeOpQuantile
	RangeOpFirst
	RangeOpLast
	RangeOpAbsent
)

// String implements fmt.Stringer.
func (op RangeOp) String() string {
	switch op {
	case RangeOpCount:
		return "count_over_time"
	case RangeOpRate:
		return "rate"
	case RangeOpRateCounter:
		return "rate_counter"
	case RangeOpBytes:
		return "bytes_over_time"
	case RangeOpBytesRate:
		return "bytes_rate"
	case RangeOpAvg:
		return "avg_over_time"
	case RangeOpSum:
		return "sum_over_time"
	case RangeOpMin:
		return "min_over_time"
	case RangeOpMax:
		return "max_over_time"
	case RangeOpStdvar:
		return "stdvar_over_time"
	case RangeOpStddev:
		return "stddev_over_time"
	case RangeOpQuantile:
		return "quantile_over_time"
	case RangeOpFirst:
		return "first_over_time"
	case RangeOpLast:
		return "last_over_time"
	case RangeOpAbsent:
		return "absent_over_time"
	default:
		return fmt.Sprintf("<unknown range op %d>", op)
	}
}

// VectorOp defines vector aggregation operation.
type VectorOp int

const (
	VectorOpSum VectorOp = iota + 1
	VectorOpAvg
	VectorOpCount
	VectorOpMax
	VectorOpMin
	VectorOpStddev
	VectorOpStdvar
	VectorOpBottomk
	VectorOpTopk
	VectorOpSort
	VectorOpSortDesc
)

// String implements fmt.Stringer.
func (op VectorOp) String() string {
	switch op {
	case VectorOpSum:
		return "sum"
	case VectorOpAvg:
		return "avg"
	case VectorOpCount:
		return "count"
	case VectorOpMax:
		return "max"
	case VectorOpMin:
		return "min"
	case VectorOpStddev:
		return "stddev"
	case VectorOpStdvar:
		return "stdvar"
	case VectorOpBottomk:
		return "bottomk"
	case VectorOpTopk:
		return "topk"
	case VectorOpSort:
		return "sort"
	case VectorOpSortDesc:
		return "sort_desc"
	default:
		return fmt.Sprintf("<unknown vector op %d>", op)
	}
}