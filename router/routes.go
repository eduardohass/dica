package router

import (
	docs "github.com/eduardohass/dica/docs"
	handlerOpening "github.com/eduardohass/dica/handler/opening"

	handlerLetter "github.com/eduardohass/dica/handler/letter"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	handlerOpening.InitializeHandler()
	handlerLetter.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group("/api/v1")
	{
		// Opening
		v1.GET("opening", handlerOpening.ShowOpeningHandler)
		v1.POST("opening", handlerOpening.CreateOpeningHandler)
		v1.PUT("opening", handlerOpening.UpdateOpeningHandler)
		v1.DELETE("opening", handlerOpening.DeleteOpeningHandler)
		v1.GET("openings", handlerOpening.ListOpeningsHandler)
		// Letter
		v1.GET("letter", handlerLetter.ShowLetterHandler)
		v1.POST("letter", handlerLetter.CreateLetterHandler)
		v1.PUT("letter", handlerLetter.UpdateLetterHandler)
		v1.DELETE("letter", handlerLetter.DeleteLetterHandler)
		v1.GET("letters", handlerLetter.ListLettersHandler)
	}
	// Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
