package lettersAnswer

import (
	"fmt"
	"net/http"

	"github.com/eduardohass/dica/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Delete lettersAnswer
// @Description Delete a job lettersAnswer
// @Tags LettersAnswer
// @Accept json
// @Produce json
// @Param id query string true "Letter identification"
// @Success 200 {object} DeleteLetterAnswerResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /lettersAnswer [delete]
func DeleteLetterAnswerHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	lettersAnswer := schemas.LetterAnswer{}

	// Find lettersAnswer
	if err := db.First(&lettersAnswer, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("lettersAnswer with id: %s not found", id))
		return
	}

	// Delete lettersAnswer
	if err := db.Delete(&lettersAnswer).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting lettersAnswer with id: %s", id))
		return
	}
	sendSuccess(ctx, "delete-lettersAnswer", lettersAnswer)
}
