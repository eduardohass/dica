package answer

import (
	"fmt"
	"net/http"

	"github.com/eduardohass/dica/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List answers
// @Description List all job answers
// @Tags Answers
// @Accept json
// @Produce json
// @Success 200 {object} ListAnswersResponse
// @Failure 500 {object} ErrorResponse
// @Router /answers [get]
func ListAnswersHandler(ctx *gin.Context) {
	fmt.Println("DBG==List Answers ")
	answers := []schemas.Answer{}
	fmt.Println("DBG==answers: ", answers)
	if err := db.Find(&answers).Error; err != nil {
		fmt.Println("DBG==Encontrou um erro ao listar as answers")
		sendError(ctx, http.StatusInternalServerError, "error listing answers")
		return
	}
	sendSuccess(ctx, "list-answers", answers)
}
