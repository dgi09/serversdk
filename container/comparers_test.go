package container

import (
	"testing"
	"time"
)

func TestIntComparer(t *testing.T) {
	tests := []struct {
		a, b, res int
	}{
		{10, 20, -1},
		{20, 10, 1},
		{10, 10, 0},
	}

	for _, tt := range tests {
		if tt.res != int(IntComparer(tt.a, tt.b)) {
			t.FailNow()
		}
	}
}

func TestStringComparer(t *testing.T) {
	tests := []struct {
		a, b string
		res  int8
	}{
		{"10", "20", -1},
		{"20", "10", 1},
		{"10", "10", 0},
	}

	for _, tt := range tests {
		if tt.res != StringComparer(tt.a, tt.b) {
			t.FailNow()
		}
	}
}

func TestTimeComparer(t *testing.T) {
	now := time.Now()

	tests := []struct {
		a, b time.Time
		res  int8
	}{
		{now, now.Add(time.Second * 3), -1},
		{now, now.Add(time.Second * -3), 1},
		{now, now, 0},
	}

	for _, tt := range tests {
		if tt.res != TimeComparer(tt.a, tt.b) {
			t.FailNow()
		}
	}
}
