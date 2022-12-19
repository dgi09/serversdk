package pwd

import "golang.org/x/crypto/bcrypt"

func GenerateSaltedPassowrd(rawPwd string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(rawPwd), bcrypt.DefaultCost)
}

func ComparePassowrds(solted []byte, raw string) bool {
	return bcrypt.CompareHashAndPassword(solted, []byte(raw)) == nil
}
