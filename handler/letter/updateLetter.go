package letter

import (
	"fmt"
	"net/http"

	"github.com/eduardohass/dica/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Update letter
// @Description Update a job letter
// @Tags Letters
// @Accept json
// @Produce json
// @Param id query string true "Letter Identification"
// @Param letter body UpdateLetterRequest true "Letter data to Update"
// @Success 200 {object} UpdateLetterResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /letter [put]
func UpdateLetterHandler(ctx *gin.Context) {
	request := UpdateLetterRequest{}
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

	letter := schemas.Letter{}

	if err := db.First(&letter, id).Error; err != nil {
		fmt.Println("DBG==Erro ao recuperar a letter")
		sendError(ctx, http.StatusNotFound, "letter not found")
		return
	}
	// Update letter
	if request.Type != "" {
		letter.Type = request.Type
	}
	// Save letter
	if err := db.Save(&letter).Error; err != nil {
		logger.Errorf("error updating letter: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error update letter")
		return
	}
	sendSuccess(ctx, "update-letter", letter)
}
