package userAnswer

import (
	"fmt"
	"net/http"

	"github.com/eduardohass/dica/schemas"
	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successful", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreateUserAnswerResponse struct {
	Message string                     `json:"message"`
	Data    schemas.UserAnswerResponse `json:"data"`
}

type DeleteUserAnswerResponse struct {
	Message string                     `json:"message"`
	Data    schemas.UserAnswerResponse `json:"data"`
}
type ShowUserAnswerResponse struct {
	Message string                     `json:"message"`
	Data    schemas.UserAnswerResponse `json:"data"`
}
type ListUserAnswersResponse struct {
	Message string                       `json:"message"`
	Data    []schemas.UserAnswerResponse `json:"data"`
}
type UpdateUserAnswerResponse struct {
	Message string                     `json:"message"`
	Data    schemas.UserAnswerResponse `json:"data"`
}
