package utils

import (
	"errors"
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

func VerifyToken(token string) (int, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("token not valid")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, errors.New("Couldn't parse token")
	}

	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return 0, errors.New("Invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("Invalid claims")
	}

	// login := claims["login"].(string)
	userId := int(claims["id"].(float64)) // actual data of type float64 and I convert data to int

	return userId, nil
}
