package config

import "testing"

func TestInMemoryReader(t *testing.T) {
	r := NewInMemoryConfigReader()

	r.SetValue("serverUrl", "localhost:5000")
	r.SetValue("secretKey", "key")
	r.SetValue("testInt", "30")
	r.SetValue("testBool", "false")

	i, pst := r.Int32("testInt", -1)

	if i != 30 || !pst {
		t.FailNow()
	}

	i, pst = r.Int32("noExt", -1)

	if i != -1 || pst {
		t.FailNow()
	}

	s, pst := r.String("secretKey", "")

	if s != "key" || !pst {
		t.FailNow()
	}

	b, pst := r.Bool("testBool", false)

	if b != false || !pst {
		t.FailNow()
	}
}
