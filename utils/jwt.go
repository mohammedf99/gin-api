package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const sKey = "supersecret"

func GenerateToken(email string, uId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": uId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(sKey))
}

func VerifyJWTToken(t string) (int64, error) {
	parsedToken, err := jwt.Parse(t, func(t *jwt.Token) (interface{}, error) {
		// _ is just data
		_, ok := t.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("unexpected singing method")
		}
		return []byte(sKey), nil
	})

	if err != nil {
		return 0, errors.New("could not parse token")
	}

	isTokenValid := parsedToken.Valid

	if !isTokenValid {
		return 0, errors.New("invalid token")
	}

	c, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("invalid token claims")
	}

	// email := c["email"].(string)
	uId := int64(c["userId"].(float64))

	return uId, nil
}
