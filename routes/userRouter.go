package routes

import (
	"goAPI/handlers"
	"goAPI/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) *gin.RouterGroup {
	userHandler := handlers.NewUserHandler()

	userRouter := router.Group("/user")
	{
		userRouter.POST("/", userHandler.CreateUser)
		userRouter.GET("/", middlewares.ValidateJWT(), userHandler.GetAllUSers)
		userRouter.GET("/:user_id", middlewares.ValidateJWT(), userHandler.GetUSer)
		userRouter.DELETE("/:user_id", middlewares.ValidateJWT(), userHandler.DeleteUSer)
		userRouter.PUT("/:user_id", middlewares.ValidateJWT(), userHandler.UpdateUser)
	}
	return userRouter
}
