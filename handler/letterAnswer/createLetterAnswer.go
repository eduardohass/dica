package lettersAnswer

import (
	"fmt"
	"net/http"

	"github.com/eduardohass/dica/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Create lettersAnswer
// @Description Create a new job lettersAnswer
// @Tags LettersAnswer
// @Accept json
// @Produce json
// @Param request body CreateLetterAnswerRequest true "Request body"
// @Success 200 {object} CreateLetterAnswerResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /lettersAnswer [post]
func CreateLetterAnswerHandler(ctx *gin.Context) {
	request := CreateLetterAnswerRequest{}

	ctx.BindJSON(&request)

	fmt.Println("CreateLetterAnswer - Antes da validação")

	if err := request.Validate(); err != nil {
		fmt.Println("CreateLetterAnswer - Caiu no erro")
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println("CreateLetterAnswer - Depois da validação")

	lettersAnswer := schemas.LetterAnswer{
		IdLetter: request.IdLetter,
		IdAnswer: request.IdAnswer,
	}
	fmt.Println("DBG==depois de popular a instancia de lettersAnswer")

	if err := db.Create(&lettersAnswer).Error; err != nil {
		fmt.Println("DBG==entrou no metodo de erro quando foi salvar no DB")
		logger.Errorf("error creating lettersAnswer: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating lettersAnswer on database")
		return
	}
	fmt.Println("DBG==Depois de gravar no DB")

	sendSuccess(ctx, "create-lettersAnswer", lettersAnswer)
}
