package question

import (
	"fmt"
	"net/http"

	"github.com/eduardohass/dica/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Create question
// @Description Create a new job question
// @Tags Questions
// @Accept json
// @Produce json
// @Param request body CreateQuestionRequest true "Request body"
// @Success 200 {object} CreateQuestionResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /question [post]
func CreateQuestionHandler(ctx *gin.Context) {
	request := CreateQuestionRequest{}

	ctx.BindJSON(&request)

	fmt.Println("CreateQuestion - Antes da validação")

	if err := request.Validate(); err != nil {
		fmt.Println("CreateQuestion - Caiu no erro")
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println("CreateQuestion - Depois da validação")

	question := schemas.Question{
		IdQuestion: request.IdQuestion,
		OptionD:    request.OptionD,
		OptionI:    request.OptionI,
		OptionC:    request.OptionC,
		OptionA:    request.OptionA,
	}
	fmt.Println("DBG==depois de popular a instancia de question")

	if err := db.Create(&question).Error; err != nil {
		fmt.Println("DBG==entrou no metodo de erro quando foi salvar no DB")
		logger.Errorf("error creating question: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating question on database")
		return
	}
	fmt.Println("DBG==Depois de gravar no DB")

	sendSuccess(ctx, "create-question", question)
}
