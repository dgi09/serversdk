package query

func calculateSliceBounds(skip int, take int, len int) (int, int) {
	if skip == -1 {
		skip = 0
	}

	if skip > len {
		skip = len
	}

	if take == -1 {
		take = len - skip
	}

	if skip+take > len {
		take = len - skip
	}

	return skip, skip + take
}

func shiftRight[T any](arr []T, index int, realLen int) {
	if realLen > len(arr) {
		return
	}

	for i := realLen; i > index; i-- {
		arr[i] = arr[i-1]
	}
}

func SliceApply[T any](value []T, query Query, clauseResolver func(clause Clause, val T) bool, orderResolver func(operand Operand, left T, right T) int8) []T {
	if len(value) == 0 {
		return nil
	}

	resLen := 0
	res := make([]T, len(value))

	for vInd := range value {
		val := value[vInd]
		match := len(query.Clauses) == 0

		for clInd := range query.Clauses {
			clause := query.Clauses[clInd]
			match = clauseResolver(clause, val)

			if !match {
				break
			}
		}

		if match {
			insertIndex := 0

			if len(query.Ordering) == 0 {
				insertIndex = resLen
			}

			for resInd := 0; resInd < resLen; resInd++ {
				for orInd := range query.Ordering {
					order := query.Ordering[orInd]
					compare := orderResolver(order.Operand, val, res[resInd])

					if compare == 0 {
						continue
					}

					if order.Direction == Asc && compare == 1 {
						insertIndex++
					}

					if order.Direction == Desc && compare == -1 {
						insertIndex++
					}

					break
				}
			}

			shiftRight(res, insertIndex, resLen)
			res[insertIndex] = val
			resLen++
		}
	}

	bondsBegin, boundsEnd := calculateSliceBounds(query.Skip, query.Take, resLen)

	return res[bondsBegin:boundsEnd]
}
