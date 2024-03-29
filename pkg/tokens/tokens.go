package tokens

import (
	"encoding/base64"
	"math/rand"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func generateRandomString(size int) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	b := []byte("abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "1234567890")
	s := make([]byte, size)

	for i := range s {
		s[i] = b[r.Intn(len(b))]
	}

	return string(s)
}

func GenerateRefreshTokenString() (string, error) {
	randomString := generateRandomString(16)
	timeString := strconv.FormatInt(time.Now().Unix(), 10)

	tokenString := randomString + timeString
	tokenStringBase64 := base64.StdEncoding.EncodeToString([]byte(tokenString))

	return tokenStringBase64, nil
}

func GenerateAccessTokenString(userId string, refreshTokenId string, lifetimeInMinutes int, secret string) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"userId":         userId,
		"refreshTokenId": refreshTokenId,
		"iat":            time.Now().Unix(),
		"exp":            time.Now().Add(time.Minute * time.Duration(lifetimeInMinutes)).Unix(),
	})

	tokenString, err := jwtToken.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func DecodeJwtToken(tokenString string, secret string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		// SigningMethodHMAC implements the HMAC-SHA and SigningMethodHS512 its specific instance
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret), nil
	})

	if err != nil {
		v, _ := err.(*jwt.ValidationError)

		// ignore expiration when decoding
		if v.Errors == jwt.ValidationErrorExpired {
			return token, nil
		}

		return nil, err
	}

	return token, nil
}
