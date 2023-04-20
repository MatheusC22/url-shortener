package routes

import (
	"goAPI/handlers"
	"goAPI/middlewares"

	"github.com/gin-gonic/gin"
)

func UrlRoutes(router *gin.Engine) *gin.RouterGroup {
	urlHandler := handlers.NewUrlHandler()

	urlRouter := router.Group("/url")
	{
		urlRouter.POST("/", middlewares.ValidateJWT(), urlHandler.CreateUrl)
		urlRouter.GET("/:url_hash", urlHandler.GetUrlByHash)
		urlRouter.GET("/", urlHandler.GetAll)
		urlRouter.GET("/redirect/:url_hash", urlHandler.RedirectToUrl)
		urlRouter.DELETE(("/:url_hash"), middlewares.ValidateJWT(), urlHandler.DeleteUrl)
	}
	return urlRouter
}
