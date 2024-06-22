package userAnswer

import (
	"fmt"
	"net/http"

	"github.com/eduardohass/dica/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List userAnswers
// @Description List all job userAnswers
// @Tags UserAnswers
// @Accept json
// @Produce json
// @Success 200 {object} ListUserAnswersResponse
// @Failure 500 {object} ErrorResponse
// @Router /userAnswers [get]
func ListUserAnswersHandler(ctx *gin.Context) {
	fmt.Println("DBG==List UserAnswers ")
	userAnswers := []schemas.UserAnswer{}
	fmt.Println("DBG==userAnswers: ", userAnswers)
	if err := db.Find(&userAnswers).Error; err != nil {
		fmt.Println("DBG==Encontrou um erro ao listar as userAnswers")
		sendError(ctx, http.StatusInternalServerError, "error listing userAnswers")
		return
	}
	sendSuccess(ctx, "list-userAnswers", userAnswers)
}
