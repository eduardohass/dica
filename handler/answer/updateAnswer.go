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
// @Param id query string true "Answer Identification"
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

	id := ctx.Query("id")
	if id == "" {
		fmt.Println("DBG==ID é vazio")
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	answer := schemas.Answer{}

	if err := db.First(&answer, id).Error; err != nil {
		fmt.Println("DBG==Erro ao recuperar a answer")
		sendError(ctx, http.StatusNotFound, "answer not found")
		return
	}
	// Update answer
	if request.Answer != "" {
		answer.Answer = request.Answer
	}
	// Save answer
	if err := db.Save(&answer).Error; err != nil {
		logger.Errorf("error updating answer: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error update answer")
		return
	}
	sendSuccess(ctx, "update-answer", answer)
}
