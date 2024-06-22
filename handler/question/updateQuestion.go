package question

import (
	"fmt"
	"net/http"

	"github.com/eduardohass/dica/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Update question
// @Description Update a job question
// @Tags Questions
// @Accept json
// @Produce json
// @Param id query string true "Question Identification"
// @Param question body UpdateQuestionRequest true "Question data to Update"
// @Success 200 {object} UpdateQuestionResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /question [put]
func UpdateQuestionHandler(ctx *gin.Context) {
	request := UpdateQuestionRequest{}
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

	question := schemas.Question{}

	if err := db.First(&question, id).Error; err != nil {
		fmt.Println("DBG==Erro ao recuperar a question")
		sendError(ctx, http.StatusNotFound, "question not found")
		return
	}
	// Update question
	if request.Question != "" {
		question.Question = request.Question
	}
	// Save question
	if err := db.Save(&question).Error; err != nil {
		logger.Errorf("error updating question: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error update question")
		return
	}
	sendSuccess(ctx, "update-question", question)
}
