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

type userHandler struct {
	H_queueService services.RabbitmqService
	app            *newrelic.Application
}

func NewUserHandler() *userHandler {
	new_app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("url-shortener"),
		newrelic.ConfigLicense(""),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	if err != nil {
		panic(err)
	}
	queueService := services.NewRabbitMQService()
	return &userHandler{H_queueService: *queueService, app: new_app}
}

func (u *userHandler) CreateUser(ctx *gin.Context) {
	var user models.UserRequest
	userService := services.NewUserService()
	txn := u.app.StartTransaction("CreateUser")
	defer txn.End()

	if err := ctx.BindJSON(&user); err != nil {
		utils.ReturnErrorMessage(ctx, utils.HtppError{Message: "Corpo da Requisicao Malformatado", HttpCode: 400})
		return
	}

	exist, err := userService.UserExists(user.User_email)
	if err != nil {
		utils.ReturnUnexpectedError(ctx, err)
		return
	}
	if exist {
		utils.ReturnErrorMessage(ctx, utils.HtppError{Message: "Usuario ja existe!", HttpCode: 400})
		return
	}

	user_id, err := userService.Insert(user)
	if err != nil {
		utils.ReturnUnexpectedError(ctx, err)
		return
	}
	//succes
	u.H_queueService.Publish(ctx.FullPath() + ";" + ctx.Request.Method)
	ctx.JSON(http.StatusCreated, gin.H{
		"Message": "usuario criado com sucesso!",
		"user_id": user_id,
	})
}

func (u *userHandler) GetAllUSers(ctx *gin.Context) {
	userService := services.NewUserService()
	users, err := userService.GetAll()
	txn := u.app.StartTransaction("GetAllUSers")
	defer txn.End()
	if err != nil {
		utils.ReturnUnexpectedError(ctx, err)
		return
	}
	u.H_queueService.Publish(ctx.FullPath() + ";" + ctx.Request.Method)
	ctx.JSON(http.StatusOK, users)
}

func (u *userHandler) GetUSer(ctx *gin.Context) {
	user_id := ctx.Param("user_id")
	userService := services.NewUserService()
	user, err := userService.GetById(user_id)
	txn := u.app.StartTransaction("GetUser")
	defer txn.End()

	if err == sql.ErrNoRows {
		utils.ReturnErrorMessage(ctx, utils.HtppError{Message: "Usuario nao encontrado", HttpCode: 400})
		return
	}
	if err != nil {
		utils.ReturnUnexpectedError(ctx, err)
		return
	}
	user.User_password = "{removed}"
	u.H_queueService.Publish(ctx.FullPath() + ";" + ctx.Request.Method)
	ctx.JSON(http.StatusOK, user)
}
func (u *userHandler) DeleteUSer(ctx *gin.Context) {
	user_id := ctx.Param("user_id")
	userService := services.NewUserService()
	txn := u.app.StartTransaction("DeleteUser")
	defer txn.End()
	response, err := userService.Delete(user_id)
	if err != nil {
		utils.ReturnUnexpectedError(ctx, err)
		return
	}

	if response != 1 {
		utils.ReturnErrorMessage(ctx, utils.HtppError{Message: "Usuario nao encontrado", HttpCode: 400})
		return
	}
	//success
	u.H_queueService.Publish(ctx.FullPath() + ";" + ctx.Request.Method)
	ctx.JSON(http.StatusNoContent, gin.H{
		"Message": "usuario deletado com sucesso!",
	})
}

func (u *userHandler) UpdateUser(ctx *gin.Context) {
	var user models.UserRequest
	user_id := ctx.Param("user_id")
	userService := services.NewUserService()
	txn := u.app.StartTransaction("UpdateUser")
	defer txn.End()

	if err := ctx.BindJSON(&user); err != nil {
		utils.ReturnErrorMessage(ctx, utils.HtppError{Message: "Corpo da Requisicao Malformatado", HttpCode: 400})
		return
	}

	resp, err := userService.Update(user_id, user)

	if err != nil {
		utils.ReturnUnexpectedError(ctx, err)
		return
	}

	if resp != 1 {
		utils.ReturnErrorMessage(ctx, utils.HtppError{Message: "Usuario nao encontrado", HttpCode: 400})
		return
	}
	//sueccess
	u.H_queueService.Publish(ctx.FullPath() + ";" + ctx.Request.Method)
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "usuario atualizado com sucesso!",
	})

}
func (u *userHandler) Login(ctx *gin.Context) {
	var request models.UserLoginRequest
	userService := services.NewUserService()
	txn := u.app.StartTransaction("Login")
	defer txn.End()

	if err := ctx.BindJSON(&request); err != nil {
		utils.ReturnErrorMessage(ctx, utils.HtppError{Message: "Corpo da Requisicao Malformatado", HttpCode: 400})
		return
	}
	user, err := userService.GetByEmail(request.User_email)
	if err != nil {
		utils.ReturnUnexpectedError(ctx, err)
		return
	}

	if request.User_password == user.User_password {
		token, err := utils.CreateJWT(models.UserJWTPayload{User_id: user.User_id})

		if err != nil {
			utils.ReturnUnexpectedError(ctx, err)
			return
		}
		//success
		u.H_queueService.Publish(ctx.FullPath() + ";" + ctx.Request.Method)
		ctx.Header("Authorization", token)
		ctx.JSON(http.StatusOK, gin.H{
			"Message": "logado com sucesso!",
		})
	}

}
