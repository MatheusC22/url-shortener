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
	urlService := services.NewUrlService()

	if err := ctx.BindJSON(&url); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Corpo da Requisição Malformatado!",
			"Error":   err,
		})
	}

	url_hash, err := urlService.Insert(url)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Erro ao inserir o usuario!",
			"Error":   err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "url criada com sucesso!",
		"user_id": url_hash,
	})
}

func (u *urlHandler) GetUrl(ctx *gin.Context) {
	url_hash := ctx.Param("url_hash")
	urlService := services.NewUrlService()
	url, err := urlService.Get(url_hash)

	if err == sql.ErrNoRows {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "URL nao encontrada",
			"Error":   err,
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

	ctx.JSON(http.StatusOK, url)
}
func (u *urlHandler) RedirectToUrl(ctx *gin.Context) {
	url_hash := ctx.Param("url_hash")
	urlService := services.NewUrlService()
	url, err := urlService.Get(url_hash)
	if err == sql.ErrNoRows {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "URL na encotrada",
			"Error":   err,
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
