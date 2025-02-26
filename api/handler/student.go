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
