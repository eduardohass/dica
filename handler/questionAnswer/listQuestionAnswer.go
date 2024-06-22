package questionAnswer

import (
	"fmt"
	"net/http"

	"github.com/eduardohass/dica/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List questionAnswers
// @Description List all job questionAnswers
// @Tags Questions
// @Accept json
// @Produce json
// @Success 200 {object} ListQuestionsResponse
// @Failure 500 {object} ErrorResponse
// @Router /questionAnswers [get]
func ListQuestionsAnswerHandler(ctx *gin.Context) {
	fmt.Println("DBG==List Questions ")
	questionAnswers := []schemas.Question{}
	fmt.Println("DBG==questionAnswers: ", questionAnswers)
	if err := db.Find(&questionAnswers).Error; err != nil {
		fmt.Println("DBG==Encontrou um erro ao listar as questionAnswers")
		sendError(ctx, http.StatusInternalServerError, "error listing questionAnswers")
		return
	}
	sendSuccess(ctx, "list-questionAnswers", questionAnswers)
}
