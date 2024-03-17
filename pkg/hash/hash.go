package hash

import (
	"golang.org/x/crypto/bcrypt"
)

func HashString(s string) (string, error) {
	hs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hs), nil
}

func CompareHashAndString(hs string, s string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hs), []byte(s))

	return err
}
