package routes

import (
	"goAPI/handlers"

	"github.com/gin-gonic/gin"
)

func UrlRoutes(router *gin.Engine) *gin.RouterGroup {
	urlHandler := handlers.NewUrlHandler()

	urlRouter := router.Group("/url")
	{
		urlRouter.POST("/", urlHandler.CreateUrl)
		urlRouter.GET("/:url_hash", urlHandler.GetUrl)
		urlRouter.GET("/redirect/:url_hash", urlHandler.RedirectToUrl)
	}
	return urlRouter
}
