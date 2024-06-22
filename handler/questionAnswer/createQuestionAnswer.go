package questionAnswer

import (
	"fmt"
	"net/http"

	"github.com/eduardohass/dica/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Create questionAnswer
// @Description Create a new job questionAnswer
// @Tags Questions
// @Accept json
// @Produce json
// @Param request body CreateQuestionAnswerRequest true "Request body"
// @Success 200 {object} CreateQuestionAnswerRequest
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /questionAnswer [post]
func CreateQuestionAnswerHandler(ctx *gin.Context) {
	request := CreateQuestionAnswerRequest{}

	ctx.BindJSON(&request)

	fmt.Println("CreateQuestion - Antes da validação")

	if err := request.Validate(); err != nil {
		fmt.Println("CreateQuestion - Caiu no erro")
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println("CreateQuestionAnswer - Depois da validação")

	questionAnswer := schemas.QuestionAnswer{
		IdQuestion: request.IdQuestion,
		IdAnswer1:  request.IdAnswer1,
		IdAnswer2:  request.IdAnswer2,
		IdAnswer3:  request.IdAnswer3,
		IdAnswer4:  request.IdAnswer4,
	}
	fmt.Println("DBG==depois de popular a instancia de questionAnswer")

	if err := db.Create(&questionAnswer).Error; err != nil {
		fmt.Println("DBG==entrou no metodo de erro quando foi salvar no DB")
		logger.Errorf("error creating questionAnswer: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating questionAnswer on database")
		return
	}
	fmt.Println("DBG==Depois de gravar no DB")

	sendSuccess(ctx, "create-questionAnswer", questionAnswer)
}
