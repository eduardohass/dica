package answer

import (
	"fmt"
	"net/http"

	"github.com/eduardohass/dica/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Create answer
// @Description Create a new job answer
// @Tags Answers
// @Accept json
// @Produce json
// @Param request body CreateAnswerRequest true "Request body"
// @Success 200 {object} CreateAnswerResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /answer [post]
func CreateAnswerHandler(ctx *gin.Context) {
	request := CreateAnswerRequest{}

	ctx.BindJSON(&request)

	fmt.Println("CreateAnswer - Antes da validação")

	if err := request.Validate(); err != nil {
		fmt.Println("CreateAnswer - Caiu no erro")
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println("CreateAnswer - Depois da validação")

	answer := schemas.Answer{
		IdAnswer:   request.IdAnswer,
		IdUser:     request.IdUser,
		IdQuestion: request.IdQuestion,
		OptionD:    request.OptionD,
		OptionI:    request.OptionI,
		OptionC:    request.OptionC,
		OptionA:    request.OptionA,
	}
	fmt.Println("DBG==depois de popular a instancia de answer")

	if err := db.Create(&answer).Error; err != nil {
		fmt.Println("DBG==entrou no metodo de erro quando foi salvar no DB")
		logger.Errorf("error creating answer: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating answer on database")
		return
	}
	fmt.Println("DBG==Depois de gravar no DB")

	sendSuccess(ctx, "create-answer", answer)
}
