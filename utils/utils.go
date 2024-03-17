package utils

import (
	"encoding/base64"
	"math/rand"
	"strconv"
	"test-task-auth-service-api/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func HashRefreshTokenString(token string) (string, error) {
	ht, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(ht), nil
}

func ValidateGuid(guid string) error {
	// GUID is actually the same 128-bit identifier as UUID but in context of Windows OS
	return uuid.Validate(guid)
}

func GenerateRandomString(size int) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	b := []byte("abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "1234567890")
	s := make([]byte, size)

	for i := range s {
		s[i] = b[r.Intn(len(b))]
	}

	return string(s)
}

func GenerateRefreshTokenString() (string, error) {
	randomString := GenerateRandomString(16)
	timeString := strconv.FormatInt(time.Now().Unix(), 10)

	tokenString := randomString + timeString
	tokenStringBase64 := base64.StdEncoding.EncodeToString([]byte(tokenString))

	return tokenStringBase64, nil
}

func GenerateAccessTokenString(userId string, refreshTokenId string) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"userId":         userId,
		"refreshTokenId": refreshTokenId,
		"iat":            time.Now().Unix(),
		"exp":            time.Now().Add(time.Minute * time.Duration(config.Envs.AccessTokenLifetimeInMinutes)).Unix(),
	})

	tokenString, err := jwtToken.SignedString([]byte(config.Envs.AccessTokenSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
