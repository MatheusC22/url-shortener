package handlers

import (
	"database/sql"
	"goAPI/models"
	"goAPI/services"
	"goAPI/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type urlHandler struct {
}

func NewUrlHandler() *urlHandler {
	return &urlHandler{}
}

func (u *urlHandler) CreateUrl(ctx *gin.Context) {
	var url models.UrlCreateRequest
	id, _ := ctx.Get("user_id_payload")
	url.User_id = id.(string)
	redisService := services.NewRedisService()
	urlService := services.NewUrlService()

	if err := ctx.BindJSON(&url); err != nil {
		utils.ReturnErrorMessage(ctx, utils.HtppError{Message: "Corpo da Requisicao Malformatado", HttpCode: 400})
		return
	}

	url_hash, err := urlService.Insert(url)
	if err != nil {
		utils.ReturnUnexpectedError(ctx, err)
		return
	}
	//success
	redisService.Set(url_hash, url.Url_original) // INSERT URL IN CACHE
	ctx.JSON(http.StatusCreated, gin.H{
		"Message":  "url criada com sucesso!",
		"url_hash": url_hash,
	})
}
func (u *urlHandler) GetAll(ctx *gin.Context) {
	urlService := services.NewUrlService()

	urls, err := urlService.GetAll()
	if err != nil {
		utils.ReturnUnexpectedError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, urls)
}

func (u *urlHandler) GetUrlByHash(ctx *gin.Context) {
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
	url, err := urlService.GetByHash(url_hash)
	if err == sql.ErrNoRows {
		utils.ReturnErrorMessage(ctx, utils.HtppError{Message: "Url nao encontrada", HttpCode: 400})
		return
	}
	if err != nil {
		utils.ReturnUnexpectedError(ctx, err)
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
		ctx.Redirect(http.StatusMovedPermanently, cache_val)
		return
	}
	url, err := urlService.GetByHash(url_hash)
	if err == sql.ErrNoRows {
		utils.ReturnErrorMessage(ctx, utils.HtppError{Message: "Url nao encontrada", HttpCode: 400})
		return
	}
	if err != nil {
		utils.ReturnUnexpectedError(ctx, err)
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
		utils.ReturnUnexpectedError(ctx, err)
		return
	}
	if response != 1 {
		utils.ReturnErrorMessage(ctx, utils.HtppError{Message: "Url nao encontrada", HttpCode: 400})
		return
	}
	//success
	redisService.Del(url_hash) // DELETE URL FROM CACHE
	ctx.JSON(http.StatusNoContent, gin.H{
		"Message": "URL deletada com sucesso!",
	})
}
