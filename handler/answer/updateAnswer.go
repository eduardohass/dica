package answer

import (
	"fmt"
	"net/http"

	"github.com/eduardohass/dica/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Update answer
// @Description Update a job answer
// @Tags Answers
// @Accept json
// @Produce json
// @Param idAnswer query string true "Answer Identification"
// @Param answer body UpdateAnswerRequest true "Answer data to Update"
// @Success 200 {object} UpdateAnswerResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /answer [put]
func UpdateAnswerHandler(ctx *gin.Context) {
	request := UpdateAnswerRequest{}
	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	idAnswer := ctx.Query("idAnswer")
	if idAnswer == "" {
		fmt.Println("DBG==idAnswer é vazio")
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("idAnswer", "queryParameter").Error())
		return
	}

	answer := schemas.Answer{}

	if err := db.First(&answer, idAnswer).Error; err != nil {
		fmt.Println("DBG==Erro ao recuperar a answer")
		sendError(ctx, http.StatusNotFound, "answer not found")
		return
	}
	// Update answer
	if request.IdAnswer != 0 {
		answer.IdAnswer = request.IdAnswer
	}
	if request.IdUser != 0 {
		answer.IdUser = request.IdUser
	}
	if request.IdQuestion != 0 {
		answer.IdQuestion = request.IdQuestion
	}
	answer.OptionD = request.OptionD
	answer.OptionI = request.OptionI
	answer.OptionC = request.OptionC
	answer.OptionA = request.OptionA

	// Save answer
	if err := db.Save(&answer).Error; err != nil {
		logger.Errorf("error updating answer: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error update answer")
		return
	}
	sendSuccess(ctx, "update-answer", answer)
}
