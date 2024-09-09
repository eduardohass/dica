package router

import (
	"github.com/eduardohass/dica/docs"
	handlerAnswer "github.com/eduardohass/dica/handler/answer"
	handlerQuestion "github.com/eduardohass/dica/handler/question"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	handlerAnswer.InitializeHandler()
	handlerQuestion.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group("/api/v1")
	{
		// Question
		v1.GET("questions", handlerQuestion.ListQuestionsHandler)
		v1.GET("question", handlerQuestion.ShowQuestionHandler)
		v1.POST("question", handlerQuestion.CreateQuestionHandler)
		v1.PUT("question", handlerQuestion.UpdateQuestionHandler)
		v1.DELETE("question", handlerQuestion.DeleteQuestionHandler)
		// Answer
		v1.GET("answers", handlerAnswer.ListAnswersHandler)
		v1.GET("answer", handlerAnswer.ShowAnswerHandler)
		v1.POST("answer", handlerAnswer.CreateAnswerHandler)
		v1.PUT("answer", handlerAnswer.UpdateAnswerHandler)
		v1.DELETE("answer", handlerAnswer.DeleteAnswerHandler)
	}
	// Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
