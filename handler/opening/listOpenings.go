package opening

import (
	"fmt"
	"net/http"

	"github.com/eduardohass/dica/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List openings
// @Description List all job openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningsResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningsHandler(ctx *gin.Context) {
	fmt.Println("DBG==List Openings")
	openings := []schemas.Opening{}
	fmt.Println("DBG==openings: ", openings)
	if err := db.Find(&openings).Error; err != nil {
		fmt.Println("DBG==Encontrou um erro ao listar os openings")
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}
	sendSuccess(ctx, "list-openings", openings)
}
