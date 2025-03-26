package handler

import (
	"edutest/pkg/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Template yaratish
// @Description Yangi test template yaratish
// @Tags Templates
// @Accept json
// @Produce json
// @Param template body model.CreateTemplateReq true "Template ma'lumotlari"
// @Success 200 {object} model.Status "Muvaffaqiyatli yaratildi"
// @Failure 400 {object} model.Error "Noto'g'ri ma'lumot kiritildi"
// @Failure 500 {object} model.Error "Serverda xatolik"
// @Router /templates/create [post]
func (h *Handler) CreateTemplate(c *gin.Context) {
	req := model.CreateTemplateReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Error is get data: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{Message: "Noto'g'ri ma'lumot kiritildi"})
		return
	}
	err = h.Service.CreateTemplate(c, &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Error is service function CreateTemplate: %v", err))
		c.JSON(http.StatusInternalServerError, model.Error{Message: "Serverda xatolik"})
		return
	}
	c.JSON(http.StatusOK, model.Status{Message: "Success!"})
}

// @Summary      Talabaning test natijalarini tekshirish
// @Description  Ushbu endpoint talabaga tegishli test javoblarini tekshirish uchun ishlatiladi
// @Tags         Templates
// @Accept       json
// @Produce      json
// @Param        request  body  model.CheckStudentTestReq  true  "Talabaning test javoblarini tekshirish uchun ma'lumot"
// @Success      200  {object}  model.Result  "Test natijalari muvaffaqiyatli tekshirildi"
// @Failure      400  {object}  model.Error  "Noto'g'ri ma'lumot kiritildi"
// @Failure      500  {object}  model.Error  "Serverda xatolik yuz berdi"
// @Router       /templates/check [post]
func (h *Handler) CheckStudentTest(c *gin.Context) {
	req := model.CheckStudentTestReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Error is get data: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{Message: "Noto'g'ri ma'lumot kiritildi"})
		return
	}
	resp, err := h.Service.CheckStudentTest(c, &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Error is service function CheckStudentTest: %v", err))
		c.JSON(http.StatusInternalServerError, model.Error{
			Message: "Serverda xatolik",
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary     Student shablonini yuklab olish
// @Description Student ID va sana bo‘yicha shablonni yuklab olish
// @Tags        Templates
// @Accept      json
// @Produce     octet-stream
// @Param       student_id query string true "Talaba ID"
// @Param       day        query string true "Kun (YYYY-MM-DD format)"
// @Success     200 {file} file "Shablon fayli"
// @Failure     400 {object} model.Error "Xato so‘rov"
// @Failure     500 {object} model.Error "Server xatosi"
// @Router      /templates/get [get]
func (h *Handler) GetStudentTemplate(c *gin.Context) {
	req := model.GetTemplatesReq{
		StudentId: c.Query("student_id"),
		Day:       c.Query("day"),
	}

	if req.StudentId == "" || req.Day == "" {
		c.JSON(http.StatusBadRequest, model.Error{Message: "student_id va day parametrlari kerak"})
		return
	}

	file, err := h.Service.GetStudentTemplates(c, &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Xatolik GetStudentTemplates: %v", err))
		c.JSON(http.StatusInternalServerError, model.Error{Message: "Serverda xatolik"})
		return
	}

	if file == nil || len(*file) == 0 {
		c.JSON(http.StatusNotFound, model.Error{Message: "Fayl topilmadi"})
		return
	}

	// Fayl nomini yaratish
	fileName := fmt.Sprintf("template_%s_%s.pdf", req.StudentId, req.Day)

	// Faylni yuklab berish
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Length", fmt.Sprintf("%d", len(*file)))
	c.Writer.Write(*file)
}
