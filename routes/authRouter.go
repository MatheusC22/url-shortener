package routes

import (
	"goAPI/handlers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) *gin.RouterGroup {
	userHandler := handlers.NewUserHandler()
	gin.SetMode(gin.ReleaseMode)
	authRouter := router.Group("/login")
	{
		authRouter.POST("/", userHandler.Login)
	}
	return authRouter
}
