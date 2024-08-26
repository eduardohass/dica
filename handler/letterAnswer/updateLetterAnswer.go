package lettersAnswer

import (
	"fmt"
	"net/http"

	"github.com/eduardohass/dica/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Update lettersAnswer
// @Description Update a job lettersAnswer
// @Tags LettersAnswer
// @Accept json
// @Produce json
// @Param id query string true "Letter Identification"
// @Param lettersAnswer body UpdateLetterAnswerRequest true "Letter data to Update"
// @Success 200 {object} UpdateLetterAnswerResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /lettersAnswer [put]
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

	lettersAnswer := schemas.LetterAnswer{}

	if err := db.First(&lettersAnswer, id).Error; err != nil {
		fmt.Println("DBG==Erro ao recuperar a lettersAnswer")
		sendError(ctx, http.StatusNotFound, "lettersAnswer not found")
		return
	}
	// Update lettersAnswer
	if request.IdLetter > 0 {
		lettersAnswer.IdLetter = request.IdLetter
	}
	if request.IdAnswer > 0 {
		lettersAnswer.IdAnswer = request.IdAnswer
	}
	// Save lettersAnswer
	if err := db.Save(&lettersAnswer).Error; err != nil {
		logger.Errorf("error updating lettersAnswer: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error update lettersAnswer")
		return
	}
	sendSuccess(ctx, "update-lettersAnswer", lettersAnswer)
}
