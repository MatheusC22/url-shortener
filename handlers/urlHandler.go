package handlers

import (
	"database/sql"
	"goAPI/models"
	"goAPI/services"
	"goAPI/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type urlHandler struct {
	app            *newrelic.Application
	H_queueService services.RabbitmqService
	H_cacheService services.RedisService
}

func NewUrlHandler() *urlHandler {
	queueService := services.NewRabbitMQService()
	redisService := services.NewRedisService()
	new_app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("url-shortener"),
		newrelic.ConfigLicense(""),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	if err != nil {
		panic(err)
	}
	return &urlHandler{H_queueService: *queueService, H_cacheService: *redisService, app: new_app}
}

func (u *urlHandler) CreateUrl(ctx *gin.Context) {
	var url models.UrlCreateRequest
	id, _ := ctx.Get("user_id_payload")
	url.User_id = id.(string)
	urlService := services.NewUrlService()
	txn := u.app.StartTransaction("CreateUrl")
	defer txn.End()

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
	u.H_queueService.Publish(ctx.FullPath() + ";" + ctx.Request.Method)
	u.H_cacheService.Set(url_hash, url.Url_original) // INSERT URL IN CACHE
	ctx.JSON(http.StatusCreated, gin.H{
		"Message":  "url criada com sucesso!",
		"url_hash": url_hash,
	})
}
func (u *urlHandler) GetAll(ctx *gin.Context) {
	urlService := services.NewUrlService()
	txn := u.app.StartTransaction("GetAllUrl")
	defer txn.End()

	urls, err := urlService.GetAll()
	if err != nil {
		utils.ReturnUnexpectedError(ctx, err)
		return
	}
	//success
	u.H_queueService.Publish(ctx.FullPath() + ";" + ctx.Request.Method)
	ctx.JSON(http.StatusOK, urls)
}

func (u *urlHandler) GetUrlByHash(ctx *gin.Context) {
	url_hash := ctx.Param("url_hash")
	urlService := services.NewUrlService()
	txn := u.app.StartTransaction("GetUrlByHash")
	defer txn.End()

	//cache
	cache_val, err := u.H_cacheService.Get(url_hash) // GET URL FROM CACHE
	if err == nil {
		//success cache
		u.H_queueService.Publish(ctx.FullPath() + ";" + ctx.Request.Method)
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
	//success
	u.H_queueService.Publish(ctx.FullPath() + ";" + ctx.Request.Method)
	ctx.JSON(http.StatusOK, gin.H{
		"db":  true,
		"url": url,
	})
}
func (u *urlHandler) RedirectToUrl(ctx *gin.Context) {
	url_hash := ctx.Param("url_hash")
	urlService := services.NewUrlService()
	txn := u.app.StartTransaction("RedirectToUrl")
	defer txn.End()

	//cache
	cache_val, err := u.H_cacheService.Get(url_hash)
	if err == nil {
		//success
		u.H_queueService.Publish(ctx.FullPath() + ";" + ctx.Request.Method)
		ctx.Redirect(http.StatusMovedPermanently, cache_val)
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
	//sucess
	u.H_queueService.Publish(ctx.FullPath() + ";" + ctx.Request.Method)
	ctx.Redirect(http.StatusMovedPermanently, url)
}

func (u *urlHandler) DeleteUrl(ctx *gin.Context) {
	url_hash := ctx.Param("url_hash")
	urlService := services.NewUrlService()
	txn := u.app.StartTransaction("DeleteUrl")
	defer txn.End()

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
	u.H_queueService.Publish(ctx.FullPath() + ";" + ctx.Request.Method)
	u.H_cacheService.Del(url_hash) // DELETE URL FROM CACHE
	ctx.JSON(http.StatusNoContent, gin.H{
		"Message": "URL deletada com sucesso!",
	})
}
