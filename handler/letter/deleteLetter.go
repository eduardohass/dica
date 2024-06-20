package letter

import (
	"fmt"
	"net/http"

	"github.com/eduardohass/dica/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Delete letter
// @Description Delete a job letter
// @Tags Letters
// @Accept json
// @Produce json
// @Param id query string true "Letter identification"
// @Success 200 {object} DeleteLetterResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /letter [delete]
func DeleteLetterHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	letter := schemas.Letter{}

	// Find letter
	if err := db.First(&letter, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("letter with id: %s not found", id))
		return
	}

	// Delete letter
	if err := db.Delete(&letter).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting letter with id: %s", id))
		return
	}
	sendSuccess(ctx, "delete-letter", letter)
}
