package middlewares

import (
	"goAPI/configs"
	"goAPI/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func ValidateJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		AuthHeader := ctx.Request.Header.Get("Authorization")
		if AuthHeader != "" {
			token, err := jwt.Parse(AuthHeader, func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					utils.ReturnErrorMessage(ctx, utils.HtppError{Message: "Nao Autorizado", HttpCode: 401})
					ctx.Abort()
				}
				return []byte(configs.GetJWTSecret()), nil
			})

			if err != nil {
				utils.ReturnErrorMessage(ctx, utils.HtppError{Message: "Nao Autorizado", HttpCode: 401})
				ctx.Abort()
			}

			if token.Valid {
				ctx.Next()
			}
		} else {
			utils.ReturnErrorMessage(ctx, utils.HtppError{Message: "Nao Autorizado", HttpCode: 401})
			ctx.Abort()
		}
	}
}
