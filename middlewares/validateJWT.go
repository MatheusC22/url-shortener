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
			claim, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				utils.ReturnErrorMessage(ctx, utils.HtppError{Message: "Payload nao encontrado", HttpCode: 401})
				ctx.Abort()
			}
			//setting the user_id to the context
			ctx.Set("user_id_payload", claim["sub"].(string))

			if token.Valid {
				ctx.Next()
			}
		} else {
			utils.ReturnErrorMessage(ctx, utils.HtppError{Message: "Nao Autorizado", HttpCode: 401})
			ctx.Abort()
		}
	}
}
