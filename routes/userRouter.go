package routes

import (
	"goAPI/handlers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) *gin.RouterGroup {
	userHandler := handlers.NewUserHandler()

	userRouter := router.Group("/user")
	{
		userRouter.POST("/", userHandler.CreateUser)
		userRouter.GET("/", userHandler.GetAll)
	}
	return userRouter
}
