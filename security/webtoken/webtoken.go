package webtoken

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/o1egl/paseto"
)

func Sign(descr SignWebTokenDescr) (string, error) {
	symmetricKey := []byte(descr.Key)
	now := time.Now()
	exp := now.Add(descr.LifeDuration)

	jti := uuid.NewString()

	jsonToken := paseto.JSONToken{
		Audience:   "service_audiance",
		Issuer:     "service",
		Jti:        jti,
		Subject:    "service_subject",
		IssuedAt:   now,
		Expiration: exp,
		NotBefore:  now,
	}

	for key := range descr.Values {
		jsonToken.Set(key, descr.Values[key])
	}

	return paseto.NewV2().Encrypt(symmetricKey, jsonToken, descr.Footer)
}

func Parse(descr ParseWebTokenDescr) (*WebToken, error) {
	key := []byte(descr.Key)
	var token paseto.JSONToken
	var footer string

	err := paseto.NewV2().Decrypt(descr.StrToken, key, &token, &footer)

	if err != nil {
		return nil, err
	}

	if descr.ValidateExp && token.Expiration.Before(time.Now()) {
		return nil, errors.New("token expired")
	}

	if descr.ValidateFooter && descr.Footer != footer {
		return nil, errors.New("footer mismatch")
	}

	result := &WebToken{
		IssuedAt: token.IssuedAt,
		ExpAt:    token.Expiration,
		Values:   make(map[string]string),
	}

	for _, extrVal := range descr.ExtractValues {
		result.Values[extrVal] = token.Get(extrVal)
	}

	return result, nil
}
