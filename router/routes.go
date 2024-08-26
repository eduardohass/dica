package router

import (
	"github.com/eduardohass/dica/docs"
	handlerAnswer "github.com/eduardohass/dica/handler/answer"
	handlerLetter "github.com/eduardohass/dica/handler/letter"
	handlerLetterAnswer "github.com/eduardohass/dica/handler/letterAnswer"
	handlerOpening "github.com/eduardohass/dica/handler/opening"
	handlerQuestion "github.com/eduardohass/dica/handler/question"
	handlerQuestionAnswer "github.com/eduardohass/dica/handler/questionAnswer"
	handlerUserQuestion "github.com/eduardohass/dica/handler/userQuestion"
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
	handlerLetterAnswer.InitializeHandler()
	handlerUserQuestion.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group("/api/v1")
	{
		// Opening
		v1.GET("openings", handlerOpening.ListOpeningsHandler)
		v1.GET("opening", handlerOpening.ShowOpeningHandler)
		v1.POST("opening", handlerOpening.CreateOpeningHandler)
		v1.PUT("opening", handlerOpening.UpdateOpeningHandler)
		v1.DELETE("opening", handlerOpening.DeleteOpeningHandler)
		// Letter
		v1.GET("letters", handlerLetter.ListLettersHandler)
		v1.GET("letter", handlerLetter.ShowLetterHandler)
		v1.POST("letter", handlerLetter.CreateLetterHandler)
		v1.PUT("letter", handlerLetter.UpdateLetterHandler)
		v1.DELETE("letter", handlerLetter.DeleteLetterHandler)
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
		// QuestionAnswer
		v1.GET("questionsAnswer", handlerQuestionAnswer.ListQuestionsAnswerHandler)
		v1.GET("questionAnswer", handlerQuestionAnswer.ShowQuestionAnswerHandler)
		v1.POST("questionAnswer", handlerQuestionAnswer.CreateQuestionAnswerHandler)
		v1.PUT("questionAnswer", handlerQuestionAnswer.UpdateQuestionAnswerHandler)
		v1.DELETE("questionAnswer", handlerQuestionAnswer.DeleteQuestionAnswerHandler)
		// LetterAnswer
		v1.GET("lettersAnswer", handlerLetterAnswer.ListLettersAnswerHandler)
		v1.GET("letterAnswer", handlerLetterAnswer.ShowLetterAnswerHandler)
		v1.POST("letterAnswer", handlerLetterAnswer.CreateLetterAnswerHandler)
		v1.PUT("letterAnswer", handlerLetterAnswer.UpdateLetterAnswerHandler)
		v1.DELETE("letterAnswer", handlerLetterAnswer.DeleteLetterAnswerHandler)
		// UserQuestion
		v1.GET("usersAnswer", handlerUserQuestion.ListUsersAnswerHandler)
		v1.GET("userQuestion", handlerUserQuestion.ShowUserAnswerHandler)
		v1.POST("userQuestion", handlerUserQuestion.CreateUserAnswerHandler)
		v1.PUT("userQuestion", handlerUserQuestion.UpdateUserAnswerHandler)
		v1.DELETE("userQuestion", handlerUserQuestion.DeleteUserAnswerHandler)
	}
	// Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
