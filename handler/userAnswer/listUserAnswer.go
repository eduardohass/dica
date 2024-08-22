package userAnswer

import (
	"fmt"
	"net/http"

	"github.com/eduardohass/dica/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List UserAnswers
// @Description List all job UserAnswers
// @Tags UsersAnswer
// @Accept json
// @Produce json
// @Success 200 {object} ListUsersAnswerResponse
// @Failure 500 {object} ErrorResponse
// @Router /UserAnswers [get]
func ListUsersAnswerHandler(ctx *gin.Context) {
	fmt.Println("DBG==List UsersAnswer ")
	UserAnswers := []schemas.UserAnswer{}
	fmt.Println("DBG==UserAnswers: ", UserAnswers)
	if err := db.Find(&UserAnswers).Error; err != nil {
		fmt.Println("DBG==Encontrou um erro ao listar as UserAnswers")
		sendError(ctx, http.StatusInternalServerError, "error listing UserAnswers")
		return
	}
	sendSuccess(ctx, "list-UserAnswers", UserAnswers)
}
