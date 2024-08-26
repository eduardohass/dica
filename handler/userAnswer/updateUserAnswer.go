package userAnswer

import (
	"fmt"
	"net/http"

	"github.com/eduardohass/dica/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Update UserAnswer
// @Description Update a job UserAnswer
// @Tags UserAnswer
// @Accept json
// @Produce json
// @Param id query string true "UserAnswer Identification"
// @Param UserAnswer body UpdateUserAnswerRequest true "UserAnswer data to Update"
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

	UserAnswer := schemas.UserAnswer{}

	if err := db.First(&UserAnswer, id).Error; err != nil {
		fmt.Println("DBG==Erro ao recuperar a UserAnswer")
		sendError(ctx, http.StatusNotFound, "UserAnswer not found")
		return
	}
	// Update UserAnswer
	if request.IdQuestion > 0 {
		UserAnswer.IdQuestion = request.IdQuestion
	}
	if request.IdAnswer1 > 0 {
		UserAnswer.IdAnswer1 = request.IdAnswer1
	}
	if request.IdAnswer2 > 0 {
		UserAnswer.IdAnswer2 = request.IdAnswer2
	}
	if request.IdAnswer3 > 0 {
		UserAnswer.IdAnswer3 = request.IdAnswer3
	}
	if request.IdAnswer4 > 0 {
		UserAnswer.IdAnswer4 = request.IdAnswer4
	}
	// Save UserAnswer
	if err := db.Save(&UserAnswer).Error; err != nil {
		logger.Errorf("error updating UserAnswer: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error update UserAnswer")
		return
	}
	sendSuccess(ctx, "update-UserAnswer", UserAnswer)
}
