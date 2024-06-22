package questionAnswer

import (
	"fmt"
	"net/http"

	"github.com/eduardohass/dica/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Update questionAnswer
// @Description Update a job questionAnswer
// @Tags Questions
// @Accept json
// @Produce json
// @Param id query string true "Question Identification"
// @Param questionAnswer body UpdateQuestionAnswerRequest true "Question data to Update"
// @Success 200 {object} UpdateQuestionResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /questionAnswer [put]
func UpdateQuestionAnswerHandler(ctx *gin.Context) {
	request := UpdateQuestionAnswerRequest{}
	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	questionAnswer := schemas.QuestionAnswer{}

	if err := db.First(&questionAnswer, id).Error; err != nil {
		fmt.Println("DBG==Erro ao recuperar a questionAnswer")
		sendError(ctx, http.StatusNotFound, "questionAnswer not found")
		return
	}
	// Update questionAnswer
	if request.IdQuestion > 0 {
		questionAnswer.IdQuestion = request.IdQuestion
	}
	if request.IdAnswer1 > 0 {
		questionAnswer.IdAnswer1 = request.IdAnswer1
	}
	if request.IdAnswer2 > 0 {
		questionAnswer.IdAnswer2 = request.IdAnswer2
	}
	if request.IdAnswer3 > 0 {
		questionAnswer.IdAnswer3 = request.IdAnswer3
	}
	if request.IdAnswer4 > 0 {
		questionAnswer.IdAnswer4 = request.IdAnswer4
	}
	// Save questionAnswer
	if err := db.Save(&questionAnswer).Error; err != nil {
		logger.Errorf("error updating questionAnswer: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error update questionAnswer")
		return
	}
	sendSuccess(ctx, "update-questionAnswer", questionAnswer)
}
