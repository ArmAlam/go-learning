package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const SECRET = "SECRET"

func GeneraToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(SECRET))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte(SECRET), nil
	})

	if err != nil {
		return 0, errors.New("Coun't verify token")
	}

	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return 0, errors.New("Token invalid")
	}

	cliams, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("Token invalid")

	}

	// email := cliams["email"].(string)
	userId := int64(cliams["userId"].(float64))

	return userId, nil

}
