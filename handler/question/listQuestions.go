package question

import (
	"fmt"
	"net/http"

	"github.com/eduardohass/dica/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List questions
// @Description List all job questions
// @Tags Questions
// @Accept json
// @Produce json
// @Success 200 {object} ListQuestionsResponse
// @Failure 500 {object} ErrorResponse
// @Router /questions [get]
func ListQuestionsHandler(ctx *gin.Context) {
	fmt.Println("DBG==List Questions ")
	questions := []schemas.Question{}
	fmt.Println("DBG==questions: ", questions)
	if err := db.Find(&questions).Error; err != nil {
		fmt.Println("DBG==Encontrou um erro ao listar as questions")
		sendError(ctx, http.StatusInternalServerError, "error listing questions")
		return
	}
	sendSuccess(ctx, "list-questions", questions)
}
