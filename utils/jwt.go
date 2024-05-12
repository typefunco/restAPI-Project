package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey string = "SecretAuthToken"

func GenerateToken(login string, UseriId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login": login,
		"id":    UseriId,
		"exp":   time.Now().Add(time.Hour * 3).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}
