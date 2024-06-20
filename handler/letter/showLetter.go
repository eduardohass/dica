package letter

import (
	"net/http"

	"github.com/eduardohass/dica/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Show letter
// @Description Show a job letter
// @Tags Letters
// @Accept json
// @Produce json
// @Param id query string true "Letter identification"
// @Success 200 {object} ShowLetterResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /letter [get]
func ShowLetterHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	letter := schemas.Letter{}

	if err := db.First(&letter, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "letter not found")
		return
	}
	sendSuccess(ctx, "show-letter", letter)
}
