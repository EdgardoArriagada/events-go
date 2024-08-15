package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const SECRET_KEY = "my_secret_key" // TODO: move to env

func GenerateJWT(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  "",
		"userId": "",
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(SECRET_KEY))
}

func VerifyToken(token string) error {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return errors.New("Could not parse token.")
	}
	if !parsedToken.Valid {
		return errors.New("Invalid token.")
	}

	/* claims, ok := parsedToken.Claims.(jwt.MapClaims) */
	/* if !ok { */
	/* 	return errors.New("Could not get claims.") */
	/* } */
	/**/
	/* email := claims["email"].(string) */
	/* userId := claims["userId"].(int64) */

	return nil
}
