package handler

import (
	"edutest/pkg/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary      User registration
// @Description  Register a new user with necessary credentials
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body model.RegisterReq true "User registration data"
// @Success      200 {object} model.Tokens
// @Failure      400 {object} model.Error "Bad request: invalid input"
// @Failure      500 {object} model.Error "Internal server error"
// @Router       /register [post]
func(h *Handler) Register(c *gin.Context){
	var req = model.RegisterReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil{
		h.Log.Error(fmt.Sprintf("Error is get user's data: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return 
	}

	resp, err := h.Service.Register(req)
	if err != nil{
		h.Log.Error(fmt.Sprintf("Error at Register function: %v", err))
		c.JSON(http.StatusInternalServerError, model.Error{
			Error: err.Error(),
			Message: "Error at register function",
		})
		return 
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary      User login
// @Description  Authenticate user and return access/refresh tokens
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body model.LoginReq true "User credentials"
// @Success      200 {object} model.Tokens
// @Failure      400 {object} model.Error "Invalid request format"
// @Failure      404 {object} model.Error "User not found or invalid credentials"
// @Router       /login [post]
func(h *Handler) Login(c *gin.Context){
	var req = model.LoginReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil{
		h.Log.Error(fmt.Sprintf("Error is get user's data: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return 
	}

	resp, err := h.Service.Login(req)
	if err != nil{
		h.Log.Error(fmt.Sprintf("Error is sign in user: %v", err))
		c.JSON(http.StatusNotFound, model.Error{
			Error: err.Error(),
			Message: "Error is sign in user",
		})
		return
	}	
	c.JSON(http.StatusOK, resp)
}

// @Summary      Refresh access token
// @Description  Get a new access token using refresh token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user_id        query string true "User ID"
// @Param        refresh_token  query string true "Refresh Token"
// @Success      200 {object} model.Token
// @Failure      401 {object} model.Error "Invalid or expired refresh token"
// @Router       /refresh-token [get]
func(h *Handler) RefreshToken(c *gin.Context){
	var req = model.RefreshTokenReq{
		UserId: c.Query("user_id"),
		RefreshToken: c.Query("refresh_token"),
	}

	token, err := h.Service.RefreshToken(req)
	if err != nil{
		h.Log.Error(fmt.Sprintf("Error is refresh access token: %v", err))
		c.JSON(http.StatusUnauthorized, model.Error{
			Message: "Error is refresh access token",
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, token)
}