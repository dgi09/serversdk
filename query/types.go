package query

type Operator uint8
type Operand uint16

const (
	Eq Operator = iota
	NEq
	Lt
	Gt
	LEt
	GEt
	In
	Like
)

type OrderDir uint8

const (
	Asc OrderDir = iota
	Desc
)

type Clause struct {
	Operand  Operand
	Operator Operator
	Value    any
}

type Order struct {
	Operand   Operand
	Direction OrderDir
}

type Query struct {
	Clauses  []Clause
	Ordering []Order
	Skip     int
	Take     int
}
