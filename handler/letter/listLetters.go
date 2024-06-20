package letter

import (
	"fmt"
	"net/http"

	"github.com/eduardohass/dica/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List letters
// @Description List all job letters
// @Tags Letters
// @Accept json
// @Produce json
// @Success 200 {object} ListLettersResponse
// @Failure 500 {object} ErrorResponse
// @Router /letters [get]
func ListLettersHandler(ctx *gin.Context) {
	fmt.Println("DBG==List Letters ")
	letters := []schemas.Letter{}
	fmt.Println("DBG==letters: ", letters)
	if err := db.Find(&letters).Error; err != nil {
		fmt.Println("DBG==Encontrou um erro ao listar as letters")
		sendError(ctx, http.StatusInternalServerError, "error listing letters")
		return
	}
	sendSuccess(ctx, "list-letters", letters)
}
