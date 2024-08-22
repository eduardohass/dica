package letterAnswer

import (
	"fmt"
	"net/http"

	"github.com/eduardohass/dica/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Delete letterAnswer
// @Description Delete a job letterAnswer
// @Tags LettersAnswer
// @Accept json
// @Produce json
// @Param id query string true "Letter identification"
// @Success 200 {object} DeleteLetterAnswerResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /letterAnswer [delete]
func DeleteLetterAnswerHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	letterAnswer := schemas.LetterAnswer{}

	// Find letterAnswer
	if err := db.First(&letterAnswer, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("letterAnswer with id: %s not found", id))
		return
	}

	// Delete letterAnswer
	if err := db.Delete(&letterAnswer).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting letterAnswer with id: %s", id))
		return
	}
	sendSuccess(ctx, "delete-letterAnswer", letterAnswer)
}
