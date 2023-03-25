package handlers

import (
	"database/sql"
	"goAPI/models"
	"goAPI/services"
	"goAPI/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
}

func NewUserHandler() *userHandler {
	return &userHandler{}
}

func (u *userHandler) CreateUser(ctx *gin.Context) {
	var user models.UserRequest
	userService := services.NewUserService()

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Corpo da Requisição Malformatado!",
		})
	}

	user_id, err := userService.Insert(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Message": "Erro ao inserir o usuario!",
			"Error":   err,
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"Message": "usuario criado com sucesso!",
		"user_id": user_id,
	})
}

func (u *userHandler) GetAllUSers(ctx *gin.Context) {
	userService := services.NewUserService()
	users, err := userService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Message": "Erro ao buscar os usuarios!",
			"Error":   err,
		})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (u *userHandler) GetUSer(ctx *gin.Context) {
	user_id := ctx.Param("user_id")
	userService := services.NewUserService()
	user, err := userService.GetById(user_id)

	if err == sql.ErrNoRows {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Usuario nao encotrado",
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
	user.User_password = ""
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
	var user models.UserRequest
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
func (u *userHandler) Login(ctx *gin.Context) {
	var request models.UserLoginRequest
	userService := services.NewUserService()

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Corpo da Requisição Malformatado",
		})
	}
	user, err := userService.GetByEmail(request.User_email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Message": "Erro ao fazer login!",
			"Error":   err,
		})
	}

	if request.User_password == user.User_password {
		token, err := utils.CreateJWT()

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Erro ao fazer login!",
				"Error":   err,
			})
			return
		}

		ctx.Header("Authorization", token)
		ctx.JSON(http.StatusOK, gin.H{
			"Message": "logado com sucesso!",
		})
	}

}
