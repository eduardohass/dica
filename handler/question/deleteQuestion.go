package question

import (
	"fmt"
	"net/http"

	"github.com/eduardohass/dica/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Delete question
// @Description Delete a job question
// @Tags Questions
// @Accept json
// @Produce json
// @Param id query string true "Question identification"
// @Success 200 {object} DeleteQuestionResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /question [delete]
func DeleteQuestionHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	question := schemas.Question{}

	// Find question
	if err := db.First(&question, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("question with id: %s not found", id))
		return
	}

	// Delete question
	if err := db.Delete(&question).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting question with id: %s", id))
		return
	}
	sendSuccess(ctx, "delete-question", question)
}
