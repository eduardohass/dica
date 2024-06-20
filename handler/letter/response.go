package letter

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

type CreateLetterResponse struct {
	Message string                 `json:"message"`
	Data    schemas.LetterResponse `json:"data"`
}

type DeleteLetterResponse struct {
	Message string                 `json:"message"`
	Data    schemas.LetterResponse `json:"data"`
}
type ShowLetterResponse struct {
	Message string                 `json:"message"`
	Data    schemas.LetterResponse `json:"data"`
}
type ListLettersResponse struct {
	Message string                   `json:"message"`
	Data    []schemas.LetterResponse `json:"data"`
}
type UpdateLetterResponse struct {
	Message string                 `json:"message"`
	Data    schemas.LetterResponse `json:"data"`
}
