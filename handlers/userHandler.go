package handlers

import (
	"database/sql"
	"goAPI/models"
	"goAPI/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
}

func NewUserHandler() *userHandler {
	return &userHandler{}
}

func (u *userHandler) CreateUser(ctx *gin.Context) {
	var user models.UserDTO
	userService := services.NewUserService()

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Corpo da Requisição Malformatado!",
			"Error":   err,
		})
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

func (u *userHandler) GetAllUSers(ctx *gin.Context) {
	userService := services.NewUserService()
	users, err := userService.GetAll()
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (u *userHandler) GetUSer(ctx *gin.Context) {
	user_id := ctx.Param("user_id")
	userService := services.NewUserService()
	user, err := userService.Get(user_id)

	if err == sql.ErrNoRows {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Usuario nao encotrado",
			"Error":   err,
		})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Message": "Erro ao buscar o usuario!",
			"Error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
func (u *userHandler) DeleteUSer(ctx *gin.Context) {
	user_id := ctx.Param("user_id")
	userService := services.NewUserService()

	response, err := userService.Delete(user_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Message": "Erro ao deletar o usuario!",
			"Error":   err,
		})
		return
	}

	if response != 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Usuario nao encontrado!",
		})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{
		"Message": "usuario deletado com sucesso!",
	})
}

func (u *userHandler) UpdateUser(ctx *gin.Context) {
	var user models.UserDTO
	user_id := ctx.Param("user_id")
	userService := services.NewUserService()

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Corpo da Requisição Malformatado!",
			"Error":   err,
		})
	}

	resp, err := userService.Update(user_id, user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Message": "Erro ao atualizar o usuario!",
			"Error":   err,
		})
		return
	}

	if resp != 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Usuario nao encontrado!",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message": "usuario atualizado com sucesso!",
	})

}
