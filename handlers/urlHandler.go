package handlers

import (
	"database/sql"
	"goAPI/models"
	"goAPI/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type urlHandler struct {
}

func NewUrlHandler() *urlHandler {
	return &urlHandler{}
}

func (u *urlHandler) CreateUrl(ctx *gin.Context) {
	var url models.UrlDTO
	redisService := services.NewRedisService()
	urlService := services.NewUrlService()

	if err := ctx.BindJSON(&url); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Corpo da Requisição Malformatado!",
			"Error":   err,
		})
	}

	url_hash, err := urlService.Insert(url)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Message": "Erro ao inserir a url!",
			"Error":   err,
		})
		return
	}

	redisService.Set(url_hash, url.Url_original) // INSERT URL IN CACHE
	ctx.JSON(http.StatusOK, gin.H{
		"Message":  "url criada com sucesso!",
		"url_hash": url_hash,
	})
}

func (u *urlHandler) GetUrl(ctx *gin.Context) {
	url_hash := ctx.Param("url_hash")
	redisService := services.NewRedisService()
	urlService := services.NewUrlService()

	//cache
	cache_val, err := redisService.Get(url_hash) // GET URL FROM CACHE
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"cache": true,
			"url":   cache_val,
		})
		return
	}

	//database
	url, err := urlService.Get(url_hash)
	if err == sql.ErrNoRows {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "URL nao encontrada",
		})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Message": "Erro ao buscar a URL!",
			"Error":   err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"db":  true,
		"url": url,
	})
}
func (u *urlHandler) RedirectToUrl(ctx *gin.Context) {
	url_hash := ctx.Param("url_hash")
	redisService := services.NewRedisService()
	urlService := services.NewUrlService()

	cache_val, err := redisService.Get(url_hash)
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"cache": true,
			"url":   cache_val,
		})
		return
	}
	url, err := urlService.Get(url_hash)
	if err == sql.ErrNoRows {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "URL nao encotrada",
		})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Message": "Erro ao buscar a URL!",
			"Error":   err,
		})
		return
	}
	ctx.Redirect(http.StatusMovedPermanently, url)
}

func (u *urlHandler) DeleteUrl(ctx *gin.Context) {
	url_hash := ctx.Param("url_hash")
	redisService := services.NewRedisService()
	urlService := services.NewUrlService()

	response, err := urlService.Delete(url_hash)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Message": "Erro ao deletar URL!",
			"Error":   err,
		})
		return
	}
	if response != 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "URL nao encontrada!",
		})
		return
	}
	redisService.Del(url_hash) // DELETE URL FROM CACHE
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "URL deletada com sucesso!",
	})
}
