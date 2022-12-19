package validation

import "testing"

func TestIsEmail(t *testing.T) {
	valid := IsEmail("email@mail.ext")

	if !valid {
		t.FailNow()
	}

	valid = IsEmail("@mail.ext")

	if valid {
		t.FailNow()
	}
}
