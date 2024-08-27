package userQuestion

import (
	"fmt"
	"net/http"

	"github.com/eduardohass/dica/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Update UserQuestion
// @Description Update a job UserQuestion
// @Tags UserQuestion
// @Accept json
// @Produce json
// @Param id query string true "UserQuestion Identification"
// @Param UserQuestion body UpdateUserQuestionRequest true "UserQuestion data to Update"
// @Success 200 {object} UpdateUserQuestionResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /userQuestion [put]
func UpdateUserQuestionHandler(ctx *gin.Context) {
	request := UpdateUserQuestionRequest{}
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

	UserQuestion := schemas.UserQuestion{}

	if err := db.First(&UserQuestion, id).Error; err != nil {
		fmt.Println("DBG==Erro ao recuperar a UserQuestion")
		sendError(ctx, http.StatusNotFound, "UserQuestion not found")
		return
	}
	// Update UserQuestion
	if request.IdQuestion > 0 {
		UserQuestion.IdQuestion = request.IdQuestion
	}
	if request.IdAnswer1 > 0 {
		UserQuestion.IdAnswer1 = request.IdAnswer1
	}
	if request.IdAnswer2 > 0 {
		UserQuestion.IdAnswer2 = request.IdAnswer2
	}
	if request.IdAnswer3 > 0 {
		UserQuestion.IdAnswer3 = request.IdAnswer3
	}
	if request.IdAnswer4 > 0 {
		UserQuestion.IdAnswer4 = request.IdAnswer4
	}
	// Save UserQuestion
	if err := db.Save(&UserQuestion).Error; err != nil {
		logger.Errorf("error updating UserQuestion: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error update UserQuestion")
		return
	}
	sendSuccess(ctx, "update-UserQuestion", UserQuestion)
}
