package letterAnswer

import (
	"fmt"
	"net/http"

	"github.com/eduardohass/dica/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Update letterAnswer
// @Description Update a job letterAnswer
// @Tags Letters
// @Accept json
// @Produce json
// @Param id query string true "Letter Identification"
// @Param letterAnswer body UpdateLetterAnswerRequest true "Letter data to Update"
// @Success 200 {object} UpdateLetterResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /letterAnswer [put]
func UpdateLetterAnswerHandler(ctx *gin.Context) {
	request := UpdateLetterAnswerRequest{}
	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		fmt.Println("DBG==ID é vazio")
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	letterAnswer := schemas.LetterAnswer{}

	if err := db.First(&letterAnswer, id).Error; err != nil {
		fmt.Println("DBG==Erro ao recuperar a letterAnswer")
		sendError(ctx, http.StatusNotFound, "letterAnswer not found")
		return
	}
	// Update letterAnswer
	if request.IdLetter > 0 {
		letterAnswer.IdLetter = request.IdLetter
	}
	if request.IdAnswer > 0 {
		letterAnswer.IdAnswer = request.IdAnswer
	}
	// Save letterAnswer
	if err := db.Save(&letterAnswer).Error; err != nil {
		logger.Errorf("error updating letterAnswer: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error update letterAnswer")
		return
	}
	sendSuccess(ctx, "update-letterAnswer", letterAnswer)
}
