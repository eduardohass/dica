package questionAnswer

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

type CreateQuestionResponse struct {
	Message string                   `json:"message"`
	Data    schemas.QuestionResponse `json:"data"`
}

type DeleteQuestionResponse struct {
	Message string                   `json:"message"`
	Data    schemas.QuestionResponse `json:"data"`
}
type ShowQuestionResponse struct {
	Message string                   `json:"message"`
	Data    schemas.QuestionResponse `json:"data"`
}
type ListQuestionsResponse struct {
	Message string                     `json:"message"`
	Data    []schemas.QuestionResponse `json:"data"`
}
type UpdateQuestionResponse struct {
	Message string                   `json:"message"`
	Data    schemas.QuestionResponse `json:"data"`
}
