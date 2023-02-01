package main

import (
	"fmt"

	"github.com/dgi09/serversdk/compare"
	"github.com/dgi09/serversdk/query"
)

type Person struct {
	Id  int
	Age int
}

func PersonOrderResolver(operand query.Operand, left Person, right Person) int8 {
	switch operand {
	case 0:
		return compare.IntComparer(left.Id, right.Id)
	case 1:
		return compare.IntComparer(left.Age, right.Age)
	default:
		return 0
	}
}

func IntOrderResolver(left int, right int) int8 {
	if right < left {
		return -1
	}

	if right > left {
		return 1
	}

	return 0
}

func PersonClauseResolve(clause query.Clause, person Person) bool {
	return true
}

func main() {
	q := query.NewQueryBuilder().Order(0, query.Asc).Order(1, query.Desc).Build()

	set := []Person{
		{1, 50},
		{2, 30},
		{1, 20},
		{3, 40},
		{3, 60},
		{1, 30},
	}

	set = query.SliceApply(set, q, PersonClauseResolve, PersonOrderResolver)

	fmt.Print(set)
}
