package text

type Char uint8

type CharClass int8

const (
	CharClass_LCapLetter CharClass = iota
	CharClass_HCapLetter
	CharClass_Digit
)
