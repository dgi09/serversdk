package validation

import (
	"testing"

	"github.com/dgi09/serversdk/text"
)

func TestMinLenghtValidator(t *testing.T) {
	tests := []struct {
		input    string
		reqLen   int
		expected bool
	}{
		{"", 5, false},
		{"a", 5, false},
		{"aaaa", 5, false},
		{"aaaaa", 5, true},
		{"aaaaaa", 5, true},
	}

	for _, tt := range tests {
		val := NewMinLengthTextValidator(tt.reqLen)
		res := val.Validate(tt.input)

		if res != tt.expected {
			t.FailNow()
		}
	}
}

func TestMaxLenghtValidator(t *testing.T) {
	tests := []struct {
		input    string
		reqLen   int
		expected bool
	}{
		{"", 5, true},
		{"a", 5, true},
		{"aaaa", 5, true},
		{"aaaaa", 5, true},
		{"aaaaaa", 5, false},
	}

	for _, tt := range tests {
		val := NewMaxLengthTextValidator(tt.reqLen)
		res := val.Validate(tt.input)

		if res != tt.expected {
			t.FailNow()
		}
	}
}

func TestOneOfCharClassValidator(t *testing.T) {

	empty := NewCompositeTextValidator()

	minLen5 := NewCompositeTextValidator(
		NewMinLengthTextValidator(5),
	)

	maxLen10 := NewCompositeTextValidator(
		NewMaxLengthTextValidator(10),
	)

	minMaxLHD := NewCompositeTextValidator(
		minLen5,
		maxLen10,
		NewOneOfCharClassTextValidator(text.CharClass_Digit),
		NewOneOfCharClassTextValidator(text.CharClass_LCapLetter),
		NewOneOfCharClassTextValidator(text.CharClass_HCapLetter),
	)

	tests := []struct {
		input     string
		validator *CompositeTextValidator
		expected  bool
	}{
		{"", empty, true},
		{"", minLen5, false},
		{"aaaaa", minLen5, true},
		{"", maxLen10, true},
		{"aaaaaaaaaaaaaaaaaaaaa", maxLen10, false},
		{"", minMaxLHD, false},
		{"aaa", minMaxLHD, false},
		{"aaaaa", minMaxLHD, false},
		{"aaaaaaaaaaaaaaaaaaaaaa", minMaxLHD, false},
		{"aaaaaaa1", minMaxLHD, false},
		{"11111111", minMaxLHD, false},
		{"AAAAAAAA", minMaxLHD, false},
		{"AAAA111", minMaxLHD, false},
		{"Aaaaa12", minMaxLHD, true},
	}

	for _, tt := range tests {
		res := tt.validator.Validate(tt.input)

		if res != tt.expected {
			t.FailNow()
		}
	}
}

func TestCompositeTextValidator(t *testing.T) {

}
