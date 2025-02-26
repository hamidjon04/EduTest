package handler

import (
	"edutest/pkg/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary      Yangi savol yaratish
// @Description  Yangi savol qo'shish uchun API
// @Tags         Questions
// @Accept       json
// @Produce      json
// @Param        request body model.CreateQuestionReq true "Savol yaratish ma'lumotlari"
// @Success      200 {object} model.Status
// @Failure      400 {object} model.Error
// @Failure      500 {object} model.Error
// @Router       /questions/create [post]
func (h *Handler) CreateQuestion(c *gin.Context) {
	req := model.CreateQuestionReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Error is get data: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{Message: "Noto'g'ri ma'lumot kiritildi"})
		return
	}
	err = h.Service.CreateQuestion(c, &req)
	if err != nil{
		h.Log.Error(fmt.Sprintf("Error is service function CreateQuestion: %v", err))
		c.JSON(http.StatusInternalServerError, model.Error{Message: "Serverda xatolik"})
		return
	}
	c.JSON(http.StatusOK, model.Status{Message: "Success!"})
}

// @Summary      Savolni yangilash
// @Description  Berilgan ID bo‘yicha savolni yangilash
// @Tags         Questions
// @Accept       json
// @Produce      json
// @Param        id path string true "Savol IDsi"
// @Param        request body model.UpdateQuestionReq true "Savolni yangilash ma'lumotlari"
// @Success      200 {object} model.Status
// @Failure      400 {object} model.Error
// @Failure      500 {object} model.Error
// @Router       /questions/update/{id} [put]
func(h *Handler) UpdateQuestion(c *gin.Context){
	req := model.UpdateQuestionReq{
		Id: c.Param("id"),
	}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Error is get data: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{Message: "Noto'g'ri ma'lumot kiritildi"})
		return
	}
	err = h.Service.UpdateQuestion(c, &req)
	if err != nil{
		h.Log.Error(fmt.Sprintf("Error is service function UpdateQuestion: %v", err))
		c.JSON(http.StatusInternalServerError, model.Error{Message: "Serverda xatolik"})
		return
	}
	c.JSON(http.StatusOK, model.Status{Message: "Success!"})
}

// @Summary      Savolni o‘chirish
// @Description  Berilgan ID bo‘yicha savolni o‘chirish
// @Tags         Questions
// @Produce      json
// @Param        id path string true "Savol IDsi"
// @Success      200 {object} model.Status
// @Failure      500 {object} model.Error
// @Router       /questions/delete/{id} [delete]
func(h *Handler) DeleteQuestion(c *gin.Context){
	err := h.Service.DeleteQuestion(c, c.Param("id"))
	if err != nil{
		h.Log.Error(fmt.Sprintf("Error is service function DeleteQuestion: %v", err))
		c.JSON(http.StatusInternalServerError, model.Error{Message: "Serverda xatolik"})
		return
	}
	c.JSON(http.StatusOK, model.Status{Message: "Success!"})
}

// @Summary      Savollarni olish
// @Description  Berilgan ID yoki fan ID bo‘yicha savollarni olish
// @Tags         Questions
// @Produce      json
// @Param        id query string false "Savol IDsi"
// @Param        subject_id query string false "Fan IDsi"
// @Param        type query string false "Savol turi"
// @Success      200 {object} model.GetQuestionsResp
// @Failure      500 {object} model.Error
// @Router       /questions [get]
func (h *Handler) GetQuestions(c *gin.Context){
	req := model.GetQuestionsReq{
		Id: c.Query("id"),
		SubjectId: c.Query("subject_id"),
		Type: c.Query("type"),
	}
	resp, err := h.Service.GetQuestions(c, &req)
	if err != nil{
		h.Log.Error(fmt.Sprintf("Error is service function GetQuestions: %v", err))
		c.JSON(http.StatusInternalServerError, model.Error{Message: "Serverda xatolik"})
		return
	}
	c.JSON(http.StatusOK, resp)
}
