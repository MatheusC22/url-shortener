package utils

import (
	"goAPI/configs"
	"goAPI/models"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateJWT(payload models.UserJWTPayload) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour).Unix()
	claims["user_id"] = payload.User_id
	var key = []byte(configs.GetJWTSecret())

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
