package userAnswer

import (
	"fmt"
	"net/http"

	"github.com/eduardohass/dica/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List usersAnswer
// @Description List all job usersAnswer
// @Tags UserAnswer
// @Accept json
// @Produce json
// @Success 200 {object} ListUsersAnswerResponse
// @Failure 500 {object} ErrorResponse
// @Router /usersAnswer [get]
func ListUsersAnswerHandler(ctx *gin.Context) {
	fmt.Println("DBG==List UsersAnswer ")
	usersAnswer := []schemas.UserAnswer{}
	fmt.Println("DBG==UsersAnswer: ", usersAnswer)
	if err := db.Find(&usersAnswer).Error; err != nil {
		fmt.Println("DBG==Encontrou um erro ao listar as UsersAnswer")
		sendError(ctx, http.StatusInternalServerError, "error listing UsersAnswer")
		return
	}
	sendSuccess(ctx, "list-UsersAnswer", usersAnswer)
}
