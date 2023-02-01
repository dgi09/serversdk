package validation

import (
	"github.com/dgi09/serversdk/text"
	"github.com/dgi09/serversdk/text/rec"
)

type MinLenTextValidator struct {
	minLen int
}

func NewMinLengthTextValidator(len int) *MinLenTextValidator {
	return &MinLenTextValidator{
		minLen: len,
	}
}

func (v *MinLenTextValidator) Validate(input string) bool {
	return len(input) >= v.minLen
}

// ========================================================================

type MaxLenTextValidator struct {
	maxLen int
}

func NewMaxLengthTextValidator(len int) *MaxLenTextValidator {
	return &MaxLenTextValidator{
		maxLen: len,
	}
}

func (v *MaxLenTextValidator) Validate(input string) bool {
	return len(input) <= v.maxLen
}

// =========================================================================

type OneOfCharClassTextValidator struct {
	class text.CharClass
}

func NewOneOfCharClassTextValidator(class text.CharClass) *OneOfCharClassTextValidator {
	return &OneOfCharClassTextValidator{
		class: class,
	}
}

func (v *OneOfCharClassTextValidator) Validate(input string) bool {
	return rec.HasAtleastOneOfClass(input, v.class)
}

// ==========================================================================

type CompositeTextValidator struct {
	validators []ITextValidator
}

func NewCompositeTextValidator(validators ...ITextValidator) *CompositeTextValidator {
	return &CompositeTextValidator{
		validators: validators,
	}
}

func (v *CompositeTextValidator) Validate(input string) bool {
	for _, val := range v.validators {
		res := val.Validate(input)

		if !res {
			return false
		}
	}

	return true
}
