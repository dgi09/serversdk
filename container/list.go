package container

import (
	"sort"

	"github.com/dgi09/serversdk/compare"
)

type List[T any] []T

func NewList[T any]() List[T] {
	return List[T]{}
}

func (l *List[T]) Add(val T) {
	*l = append(*l, val)
}

func (l *List[T]) AddRange(values []T) {
	*l = append(*l, values...)
}

func (l *List[T]) AddRangeRolled(values ...T) {
	*l = append(*l, values...)
}

func (list *List[T]) RemoveAt(index int) {
	l := len(*list)

	lVal := *list
	if index < 0 || index >= l {
		panic("Index out of bounds")
	}

	*list = append(lVal[:index], lVal[index+1:]...)
}

func (list List[T]) IndexOf(val T, comp compare.Comparer) int {
	result := -1

	for i, v := range list {
		if comp(v, val) == 0 {
			return i
		}
	}

	return result
}

func (list List[T]) FindIndex(expr func(T) bool) int {
	for i, v := range list {
		if expr(v) {
			return i
		}
	}

	return -1
}

func (list *List[T]) Set(i int, v T) {
	if i < 0 || i >= list.Count() {
		panic("Index out of bounds")
	}

	lVal := *list
	lVal[i] = v
}

func (list *List[T]) Remove(val T, comp compare.Comparer) {
	index := list.IndexOf(val, comp)

	if index != -1 {
		list.RemoveAt(index)
	}
}

func (list List[T]) Count() int {
	return len(list)
}

func (list List[T]) Where(expr func(T) bool) List[T] {
	res := make([]T, 0, list.Count())

	for _, v := range list {
		if expr(v) {
			res = append(res, v)
		}
	}

	return res
}

func (list List[T]) Any(expr func(T) bool) bool {
	for _, v := range list {
		if expr(v) {
			return true
		}
	}

	return false
}

func (list List[T]) FirstOrDefault(out *T, expr func(v T) bool) bool {
	for i, v := range list {
		if expr(v) {
			*out = list[i]
			return true
		}
	}

	return false
}

func (list *List[T]) SortAsc(comp func(l T, r T) int8) {
	lVal := *list
	sort.Slice(lVal, func(i, j int) bool {
		compRes := comp(lVal[i], lVal[j])

		return compRes == -1
	})
}

func (list *List[T]) SortDesc(comp func(l T, r T) int8) {
	lVal := *list
	sort.Slice(lVal, func(i, j int) bool {
		compRes := comp(lVal[i], lVal[j])

		return compRes == 1
	})
}
