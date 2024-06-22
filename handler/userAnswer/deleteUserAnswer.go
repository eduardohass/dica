package userAnswer

import (
	"fmt"
	"net/http"

	"github.com/eduardohass/dica/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Delete userAnswer
// @Description Delete a job userAnswer
// @Tags UserAnswers
// @Accept json
// @Produce json
// @Param id query string true "UserAnswer identification"
// @Success 200 {object} DeleteUserAnswerResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /userAnswer [delete]
func DeleteUserAnswerHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	userAnswer := schemas.UserAnswer{}

	// Find userAnswer
	if err := db.First(&userAnswer, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("userAnswer with id: %s not found", id))
		return
	}

	// Delete userAnswer
	if err := db.Delete(&userAnswer).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting userAnswer with id: %s", id))
		return
	}
	sendSuccess(ctx, "delete-userAnswer", userAnswer)
}
