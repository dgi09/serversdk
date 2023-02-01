package validation

type ITextValidator interface {
	Validate(input string) bool
}
