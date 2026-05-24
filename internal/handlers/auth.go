package handlers

import (
	"e-money/internal/dto"
	"e-money/internal/services"
	"e-money/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		err := utils.ValidateHandler(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"success":  false,
			"messsage": "Invalid Request Body",
			"error":    err,
		})
		return
	}

	response, err := h.authService.Register(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Register Failed",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Register Success",
		"data":    response,
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Request Body",
			"error":   err.Error(),
		})
		return
	}

	response, err := h.authService.Login(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Login Failed",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login Success",
		"data":    response,
	})
}
