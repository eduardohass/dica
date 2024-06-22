package questionAnswer

import (
	"net/http"

	"github.com/eduardohass/dica/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Show questionAnswer
// @Description Show a job questionAnswer
// @Tags Questions
// @Accept json
// @Produce json
// @Param id query string true "Question identification"
// @Success 200 {object} ShowQuestionResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /questionAnswer [get]
func ShowQuestionAnswerHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	questionAnswer := schemas.Question{}

	if err := db.First(&questionAnswer, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "questionAnswer not found")
		return
	}
	sendSuccess(ctx, "show-questionAnswer", questionAnswer)
}
