package letterAnswer

import (
	"net/http"

	"github.com/eduardohass/dica/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Show letterAnswer
// @Description Show a job letterAnswer
// @Tags LettersAnswer
// @Accept json
// @Produce json
// @Param id query string true "Letter identification"
// @Success 200 {object} ShowLetterAnswerResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /letterAnswer [get]
func ShowLetterAnswerHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	letterAnswer := schemas.LetterAnswer{}

	if err := db.First(&letterAnswer, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "letterAnswer not found")
		return
	}
	sendSuccess(ctx, "show-letterAnswer", letterAnswer)
}
