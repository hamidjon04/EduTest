package api

import (
	"edutest/api/handler"
	"edutest/api/middleware"
	"edutest/pkg/config"
	"edutest/service"
	"log/slog"

	_ "edutest/api/docs"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Tokenni faqat o‘zi yozing, masalan: eyJhbGciOiJIUzI1NiIs..
func Router(service service.Service, log *slog.Logger, cfg config.Config) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	h := handler.NewHandler(service, log)

	router.POST("/register", h.Register)
	router.POST("/login", h.Login)	
	router.GET("/refresh-token", h.RefreshToken)
	

	router.Use(middleware.AuthMiddleware(cfg.JWT_KEY))
	students := router.Group("/students")
	{
		students.POST("/create", h.CreateStudent)
		students.PUT("/update/:id", h.UpdateStudent)
		students.DELETE("/delete/:id", h.DeleteStudent)
		students.GET("", h.GetStudents)
		students.GET("/:student_id/result", h.GetStudentResult)
		students.GET("/results", h.GetStudentsResults)
		students.POST("/upload", h.UploadStudentsExelFile)
	}

	subjects := router.Group("/subjects")
	{
		subjects.POST("/create", h.CerateSubject)
		subjects.DELETE("/update/:id", h.UpdateSubject)
		subjects.GET("/get", h.GetSubjects)
	}

	questions := router.Group("/questions")
	{
		questions.POST("/create", h.CreateQuestion)
		questions.PUT("/update/:id", h.UpdateQuestion)
		questions.DELETE("/delete/:subject_id", h.DeleteQuestion)
		questions.GET("", h.GetQuestions)
		questions.POST("/upload", h.UploadQuestionsExelFile)
		questions.POST("/image/upload", h.UploadFile)
	}

	templates := router.Group("/templates")
	{
		templates.POST("/create", h.CreateTemplate)
		templates.POST("/check", h.CheckStudentTest)
		templates.GET("/get", h.GetStudentTemplate)
	}

	return router
}
