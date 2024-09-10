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
// @Param idAnswer query string true "Answer identification"
// @Success 200 {object} ShowAnswerResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /answer [get]
func ShowAnswerHandler(ctx *gin.Context) {
	idAnswer := ctx.Query("idAnswer")
	if idAnswer == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("idAnswer", "queryParameter").Error())
		return
	}
	answer := schemas.Answer{}

	if err := db.First(&answer, idAnswer).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "answer not found")
		return
	}
	sendSuccess(ctx, "show-answer", answer)
}
