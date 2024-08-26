package questionAnswer

import (
	"fmt"
	"net/http"

	"github.com/eduardohass/dica/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List questionsAnswer
// @Description List all job questionsAnswer
// @Tags QuestionsAnswer
// @Accept json
// @Produce json
// @Success 200 {object} ListQuestionsAnswerResponse
// @Failure 500 {object} ErrorResponse
// @Router /questionsAnswer [get]
func ListQuestionsAnswerHandler(ctx *gin.Context) {
	fmt.Println("DBG==List Questions Answer ")
	questionsAnswer := []schemas.QuestionAnswer{}
	fmt.Println("DBG==questionsAnswer: ", questionsAnswer)
	if err := db.Find(&questionsAnswer).Error; err != nil {
		fmt.Println("DBG==Encontrou um erro ao listar as questionsAnswer")
		sendError(ctx, http.StatusInternalServerError, "error listing questionsAnswer")
		return
	}
	sendSuccess(ctx, "list-questionsAnswer", questionsAnswer)
}
