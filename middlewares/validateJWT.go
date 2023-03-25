package middlewares

import (
	"fmt"
	"goAPI/configs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func ValidateJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		AuthHeader := ctx.Request.Header.Get("Authorization")
		fmt.Println(AuthHeader)
		if AuthHeader != "" {
			token, err := jwt.Parse(AuthHeader, func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					ctx.JSON(http.StatusUnauthorized, gin.H{
						"Message": "Nao autorizado",
					})
					ctx.Abort()
				}
				return []byte(configs.GetJWTSecret()), nil
			})

			if err != nil {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"Message": "Nao autorizado",
				})
				ctx.Abort()
			}

			if token.Valid {
				ctx.Next()
			}
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"Message": "Nao autorizado",
				"last":    true,
			})
			ctx.Abort()
		}
	}
}
