package handler

import (
	"edutest/pkg/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary      Test uchun savollarni olish
// @Description  Berilgan foydalanuvchi va fan uchun test savollarini olish (tasodifiy tanlanadi)
// @Tags         test
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        user_id     query    string true  "Foydalanuvchi IDsi"
// @Param        subject_id  query    string true  "Fan IDsi"
// @Param        count       query    int    true  "Nechta savol kerakligi (soni)"
// @Success      200  {object}  model.GetTestResp  "Savollar ro'yxati"
// @Failure      500  {object}  model.Error       "Server xatosi haqida ma'lumot"
// @Router       /tests/get [get]
func (h *Handler) GetTests(c *gin.Context) {
	var req = model.GetTest{
		UserId:     c.Query("user_id"),
		Subject_Id: c.Query("subject_id"),
	}
	var count string = c.Query("count")
	req.Count, _ = strconv.Atoi(count)

	resp, err := h.Service.GetQuestionsForTest(&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Error is service function GetQuestionsForTest: %v", err))
		c.JSON(http.StatusInternalServerError, model.Error{
			Message: "Serverda xatolik",
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary      Test natijasini tekshirish
// @Description  Foydalanuvchidan kelgan javoblarni tekshiradi va natijani qaytaradi
// @Tags         test
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        data  body      model.CheckReq  true  "Foydalanuvchining test javoblari"
// @Success      200   {object}  model.TestResult      "Test natijasi"
// @Failure      400   {object}  model.Error          "Noto'g'ri ma'lumot"
// @Failure      500   {object}  model.Error          "Server xatosi"
// @Router       /tests/check [post]
func (h *Handler) CheckTest(c *gin.Context){
	var req = model.CheckReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil{
		h.Log.Error(fmt.Sprintf("Error is get data: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{Message: "Noto'g'ri ma'lumot kiritildi"})
		return
	}

	resp, err := h.Service.CheckTest(req)
	if err != nil{
		h.Log.Error(fmt.Sprintf("Error is service function CheckTest: %v", err))
		c.JSON(http.StatusInternalServerError, model.Error{
			Message: "Serverda xatolik",
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}
