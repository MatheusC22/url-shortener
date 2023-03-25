package routes

import (
	"goAPI/handlers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) *gin.RouterGroup {
	userHandler := handlers.NewUserHandler()

	authRouter := router.Group("/login")
	{
		authRouter.POST("/", userHandler.Login)
	}
	return authRouter
}
