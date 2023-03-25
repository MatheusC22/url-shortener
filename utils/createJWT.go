package utils

import (
	"goAPI/configs"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour).Unix()
	var key = []byte(configs.GetJWTSecret())

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
