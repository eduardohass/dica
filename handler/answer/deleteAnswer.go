package answer

import (
	"fmt"
	"net/http"

	"github.com/eduardohass/dica/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Delete answer
// @Description Delete a job answer
// @Tags Answers
// @Accept json
// @Produce json
// @Param id query string true "Answer identification"
// @Success 200 {object} DeleteAnswerResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /answer [delete]
func DeleteAnswerHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	answer := schemas.Answer{}

	// Find answer
	if err := db.First(&answer, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("answer with id: %s not found", id))
		return
	}

	// Delete answer
	if err := db.Delete(&answer).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting answer with id: %s", id))
		return
	}
	sendSuccess(ctx, "delete-answer", answer)
}
