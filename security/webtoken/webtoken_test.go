package webtoken

import (
	"testing"
	"time"
)

func TestWebtoken(t *testing.T) {
	key := "YELLOW SUBMARINE, BLACK WIZARDRY"

	values := map[string]string{
		"id":   "100",
		"name": "Dimitar",
	}

	footer := "testfooter"

	signDescr := SignWebTokenDescr{
		Key:          key,
		Values:       values,
		LifeDuration: time.Hour * 10,
		Footer:       footer,
	}

	resToken, err := Sign(signDescr)

	if err != nil {
		t.FailNow()
	}

	parseDescr := ParseWebTokenDescr{
		Key:           key,
		StrToken:      resToken,
		Footer:        footer,
		ExtractValues: []string{"id", "name"},
	}

	webToken, err := Parse(parseDescr)

	if err != nil || webToken == nil {
		t.FailNow()
	}

	if webToken.Values["id"] != "100" || webToken.Values["name"] != "Dimitar" {
		t.FailNow()
	}
}
