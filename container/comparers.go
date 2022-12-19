package container

import "time"

type Comparer = func(l interface{}, r interface{}) int8

func IntComparer(l interface{}, r interface{}) int8 {
	il := l.(int)
	ir := r.(int)

	if il < ir {
		return -1
	}

	if il > ir {
		return 1
	}

	return 0
}

func StringComparer(l interface{}, r interface{}) int8 {
	sl := l.(string)
	sr := r.(string)

	if sl < sr {
		return -1
	}

	if sl > sr {
		return 1
	}

	return 0

}

func TimeComparer(l interface{}, r interface{}) int8 {
	tl := l.(time.Time)
	tr := r.(time.Time)

	if tl.Before(tr) {
		return -1
	}

	if tr.Before(tl) {
		return 1
	}

	return 0
}
