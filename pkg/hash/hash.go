package hash

import (
	"golang.org/x/crypto/bcrypt"
)

func HashRefreshTokenString(token string) (string, error) {
	ht, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(ht), nil
}
