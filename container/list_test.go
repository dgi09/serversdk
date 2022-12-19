package container

import "testing"

func TestAdd(t *testing.T) {
	l := NewList[int]()

	l.Add(1)

	if len(l) != 1 || l[0] != 1 {
		t.FailNow()
	}

	l.Add(2)

	if len(l) != 2 || l[1] != 2 {
		t.FailNow()
	}
}

func TestAddRange(t *testing.T) {
	l := NewList[int]()

	l.AddRange([]int{1, 2})

	if len(l) != 2 || l[0] != 1 || l[1] != 2 {
		t.FailNow()
	}
}

func TestAddRangeRolled(t *testing.T) {
	l := NewList[int]()

	l.AddRangeRolled(1, 2)

	if len(l) != 2 || l[0] != 1 || l[1] != 2 {
		t.FailNow()
	}
}

func TestCount(t *testing.T) {
	l := NewList[int]()

	if l.Count() != len(l) {
		t.FailNow()
	}

	l.Add(1)

	if l.Count() != len(l) {
		t.FailNow()
	}
}

func TestRemoveAt(t *testing.T) {
	l := NewList[int]()

	l.AddRangeRolled(1, 2, 3)

	l.RemoveAt(1)

	if l.Count() != 2 || l[0] != 1 || l[1] != 3 {
		t.FailNow()
	}
}

func TestFindIndex(t *testing.T) {
	l := NewList[int]()

	l.AddRangeRolled(1, 2, 3, 4, 5, 6)

	tests := []struct {
		val, res int
	}{
		{1, 0},
		{2, 1},
		{3, 2},
		{10, -1},
		{6, 5},
		{0, -1},
	}

	for _, tt := range tests {
		i := l.FindIndex(func(v int) bool {
			return v == tt.val
		})

		if i != tt.res {
			t.FailNow()
		}
	}
}

func TestAny(t *testing.T) {
	l := NewList[int]()

	l.AddRangeRolled(1, 2, 3, 4, 5, 6)

	tests := []struct {
		val int
		res bool
	}{
		{1, true},
		{2, true},
		{3, true},
		{10, false},
		{6, true},
		{0, false},
	}

	for _, tt := range tests {
		an := l.Any(func(v int) bool {
			return v == tt.val
		})

		if an != tt.res {
			t.FailNow()
		}
	}
}
