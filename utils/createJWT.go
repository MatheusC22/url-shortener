package utils

import (
	"goAPI/configs"
	"goAPI/models"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateJWT(payload models.UserJWTPayload) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour).Unix(),
		Subject:   payload.User_id,
	})
	var key = []byte(configs.GetJWTSecret())

	return token.SignedString(key)
}
