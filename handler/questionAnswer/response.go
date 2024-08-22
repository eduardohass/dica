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

type CreateQuestionAnswerResponse struct {
	Message string                         `json:"message"`
	Data    schemas.QuestionAnswerResponse `json:"data"`
}

type DeleteQuestionAnswerResponse struct {
	Message string                         `json:"message"`
	Data    schemas.QuestionAnswerResponse `json:"data"`
}
type ShowQuestionAnswerResponse struct {
	Message string                         `json:"message"`
	Data    schemas.QuestionAnswerResponse `json:"data"`
}
type ListQuestionsAnswerResponse struct {
	Message string                           `json:"message"`
	Data    []schemas.QuestionAnswerResponse `json:"data"`
}
type UpdateQuestionAnswerResponse struct {
	Message string                         `json:"message"`
	Data    schemas.QuestionAnswerResponse `json:"data"`
}
