package compare

import (
	"strings"
	"time"
)

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

func BoolComparer(l interface{}, r interface{}) int8 {
	bl := l.(bool)
	br := r.(bool)

	il := 0
	if bl {
		il = 1
	}

	ir := 0

	if br {
		ir = 1
	}

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

	return int8(strings.Compare(sl, sr))
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

func FloatComparer(l interface{}, r interface{}) int8 {
	fl := l.(float32)
	fr := r.(float32)

	if fl < fr {
		return -1
	}

	if fl > fr {
		return 1
	}

	return 0
}
