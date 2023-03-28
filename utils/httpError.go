package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HtppError struct {
	Message  string
	HttpCode int
}

func ReturnErrorMessage(ctx *gin.Context, err HtppError) {
	ctx.JSON(err.HttpCode, gin.H{
		"Message": err.Message,
	})
}

func ReturnUnexpectedError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"Message": "Um erro inesperado ocorreu",
		"Error":   err.Error(),
	})
}
