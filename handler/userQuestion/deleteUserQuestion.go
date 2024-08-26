package userQuestion

import (
	"fmt"
	"net/http"

	"github.com/eduardohass/dica/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Delete userQuestion
// @Description Delete a job userQuestion
// @Tags UserQuestion
// @Accept json
// @Produce json
// @Param id query string true "UserQuestion identification"
// @Success 200 {object} DeleteUserAnswerResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /userQuestion [delete]
func DeleteUserAnswerHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	userQuestion := schemas.UserQuestion{}

	// Find userQuestion
	if err := db.First(&userQuestion, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("userQuestion with id: %s not found", id))
		return
	}

	// Delete userQuestion
	if err := db.Delete(&userQuestion).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting userQuestion with id: %s", id))
		return
	}
	sendSuccess(ctx, "delete-userQuestion", userQuestion)
}
