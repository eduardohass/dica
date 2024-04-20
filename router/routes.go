package router

import (
	"github.com/eduardohass/dica/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("openning", handler.ShowOpeningHandler)
		v1.POST("openning", handler.CreateOpeningHandler)
		v1.PUT("openning", handler.UpdateOpeningHandler)
		v1.DELETE("openning", handler.DeleteOpeningHandler)
		v1.GET("opennings", handler.ListOpeningsHandler)
	}
}
