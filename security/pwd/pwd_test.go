package pwd

import "testing"

func TestEq(t *testing.T) {
	pwd1 := "pwd"
	pwd2 := "pwd"

	gen1, err := GenerateSaltedPassowrd(pwd1)

	if err != nil || gen1 == nil {
		t.FailNow()
	}

	if !ComparePassowrds(gen1, pwd2) {
		t.FailNow()
	}
}

func TestNotEq(t *testing.T) {
	pwd1 := "pwd1"
	pwd2 := "pwd2"

	gen1, err := GenerateSaltedPassowrd(pwd1)

	if err != nil || gen1 == nil {
		t.FailNow()
	}

	if ComparePassowrds(gen1, pwd2) {
		t.FailNow()
	}
}
