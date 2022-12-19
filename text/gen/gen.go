package gen

import (
	"math/rand"

	"github.com/dgi09/serversdk/text"
)

func Random(lenght int, classes ...text.CharClass) string {
	buffer := make([]text.Char, lenght)

	for i := 0; i < lenght; i++ {
		class := rand.Intn(len(classes))
		buffer[i] = RandomChar(text.CharClass(class))
	}

	return string(buffer)
}

func RandomChar(class text.CharClass) text.Char {
	switch class {
	case text.CharClass_LCapLetter:
		return randomLCapLetter()
	case text.CharClass_HCapLetter:
		return randomHCapLetter()
	case text.CharClass_Digit:
		return randomDigit()
	default:
		return 0
	}
}

func randomLCapLetter() text.Char {
	return 97 + text.Char(rand.Intn(26))
}

func randomHCapLetter() text.Char {
	return 65 + text.Char(rand.Intn(26))
}

func randomDigit() text.Char {
	return 48 + text.Char(rand.Intn(10))
}
