package api

import (
	"edutest/api/handler"
	"edutest/api/middleware"
	"edutest/service"
	"log/slog"

	_ "edutest/api/docs"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func Router(service service.Service, log *slog.Logger) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	h := handler.NewHandler(service, log)

	students := router.Group("/students")
	{
		students.POST("/create", h.CreateStudent)
		students.PUT("/update/:id", h.UpdateStudent)
		students.DELETE("/delete/:id", h.DeleteStudent)
		students.GET("", h.GetStudents)
		students.GET("/:student_id/result", h.GetStudentResult)
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
		templates.GET("/get", h.GetStudentTemplate)
	}

	return router
}
