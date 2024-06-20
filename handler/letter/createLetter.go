package letter

import (
	"fmt"
	"net/http"

	"github.com/eduardohass/dica/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Create letter
// @Description Create a new job letter
// @Tags Letters
// @Accept json
// @Produce json
// @Param request body CreateLetterRequest true "Request body"
// @Success 200 {object} CreateLetterResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /letter [post]
func CreateLetterHandler(ctx *gin.Context) {
	request := CreateLetterRequest{}

	ctx.BindJSON(&request)

	fmt.Println("CreateLetter - Antes da validação")
	fmt.Println("DBG==type " + request.Type)

	if err := request.Validate(); err != nil {
		fmt.Println("CreateLetter - Caiu no erro")
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println("CreateLetter - Depois da validação")
	fmt.Println("DBG==type " + request.Type)

	letter := schemas.Letter{
		Type: request.Type,
	}
	fmt.Println("DBG==depois de popular a instancia de letter")
	fmt.Println("DBG== " + letter.Type)

	if err := db.Create(&letter).Error; err != nil {
		fmt.Println("DBG==entrou no metodo de erro quando foi salvar no DB")
		logger.Errorf("error creating letter: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating letter on database")
		return
	}
	fmt.Println("DBG==Depois de gravar no DB")

	sendSuccess(ctx, "create-letter", letter)
}
