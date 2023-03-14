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
		userRouter.GET("/", userHandler.GetAllUSers)
		userRouter.GET("/:user_id", userHandler.GetUSer)
		userRouter.DELETE("/:user_id", userHandler.DeleteUSer)
		userRouter.PUT("/:user_id", userHandler.UpdateUser)
	}
	return userRouter
}
