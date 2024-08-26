package questionAnswer

import (
	"fmt"
	"net/http"

	"github.com/eduardohass/dica/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Delete questionAnswer
// @Description Delete a job questionAnswer
// @Tags QuestionsAnswer
// @Accept json
// @Produce json
// @Param id query string true "Question identification"
// @Success 200 {object} DeleteQuestionAnswerResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /questionAnswer [delete]
func DeleteQuestionAnswerHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	questionAnswer := schemas.QuestionAnswer{}

	// Find questionAnswer
	if err := db.First(&questionAnswer, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("questionAnswer with id: %s not found", id))
		return
	}

	// Delete questionAnswer
	if err := db.Delete(&questionAnswer).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting questionAnswer with id: %s", id))
		return
	}
	sendSuccess(ctx, "delete-questionAnswer", questionAnswer)
}
