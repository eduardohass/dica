package answer

import (
	"net/http"

	"github.com/eduardohass/dica/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Show answer
// @Description Show a job answer
// @Tags Answers
// @Accept json
// @Produce json
// @Param id query string true "Answer identification"
// @Success 200 {object} ShowAnswerResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /answer [get]
func ShowAnswerHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	answer := schemas.Answer{}

	if err := db.First(&answer, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "answer not found")
		return
	}
	sendSuccess(ctx, "show-answer", answer)
}
