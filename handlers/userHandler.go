package handlers

import (
	"goAPI/models"
	"goAPI/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
}

func NewUserHandler() userHandler {
	var uHand userHandler
	return uHand
}

func (u userHandler) CreateUser(ctx *gin.Context) {
	var user models.User
	userService := services.NewUserService()

	if err := ctx.BindJSON(&user); err != nil {
		return
	}

	user_id, err := userService.Insert(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Erro ao inserir o usuario!",
			"Error":   err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "usuario criado com sucesso!",
		"user_id": user_id,
	})
}

func (u userHandler) GetAll(ctx *gin.Context) {
	userService := services.NewUserService()

	teste := userService.Teste()

	ctx.JSON(http.StatusOK, gin.H{
		"teste": teste,
	})
}
