package rec

import (
	"testing"

	"github.com/dgi09/serversdk/text"
)

func TestAtleastOneOfClass(t *testing.T) {
	tests := []struct {
		val   string
		class text.CharClass
		res   bool
	}{
		{"", text.CharClass_LCapLetter, false},
		{"t", text.CharClass_LCapLetter, true},
		{"t1", text.CharClass_LCapLetter, true},
		{"A", text.CharClass_LCapLetter, false},
		{"AA", text.CharClass_LCapLetter, false},
		{"Aa", text.CharClass_LCapLetter, true},
		{"", text.CharClass_HCapLetter, false},
		{"T", text.CharClass_HCapLetter, true},
		{"T1", text.CharClass_HCapLetter, true},
		{"a", text.CharClass_HCapLetter, false},
		{"aa", text.CharClass_HCapLetter, false},
		{"aA", text.CharClass_HCapLetter, true},
		{"", text.CharClass_Digit, false},
		{"T", text.CharClass_Digit, false},
		{"T1", text.CharClass_Digit, true},
		{"t1", text.CharClass_Digit, true},
		{"11", text.CharClass_Digit, true},
		{"aa", text.CharClass_Digit, false},
		{"aA", text.CharClass_Digit, false},
	}

	for _, tt := range tests {
		if tt.res != HasAtleastOneOfClass(tt.val, tt.class) {
			t.Error(tt.val)
		}
	}
}
