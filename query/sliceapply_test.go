package query

import (
	"reflect"
	"testing"

	"github.com/dgi09/serversdk/compare"
)

type testStruct struct {
	A int
	B int
}

func TestCalculateSliceBounds(t *testing.T) {
	type args struct {
		skip int
		take int
		len  int
	}
	tests := []struct {
		args  args
		want  int
		want1 int
	}{
		{args{-1, -1, 10}, 0, 10},
		{args{0, -1, 10}, 0, 10},
		{args{3, -1, 10}, 3, 10},
		{args{10, -1, 10}, 10, 10},
		{args{12, -1, 10}, 10, 10},
		{args{-1, 4, 10}, 0, 4},
		{args{1, 4, 10}, 1, 5},
		{args{0, 4, 10}, 0, 4},
		{args{0, 10, 10}, 0, 10},
		{args{0, 12, 10}, 0, 10},
		{args{5, 14, 10}, 5, 10},
	}
	for _, tt := range tests {
		got, got1 := calculateSliceBounds(tt.args.skip, tt.args.take, tt.args.len)

		if got != tt.want {
			t.Errorf("calculateSliceBounds() got = %v, want %v", got, tt.want)
		}
		if got1 != tt.want1 {
			t.Errorf("calculateSliceBounds() got1 = %v, want %v", got1, tt.want1)
		}
	}
}

func TestShiftRight(t *testing.T) {
	tests := []struct {
		arr     []int
		index   int
		realLen int
		exp     []int
	}{
		{[]int{}, 0, 0, []int{}},
		{[]int{1}, 0, 0, []int{1}},
		{[]int{1, 0}, 0, 1, []int{1, 1}},
		{[]int{1, 2, 0}, 0, 2, []int{1, 1, 2}},
		{[]int{1, 2, 0}, 1, 2, []int{1, 2, 2}},
	}
	for _, tt := range tests {
		shiftRight(tt.arr, tt.index, tt.realLen)

		if !reflect.DeepEqual(tt.arr, tt.exp) {
			t.FailNow()
		}
	}
}

func clauseResolver(clause Clause, val int) bool {

	cmp := clause.Value.(int)

	switch clause.Operator {
	case Eq:
		return val == cmp
	case NEq:
		return val != cmp
	case Gt:
		return val > cmp
	case Lt:
		return val < cmp
	case LEt:
		return val <= cmp
	case GEt:
		return val >= cmp
	default:
		return false
	}
}

func orderResolver(operand Operand, left int, right int) int8 {
	return compare.IntComparer(left, right)
}

func TestSliceApplyPrimitive(t *testing.T) {
	emptyQ := NewQueryBuilder().Build()
	sortAscQ := NewQueryBuilder().Order(0, Asc).Build()
	sortDescQ := NewQueryBuilder().Order(0, Desc).Build()
	eq10Q := NewQueryBuilder().Equal(0, 10).Build()
	gt10Q := NewQueryBuilder().GreaterThan(0, 10).Build()
	gt10AscQ := NewQueryBuilder().GreaterThan(0, 10).Order(0, Asc).Build()
	bet10and20AscQ := NewQueryBuilder().GreaterOrEqualThan(0, 10).LessOrEqualhan(0, 20).Order(0, Asc).Build()

	tests := []struct {
		arr   []int
		query Query
		exp   []int
	}{
		{[]int{}, emptyQ, nil},
		{[]int{1, 2, 3}, emptyQ, []int{1, 2, 3}},
		{[]int{1, 2, 3}, sortAscQ, []int{1, 2, 3}},
		{[]int{3, 1, 2}, sortAscQ, []int{1, 2, 3}},
		{[]int{3, 2, 1}, sortDescQ, []int{3, 2, 1}},
		{[]int{3, 1, 2}, sortDescQ, []int{3, 2, 1}},
		{[]int{1, 2, 3}, eq10Q, []int{}},
		{[]int{3, 10, 2}, eq10Q, []int{10}},
		{[]int{3, 10, 2}, gt10Q, []int{}},
		{[]int{3, 11, 2}, gt10Q, []int{11}},
		{[]int{13, 3, 11, 2}, gt10AscQ, []int{11, 13}},
		{[]int{1, 2, 3, 10, 16, 15, 0, 11, 13, 12}, bet10and20AscQ, []int{10, 11, 12, 13, 15, 16}},
	}

	for _, tt := range tests {
		res := SliceApply(tt.arr, tt.query, clauseResolver, orderResolver)

		if !reflect.DeepEqual(res, tt.exp) {
			t.FailNow()
		}
	}
}

func testStructOrderResolver(operand Operand, left testStruct, right testStruct) int8 {
	switch operand {
	case 0:
		return compare.IntComparer(left.A, right.A)
	case 1:
		return compare.IntComparer(left.B, right.B)
	default:
		return 0
	}
}

func testStructClauseResolver(clause Clause, val testStruct) bool {
	return true
}

func TestSliceApplyComplex(t *testing.T) {
	emptyQ := NewQueryBuilder().Build()
	sortAAscQ := NewQueryBuilder().Order(0, Asc).Build()
	sortADescQ := NewQueryBuilder().Order(0, Desc).Build()
	sortAAscBAscQ := NewQueryBuilder().Order(0, Asc).Order(1, Asc).Build()
	sortAAscBDescQ := NewQueryBuilder().Order(0, Asc).Order(1, Desc).Build()

	tests := []struct {
		arr   []testStruct
		query Query
		exp   []testStruct
	}{
		{[]testStruct{}, emptyQ, nil},
		{[]testStruct{{1, 3}, {2, 3}}, sortAAscQ, []testStruct{{1, 3}, {2, 3}}},
		{[]testStruct{{2, 3}, {1, 3}}, sortAAscQ, []testStruct{{1, 3}, {2, 3}}},
		{[]testStruct{{2, 3}, {1, 3}}, sortADescQ, []testStruct{{2, 3}, {1, 3}}},
		{[]testStruct{{1, 3}, {2, 3}}, sortADescQ, []testStruct{{2, 3}, {1, 3}}},
		{[]testStruct{{1, 1}, {1, 2}, {2, 1}, {2, 2}}, sortAAscBAscQ, []testStruct{{1, 1}, {1, 2}, {2, 1}, {2, 2}}},
		{[]testStruct{{2, 2}, {1, 2}, {1, 1}, {2, 1}}, sortAAscBAscQ, []testStruct{{1, 1}, {1, 2}, {2, 1}, {2, 2}}},
		{[]testStruct{{1, 1}, {1, 2}, {2, 1}, {2, 2}}, sortAAscBDescQ, []testStruct{{1, 2}, {1, 1}, {2, 2}, {2, 1}}},
		{[]testStruct{{2, 2}, {1, 2}, {1, 1}, {2, 1}}, sortAAscBDescQ, []testStruct{{1, 2}, {1, 1}, {2, 2}, {2, 1}}},
	}

	for _, tt := range tests {
		res := SliceApply(tt.arr, tt.query, testStructClauseResolver, testStructOrderResolver)

		if !reflect.DeepEqual(res, tt.exp) {
			t.FailNow()
		}
	}
}
