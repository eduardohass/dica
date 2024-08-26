package lettersAnswer

import (
	"net/http"

	"github.com/eduardohass/dica/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Show lettersAnswer
// @Description Show a job lettersAnswer
// @Tags LettersAnswer
// @Accept json
// @Produce json
// @Param id query string true "Letter identification"
// @Success 200 {object} ShowLetterAnswerResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /lettersAnswer [get]
func ShowLetterAnswerHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	lettersAnswer := schemas.LetterAnswer{}

	if err := db.First(&lettersAnswer, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "lettersAnswer not found")
		return
	}
	sendSuccess(ctx, "show-lettersAnswer", lettersAnswer)
}
