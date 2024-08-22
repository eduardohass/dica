package userAnswer

import (
	"net/http"

	"github.com/eduardohass/dica/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Show userAnswer
// @Description Show a job userAnswer
// @Tags UserAnswer
// @Accept json
// @Produce json
// @Param id query string true "UserAnswer identification"
// @Success 200 {object} ShowUserAnswerResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /userAnswer [get]
func ShowUserAnswerHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	userAnswer := schemas.UserAnswer{}

	if err := db.First(&userAnswer, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "userAnswer not found")
		return
	}
	sendSuccess(ctx, "show-userAnswer", userAnswer)
}
