package webtoken

import "time"

type SignWebTokenDescr struct {
	LifeDuration time.Duration
	Key          string
	Values       map[string]string
	Footer       string
}

type ParseWebTokenDescr struct {
	Key            string
	StrToken       string
	Footer         string
	ValidateExp    bool
	ValidateFooter bool
	ExtractValues  []string
}

type WebToken struct {
	IssuedAt time.Time
	ExpAt    time.Time
	Values   map[string]string
}
