package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const SECRET_KEY = "my_secret_key" // TODO: move to env

func GenerateJWT(email string, userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  "",
		"userId": "",
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString(SECRET_KEY)
}
