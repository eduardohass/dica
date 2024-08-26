package userQuestion

import (
	"fmt"
	"net/http"

	"github.com/eduardohass/dica/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List usersAnswer
// @Description List all job usersAnswer
// @Tags UserQuestion
// @Accept json
// @Produce json
// @Success 200 {object} ListUsersQuestionResponse
// @Failure 500 {object} ErrorResponse
// @Router /usersAnswer [get]
func ListUsersQuestionHandler(ctx *gin.Context) {
	fmt.Println("DBG==List UsersQuestion ")
	usersAnswer := []schemas.UserQuestion{}
	fmt.Println("DBG==UsersQuestion: ", usersAnswer)
	if err := db.Find(&usersAnswer).Error; err != nil {
		fmt.Println("DBG==Encontrou um erro ao listar as UsersQuestion")
		sendError(ctx, http.StatusInternalServerError, "error listing UsersQuestion")
		return
	}
	sendSuccess(ctx, "list-UsersQuestion", usersAnswer)
}
