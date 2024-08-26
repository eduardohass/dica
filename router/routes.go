package router

import (
	docs "github.com/eduardohass/dica/docs"
	handlerAnswer "github.com/eduardohass/dica/handler/answer"
	handlerLetter "github.com/eduardohass/dica/handler/letter"
	handlerOpening "github.com/eduardohass/dica/handler/opening"
	handlerQuestion "github.com/eduardohass/dica/handler/question"
	handlerQuestionAnswer "github.com/eduardohass/dica/handler/questionanswer"
	handlerUserAnswer "github.com/eduardohass/dica/handler/useranswer"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	handlerOpening.InitializeHandler()
	handlerLetter.InitializeHandler()
	handlerAnswer.InitializeHandler()
	handlerQuestion.InitializeHandler()
	handlerQuestionAnswer.InitializeHandler()
	handlerUserAnswer.InitializeHandler()
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
		// // Question
		v1.GET("question", handlerQuestion.ShowQuestionHandler)
		v1.POST("question", handlerQuestion.CreateQuestionHandler)
		v1.PUT("question", handlerQuestion.UpdateQuestionHandler)
		v1.DELETE("question", handlerQuestion.DeleteQuestionHandler)
		v1.GET("questions", handlerQuestion.ListQuestionsHandler)
		// QuestionAnswer
		v1.GET("questionanswer", handlerQuestionAnswer.ShowQuestionAnswerHandler)
		v1.POST("questionanswer", handlerQuestionAnswer.CreateQuestionAnswerHandler)
		v1.PUT("questionanswer", handlerQuestionAnswer.UpdateQuestionAnswerHandler)
		v1.DELETE("questionanswer", handlerQuestionAnswer.DeleteQuestionAnswerHandler)
		v1.GET("questionanswers", handlerQuestionAnswer.ListQuestionsAnswerHandler)
		// // Answer
		v1.GET("answer", handlerAnswer.ShowAnswerHandler)
		v1.POST("answer", handlerAnswer.CreateAnswerHandler)
		v1.PUT("answer", handlerAnswer.UpdateAnswerHandler)
		v1.DELETE("answer", handlerAnswer.DeleteAnswerHandler)
		v1.GET("answers", handlerAnswer.ListAnswersHandler)
		// UserAnswer
		v1.GET("useranswer", handlerUserAnswer.ShowUserAnswerHandler)
		v1.POST("useranswer", handlerUserAnswer.CreateUserAnswerHandler)
		v1.PUT("useranswer", handlerUserAnswer.UpdateUserAnswerHandler)
		v1.DELETE("useranswer", handlerUserAnswer.DeleteUserAnswerHandler)
		v1.GET("useranswers", handlerUserAnswer.ListUsersAnswerHandler)
	}
	// Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
