package handler

import (
	"edutest/pkg/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateSubject godoc
// @Summary Fan yaratish
// @Description Yangi fan qo‘shish
// @Tags Subjects
// @Accept json
// @Produce json
// @Param request body model.CreateSubjectReq true "Yangi fan ma'lumotlari"
// @Success 200 {object} model.Status "Muvaffaqiyatli qo‘shildi"
// @Failure 400 {object} model.Error "Noto'g'ri ma'lumot kiritildi"
// @Failure 500 {object} model.Error "Serverda xatolik yuz berdi"
// @Router /subjects/create [post]
func (h *Handler) CerateSubject(c *gin.Context) {
	req := model.CreateSubjectReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Error is get data: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{Message: "Noto'g'ri ma'lumot kiritildi"})
		return
	}
	err = h.Service.CreateSubject(c, &req)
	if err != nil{
		h.Log.Error(fmt.Sprintf("Error is service function CreateSubject: %v", err))
		c.JSON(http.StatusInternalServerError, model.Error{Message: "Serverda xatolik"})
		return
	}
	c.JSON(http.StatusOK, model.Status{Message: "Success!"})
}

// @Summary Fan o'chirish
// @Description Berilgan ID bo‘yicha fanni o‘chirish
// @Tags Subjects
// @Param id path string true "Fan ID'si"
// @Success 200 {object} model.Status "Muvaffaqiyatli o‘chirildi"
// @Failure 500 {object} model.Error "Serverda xatolik yuz berdi"
// @Router /subjects/delete/{subject_id} [delete]
func (h *Handler) DeleteSubject(c *gin.Context){
	err := h.Service.DeleteSubject(c, c.Param("subject_id"))
	if err != nil{
		h.Log.Error(fmt.Sprintf("Error is service function DeleteSubject: %v", err))
		c.JSON(http.StatusInternalServerError, model.Error{Message: "Serverda xatolik"})
		return
	}
	c.JSON(http.StatusOK, model.Status{Message: "Success!"})
}

// @Summary Fan(lar) ma'lumotlarini olish
// @Description Agar ID berilsa, shu fanni, aks holda barcha fanlarni qaytaradi
// @Tags Subjects
// @Param id query string false "Fan ID'si (Majburiy emas, agar berilmasa barcha fanlar qaytariladi)"
// @Success 200 {object} model.GetSubjectsResp "Fanlar ma'lumotlari"
// @Failure 500 {object} model.Error "Serverda xatolik yuz berdi"
// @Router /subjects/get [get]
func (h *Handler) GetSubjects(c *gin.Context){
	resp, err := h.Service.GetSubjects(c, c.Query("id"))
	if err != nil{
		h.Log.Error(fmt.Sprintf("Error is service function GetSubjects: %v", err))
		c.JSON(http.StatusInternalServerError, model.Error{Message: "Serverda xatolik"})
		return
	}
	c.JSON(http.StatusOK, resp)
}
