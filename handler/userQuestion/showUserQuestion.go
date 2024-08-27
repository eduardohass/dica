package userQuestion

import (
	"net/http"

	"github.com/eduardohass/dica/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Show userQuestion
// @Description Show a job userQuestion
// @Tags UserQuestion
// @Accept json
// @Produce json
// @Param id query string true "UserQuestion identification"
// @Success 200 {object} ShowUserQuestionResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /userQuestion [get]
func ShowUserQuestionHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	userQuestion := schemas.UserQuestion{}

	if err := db.First(&userQuestion, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "userQuestion not found")
		return
	}
	sendSuccess(ctx, "show-userQuestion", userQuestion)
}
