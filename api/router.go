package api

import (
	"edutest/api/handler"
	"edutest/service"
	"log/slog"

	_ "edutest/api/docs"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
)

func Router(service service.Service, log *slog.Logger) *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	h := handler.NewHandler(service, log)

	students := router.Group("/students")
	{
		students.POST("/create", h.CreateStudent)
		students.PUT("/update/:id", h.UpdateStudent)
		students.DELETE("/delete/:id", h.DeleteStudent)
		students.GET("", h.GetStudents)
	}

	subjects := router.Group("/subjects")
	{
		subjects.POST("/create", h.CerateSubject)
		subjects.DELETE("/delete/:id", h.DeleteSubject)
		subjects.GET("/", h.GetSubjects)
	}

	questions := router.Group("/questions")
	{
		questions.POST("/create", h.CreateQuestion)
		questions.PUT("/update/:id", h.UpdateQuestion)
		questions.DELETE("/delete/:id", h.DeleteQuestion)
		questions.GET("/", h.GetQuestions)
	}

	templates := router.Group("/templates")
	{
		templates.POST("/create", h.CreateTemplate)
		templates.POST("/check", h.CheckStudentTest)
	}

	return router
}
