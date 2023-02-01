package query

type QueryBuilder struct {
	query Query
}

func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{
		query: Query{
			Skip: -1,
			Take: -1,
		},
	}
}

func (b *QueryBuilder) addClause(operand Operand, operator Operator, value any) {
	clause := Clause{
		Operand:  operand,
		Operator: operator,
		Value:    value,
	}

	b.query.Clauses = append(b.query.Clauses, clause)
}

func (b *QueryBuilder) Equal(operand Operand, value any) *QueryBuilder {

	b.addClause(operand, Eq, value)

	return b
}

func (b *QueryBuilder) NotEqual(operand Operand, value any) *QueryBuilder {

	b.addClause(operand, NEq, value)

	return b
}

func (b *QueryBuilder) LessThan(operand Operand, value any) *QueryBuilder {

	b.addClause(operand, Lt, value)

	return b
}

func (b *QueryBuilder) LessOrEqualhan(operand Operand, value any) *QueryBuilder {

	b.addClause(operand, LEt, value)

	return b
}

func (b *QueryBuilder) GreaterThan(operand Operand, value any) *QueryBuilder {

	b.addClause(operand, Gt, value)

	return b
}

func (b *QueryBuilder) GreaterOrEqualThan(operand Operand, value any) *QueryBuilder {

	b.addClause(operand, GEt, value)

	return b
}

func (b *QueryBuilder) Like(operand Operand, value any) *QueryBuilder {

	b.addClause(operand, Like, value)

	return b
}

func (b *QueryBuilder) In(operand Operand, value any) *QueryBuilder {

	b.addClause(operand, In, value)

	return b
}

func (b *QueryBuilder) Paginate(skip int, take int) *QueryBuilder {
	b.query.Skip = skip
	b.query.Take = take

	return b
}

func (b *QueryBuilder) Order(operand Operand, dir OrderDir) *QueryBuilder {
	order := Order{
		Operand:   operand,
		Direction: dir,
	}

	b.query.Ordering = append(b.query.Ordering, order)

	return b
}

func (b *QueryBuilder) Build() Query {
	return b.query
}
