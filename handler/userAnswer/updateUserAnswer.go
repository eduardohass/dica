package userAnswer

import (
	"fmt"
	"net/http"

	"github.com/eduardohass/dica/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Update userAnswer
// @Description Update a job userAnswer
// @Tags UserAnswers
// @Accept json
// @Produce json
// @Param id query string true "UserAnswer Identification"
// @Param userAnswer body UpdateUserAnswerRequest true "UserAnswer data to Update"
// @Success 200 {object} UpdateUserAnswerResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /userAnswer [put]
func UpdateUserAnswerHandler(ctx *gin.Context) {
	request := UpdateUserAnswerRequest{}
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

	userAnswer := schemas.UserAnswer{}

	if err := db.First(&userAnswer, id).Error; err != nil {
		fmt.Println("DBG==Erro ao recuperar a questionAnswer")
		sendError(ctx, http.StatusNotFound, "questionAnswer not found")
		return
	}
	// Update questionAnswer
	if request.IdQuestion > 0 {
		userAnswer.IdQuestion = request.IdQuestion
	}
	if userAnswer.IdAnswer1 > 0 {
		userAnswer.IdAnswer1 = request.IdAnswer1
	}
	if request.IdAnswer2 > 0 {
		userAnswer.IdAnswer2 = request.IdAnswer2
	}
	if request.IdAnswer3 > 0 {
		userAnswer.IdAnswer3 = request.IdAnswer3
	}
	if request.IdAnswer4 > 0 {
		userAnswer.IdAnswer4 = request.IdAnswer4
	}
	// Save questionAnswer
	if err := db.Save(&userAnswer).Error; err != nil {
		logger.Errorf("error updating questionAnswer: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error update questionAnswer")
		return
	}
	sendSuccess(ctx, "update-userAnswer", userAnswer)
}
