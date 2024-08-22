package letterAnswer

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

type CreateLetterAnswerResponse struct {
	Message string                       `json:"message"`
	Data    schemas.LetterAnswerResponse `json:"data"`
}

type DeleteLetterAnswerResponse struct {
	Message string                       `json:"message"`
	Data    schemas.LetterAnswerResponse `json:"data"`
}
type ShowLetterAnswerResponse struct {
	Message string                       `json:"message"`
	Data    schemas.LetterAnswerResponse `json:"data"`
}
type ListLettersAnswerResponse struct {
	Message string                         `json:"message"`
	Data    []schemas.LetterAnswerResponse `json:"data"`
}
type UpdateLetterAnswerResponse struct {
	Message string                       `json:"message"`
	Data    schemas.LetterAnswerResponse `json:"data"`
}
