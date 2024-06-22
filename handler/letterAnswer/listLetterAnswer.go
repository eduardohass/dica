package letterAnswer

import (
	"fmt"
	"net/http"

	"github.com/eduardohass/dica/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List letterAnswers
// @Description List all job letterAnswers
// @Tags Letters
// @Accept json
// @Produce json
// @Success 200 {object} ListLettersResponse
// @Failure 500 {object} ErrorResponse
// @Router /letterAnswers [get]
func ListLettersAnswerHandler(ctx *gin.Context) {
	fmt.Println("DBG==List LettersAnswer ")
	letterAnswers := []schemas.LetterAnswer{}
	fmt.Println("DBG==letterAnswers: ", letterAnswers)
	if err := db.Find(&letterAnswers).Error; err != nil {
		fmt.Println("DBG==Encontrou um erro ao listar as letterAnswers")
		sendError(ctx, http.StatusInternalServerError, "error listing letterAnswers")
		return
	}
	sendSuccess(ctx, "list-letterAnswers", letterAnswers)
}
