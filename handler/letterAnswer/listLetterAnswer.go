package lettersAnswer

import (
	"fmt"
	"net/http"

	"github.com/eduardohass/dica/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List lettersAnswer
// @Description List all job lettersAnswer
// @Tags LettersAnswer
// @Accept json
// @Produce json
// @Success 200 {object} ListLettersAnswerResponse
// @Failure 500 {object} ErrorResponse
// @Router /lettersAnswer [get]
func ListLettersAnswerHandler(ctx *gin.Context) {
	fmt.Println("DBG==List Letters Answer ")
	lettersAnswer := []schemas.LetterAnswer{}
	fmt.Println("DBG==lettersAnswer: ", lettersAnswer)
	if err := db.Find(&lettersAnswer).Error; err != nil {
		fmt.Println("DBG==Encontrou um erro ao listar as lettersAnswer")
		sendError(ctx, http.StatusInternalServerError, "error listing lettersAnswer")
		return
	}
	sendSuccess(ctx, "list-lettersAnswer", lettersAnswer)
}
