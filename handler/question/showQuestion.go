package question

import (
	"net/http"

	"github.com/eduardohass/dica/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Show question
// @Description Show a job question
// @Tags Questions
// @Accept json
// @Produce json
// @Param id query string true "Question identification"
// @Success 200 {object} ShowQuestionResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /question [get]
func ShowQuestionHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	question := schemas.Question{}

	if err := db.First(&question, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "question not found")
		return
	}
	sendSuccess(ctx, "show-question", question)
}
