package gen

import (
	"testing"

	"github.com/dgi09/serversdk/text"
)

func TestRandomChar(t *testing.T) {
	tests := []struct {
		class    text.CharClass
		min, max text.Char
	}{
		{text.CharClass_LCapLetter, 97, 97 + 25},
		{text.CharClass_HCapLetter, 65, 65 + 25},
		{text.CharClass_Digit, 48, 48 + 9},
	}

	for _, tt := range tests {
		for i := 0; i < 10_000_000; i++ {
			res := RandomChar(tt.class)
			if res < tt.min || res > tt.max {
				t.Error(tt)
			}
		}
	}
}

func TestRandom(t *testing.T) {
	tests := []struct {
		len int
	}{
		{1}, {2},
		{10},
		{5},
		{100},
	}

	for _, tt := range tests {
		res := Random(tt.len, text.CharClass_LCapLetter, text.CharClass_HCapLetter, text.CharClass_Digit)

		if len(res) != tt.len {
			t.Error(res)
		}
	}
}
