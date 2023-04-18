package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func EnsureRightUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user_id_param := ctx.Param("user_id")
		user_id_token, _ := ctx.Get("user_id_payload")

		if user_id_param == user_id_token {
			ctx.Next()
		} else {
			ctx.JSON(http.StatusForbidden, gin.H{
				"message": "Resource Belong to another User!",
			})
			ctx.Abort()
		}

	}
}
