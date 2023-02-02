package query

import (
	"strings"
	"time"

	"golang.org/x/exp/slices"
)

func IntClauseResolver(clause Clause, val int) bool {
	switch clause.Operator {
	case Eq:
		return val == clause.Value.(int)
	case NEq:
		return val != clause.Value.(int)
	case Lt:
		return val < clause.Value.(int)
	case Gt:
		return val > clause.Value.(int)
	case LEt:
		return val <= clause.Value.(int)
	case GEt:
		return val >= clause.Value.(int)
	case In:
		if clause.Value == nil {
			return false
		}

		return slices.Contains(clause.Value.([]int), val)
	default:
		return false
	}
}

func Uint64ClauseResolver(clause Clause, val uint64) bool {
	switch clause.Operator {
	case Eq:
		return val == clause.Value.(uint64)
	case NEq:
		return val != clause.Value.(uint64)
	case Lt:
		return val < clause.Value.(uint64)
	case Gt:
		return val > clause.Value.(uint64)
	case LEt:
		return val <= clause.Value.(uint64)
	case GEt:
		return val >= clause.Value.(uint64)
	case In:
		if clause.Value == nil {
			return false
		}

		return slices.Contains(clause.Value.([]uint64), val)
	default:
		return false
	}
}

func Uint8ClauseResolver(clause Clause, val uint8) bool {
	switch clause.Operator {
	case Eq:
		return val == clause.Value.(uint8)
	case NEq:
		return val != clause.Value.(uint8)
	case Lt:
		return val < clause.Value.(uint8)
	case Gt:
		return val > clause.Value.(uint8)
	case LEt:
		return val <= clause.Value.(uint8)
	case GEt:
		return val >= clause.Value.(uint8)
	case In:
		if clause.Value == nil {
			return false
		}

		return slices.Contains(clause.Value.([]uint8), val)
	default:
		return false
	}
}

func UintClauseResolver(clause Clause, val uint) bool {
	switch clause.Operator {
	case Eq:
		return val == clause.Value.(uint)
	case NEq:
		return val != clause.Value.(uint)
	case Lt:
		return val < clause.Value.(uint)
	case Gt:
		return val > clause.Value.(uint)
	case LEt:
		return val <= clause.Value.(uint)
	case GEt:
		return val >= clause.Value.(uint)
	case In:
		if clause.Value == nil {
			return false
		}

		return slices.Contains(clause.Value.([]uint), val)
	default:
		return false
	}
}

func FloatClauseResolver(clause Clause, val float32) bool {
	switch clause.Operator {
	case Eq:
		return val == clause.Value.(float32)
	case NEq:
		return val != clause.Value.(float32)
	case Lt:
		return val < clause.Value.(float32)
	case Gt:
		return val > clause.Value.(float32)
	case LEt:
		return val <= clause.Value.(float32)
	case GEt:
		return val >= clause.Value.(float32)
	case In:
		if clause.Value == nil {
			return false
		}

		return slices.Contains(clause.Value.([]float32), val)
	default:
		return false
	}
}

func TimeClauseResolver(clause Clause, val time.Time) bool {
	switch clause.Operator {
	case Eq:
		return val == clause.Value.(time.Time)
	case NEq:
		return val != clause.Value.(time.Time)
	case Lt:
		return val.Before(clause.Value.(time.Time))
	case Gt:
		return val.After(clause.Value.(time.Time))
	case LEt:
		cmp := clause.Value.(time.Time)
		return val == cmp || val.Before(cmp)
	case GEt:
		cmp := clause.Value.(time.Time)
		return val == cmp || val.After(cmp)
	default:
		return false
	}
}

func BoolClauseResolver(clause Clause, val bool) bool {
	switch clause.Operator {
	case Eq:
		return val == clause.Value.(bool)
	case NEq:
		return val != clause.Value.(bool)

	case In:
		if clause.Value == nil {
			return false
		}

		return slices.Contains(clause.Value.([]bool), val)
	default:
		return false
	}
}

func StringClauseResolver(clause Clause, val string) bool {
	switch clause.Operator {
	case Eq:
		return val == clause.Value.(string)
	case NEq:
		return val != clause.Value.(string)
	case In:
		if clause.Value == nil {
			return false
		}

		return slices.Contains(clause.Value.([]string), val)
	case Like:
		if clause.Value == nil {
			return false
		}

		return strings.Contains(strings.ToLower(val), strings.ToLower(clause.Value.(string)))

	default:
		return false
	}
}
