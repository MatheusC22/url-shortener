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
		userRouter.GET("/", userHandler.GetAllUSers)
		userRouter.GET("/:user_id", middlewares.ValidateJWT(), middlewares.EnsureRightUser(), userHandler.GetUSer)
		userRouter.DELETE("/:user_id", middlewares.ValidateJWT(), middlewares.EnsureRightUser(), userHandler.DeleteUSer)
		userRouter.PUT("/:user_id", middlewares.ValidateJWT(), middlewares.EnsureRightUser(), userHandler.UpdateUser)
	}
	return userRouter
}
