package rec

import (
	"regexp"

	"github.com/dgi09/serversdk/text"
)

func HasAtleastOneOfClass(input string, class text.CharClass) bool {
	switch class {
	case text.CharClass_LCapLetter:
		return hasAtleastOneLCapLetter(input)
	case text.CharClass_HCapLetter:
		return hasAtleastOneHCapLetter(input)
	case text.CharClass_Digit:
		return hasAtleasOneDigit(input)
	default:
		return false
	}
}

func hasAtleastOneLCapLetter(input string) bool {
	regex := regexp.MustCompile("[a-z]")

	return regex.MatchString(input)
}

func hasAtleastOneHCapLetter(input string) bool {
	regex := regexp.MustCompile("[A-Z]")

	return regex.MatchString(input)
}

func hasAtleasOneDigit(input string) bool {
	regex := regexp.MustCompile("[0-9]")

	return regex.MatchString(input)
}
