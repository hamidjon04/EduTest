package handler

import (
	"edutest/pkg/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Talaba yaratish
// @Description Yangi talabani ro‘yxatga olish
// @Tags Students
// @Security     BearerAuth
// @Accept json
// @Produce json
// @Param request body model.CreateStudentReq true "Talaba ma'lumotlari"
// @Success 200 {object} model.CreateStudentResp
// @Failure 400 {object} model.Error "Noto'g'ri ma'lumot kiritildi"
// @Failure 500 {object} model.Error "Serverda xatolik yuz berdi"
// @Router /students/create [post]
func (h *Handler) CreateStudent(c *gin.Context) {
	req := model.CreateStudentReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Error is get data: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{Message: "Noto'g'ri ma'lumot kiritildi"})
		return
	}
	resp, err := h.Service.CreateStudent(c, &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Error is service function CreateStudent: %v", err))
		c.JSON(http.StatusInternalServerError, model.Error{Message: "Serverda xatolik"})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Talaba ma'lumotlarini yangilash
// @Description Berilgan ID bo‘yicha talabani yangilash
// @Tags Students
// @Security     BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "Talabaning unikal ID'si"
// @Param request body model.UpdateStudentReq true "Yangilangan talaba ma'lumotlari"
// @Success 200 {object} model.Status "Muvaffaqiyatli yangilandi"
// @Failure 400 {object} model.Error "Noto'g'ri ma'lumot kiritildi"
// @Failure 500 {object} model.Error "Serverda xatolik yuz berdi"
// @Router /students/update/{id} [put]
func (h *Handler) UpdateStudent(c *gin.Context) {
	req := model.UpdateStudentReq{
		Id: c.Param("id"),
	}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Error is get data: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{Message: "Noto'g'ri ma'lumot kiritildi"})
		return
	}
	err = h.Service.UpdateStudent(c, &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Error is service function UpdateStudent: %v", err))
		c.JSON(http.StatusInternalServerError, model.Error{Message: "Serverda xatolik"})
		return
	}
	c.JSON(http.StatusOK, model.Status{Message: "Success!"})
}

// @Summary Talabani o'chirish
// @Description Berilgan ID bo‘yicha talabani o‘chirish
// @Tags Students
// @Security     BearerAuth
// @Param id path string true "Talabaning unikal ID'si"
// @Success 200 {object} model.Status "Muvaffaqiyatli o‘chirildi"
// @Failure 500 {object} model.Error "Serverda xatolik yuz berdi"
// @Router /students/delete/{id} [delete]
func (h *Handler) DeleteStudent(c *gin.Context) {
	req := model.StudentId{
		Id: c.Param("id"),
	}
	err := h.Service.DeleteStudent(c, &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Error is service function DeleteStudent: %v", err))
		c.JSON(http.StatusInternalServerError, model.Error{Message: "Serverda xatolik"})
		return
	}
	c.JSON(http.StatusOK, model.Status{Message: "Success!"})
}

// @Summary Talabalar ma'lumotlarini olish
// @Description Berilgan ID bo‘yicha talabani yoki barcha talabalarni olish
// @Tags Students
// @Security     BearerAuth
// @Param id query string false "Talabaning unikal ID'si (Majburiy emas, agar berilmasa barcha talabalar qaytariladi)"
// @Success 200 {object} model.GetStudentsResp "Talabalar ma'lumotlari"
// @Failure 500 {object} model.Error "Serverda xatolik yuz berdi"
// @Router /students [get]
func (h *Handler) GetStudents(c *gin.Context) {
	req := model.StudentId{
		Id: c.Query("id"),
	}
	resp, err := h.Service.GetStudents(c, &req)
	if err != nil{
		h.Log.Error(fmt.Sprintf("Error is service function GetStudents: %v", err))
		c.JSON(http.StatusInternalServerError, model.Error{Message: "Serverda xatolik"})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Get student result
// @Description Retrieves the result of a student based on student_id and template_id
// @Tags Students
// @Security     BearerAuth
// @Accept  json
// @Produce  json
// @Param student_id path string true "Student ID"
// @Param template_id query string false "Template ID"
// @Success 200 {object} model.GetStudentResultResp
// @Failure 500 {object} model.Error "Serverda xatolik"
// @Router /students/{student_id}/result [get]
func (h *Handler) GetStudentResult(c *gin.Context){
	req := model.GetStudentResultReq{
		StudentId: c.Param("student_id"),
		TemplateId: c.Query("template_id"),
	}
	resp, err := h.Service.GetStudentResult(c, &req)
	if err != nil{
		h.Log.Error(fmt.Sprintf("Error is service function GetStudentResult: %v", err))
		c.JSON(http.StatusInternalServerError, model.Error{
			Message: "Serverda xatolik",
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Talabalarning natijalarini olish
// @Description Berilgan kun va fanlarga mos keluvchi talabalarning natijalarini qaytaradi
// @Tags Students
// @Security     BearerAuth
// @Accept json
// @Produce json
// @Param day query string false "Natija kuni"
// @Param subject1_id query string false "Birinchi fan ID si"
// @Param subject2_id query string false "Ikkinchi fan ID si"
// @Success 200 {object} model.GetStudentsResultResp "Talabalarning natijalari"
// @Failure 500 {object} model.Error "Serverda xatolik yuz berdi"
// @Router /students/results [get]
func (h *Handler) GetStudentsResults(c *gin.Context){
	req := model.GetStudentsResultReq{
		Day: c.Query("day"),
		Subject1: c.Query("subject1_id"),
		Subject2: c.Query("subject2_id"),
	}
	resp, err := h.Service.GetStudentsResult(c, &req)
	if err != nil{
		h.Log.Error(fmt.Sprintf("Error is service function GetStudentsResult: %v", err))
		c.JSON(http.StatusInternalServerError, model.Error{
			Message: "Serverda xatolik",
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Talabalar ro'yxatini yuklash
// @Description Excel fayl orqali talabalar ma'lumotlarini yuklash
// @Tags Files
// @Security     BearerAuth
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Excel fayl (XLSX format)"
// @Success 200 {object} model.StudentsStatus "Talabalar muvaffaqiyatli yuklandi"
// @Failure 400 {object} model.Error "Fayl yuklashda xatolik"
// @Failure 500 {object} model.Error "Server xatosi"
// @Router /students/upload [post]
func (h *Handler) UploadStudentsExelFile(c *gin.Context){
	file, err := c.FormFile("file")
	if err != nil{
		h.Log.Error(fmt.Sprintf("Error is upload file: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Faylni yuklab olishda xatolik",
			Error: err.Error(),
		})
	}

	filePath := "./" + file.Filename
	if err = c.SaveUploadedFile(file, filePath); err != nil{
		h.Log.Error(fmt.Sprintf("Error is save file: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Faylni saqlashda xatolik",
			Error: err.Error(),
		})
	}

	resp, err := h.Service.OpenStudentsExelFile(c, filePath)
	if err != nil{
		h.Log.Error(fmt.Sprintf("Error is service function OpenStudentsExelFile: %v", err))
		c.JSON(http.StatusInternalServerError, model.Error{
			Message: "Serverda xatolik",
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}