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
	if err != nil{
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
func(h *Handler) CheckStudentTest(c *gin.Context){
	req := model.CheckStudentTestReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil{
		h.Log.Error(fmt.Sprintf("Error is get data: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{Message: "Noto'g'ri ma'lumot kiritildi"})
		return
	}
	resp, err := h.Service.CheckStudentTest(c, &req)
	if err != nil{
		h.Log.Error(fmt.Sprintf("Error is service function CheckStudentTest: %v", err))
		c.JSON(http.StatusInternalServerError, model.Error{Message: "Serverda xatolik"})
		return
	}
	c.JSON(http.StatusOK, resp)
}