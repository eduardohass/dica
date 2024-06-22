package userAnswer

import (
	"fmt"
	"net/http"

	"github.com/eduardohass/dica/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Create userAnswer
// @Description Create a new job userAnswer
// @Tags UserAnswers
// @Accept json
// @Produce json
// @Param request body CreateUserAnswerRequest true "Request body"
// @Success 200 {object} CreateUserAnswerResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /userAnswer [post]
func CreateUserAnswerHandler(ctx *gin.Context) {
	request := CreateUserAnswerRequest{}

	ctx.BindJSON(&request)

	fmt.Println("CreateUserAnswer - Antes da validação")

	if err := request.Validate(); err != nil {
		fmt.Println("CreateUserAnswer - Caiu no erro")
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println("CreateUserAnswer - Depois da validação")

	userAnswer := schemas.UserAnswer{
		IdQuestion: request.IdQuestion,
		IdAnswer1:  request.IdAnswer1,
		IdAnswer2:  request.IdAnswer2,
		IdAnswer3:  request.IdAnswer3,
		IdAnswer4:  request.IdAnswer4,
	}
	fmt.Println("DBG==depois de popular a instancia de userAnswer")

	if err := db.Create(&userAnswer).Error; err != nil {
		fmt.Println("DBG==entrou no metodo de erro quando foi salvar no DB")
		logger.Errorf("error creating userAnswer: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating userAnswer on database")
		return
	}
	fmt.Println("DBG==Depois de gravar no DB")

	sendSuccess(ctx, "create-userAnswer", userAnswer)
}
