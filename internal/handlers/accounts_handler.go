package handlers

import (
	"e-money/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	accountService services.AccountService
}

func NewAccountHandler(accountService services.AccountService) *AccountHandler {
	return &AccountHandler{accountService: accountService}
}

func (h *AccountHandler) GetAccountByUserID(c *gin.Context) {

	userIDAny, ok := c.Get("user_id")

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Unauthorized",
		})
		return
	}

	userID := userIDAny.(int)

	response, err := h.accountService.GetAccountByUserID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Get Account Failed",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Get Account Success",
		"data":    response,
	})
}

func (h *AccountHandler) GetSummary(c *gin.Context) {

	userIDAny, ok := c.Get("user_id")

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Unauthorized",
		})
		return
	}

	userID := userIDAny.(int)

	response, err := h.accountService.GetSummary(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Get Summary Failed",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Get Summary Success",
		"data": gin.H{
			"total_balance": response,
		},
	})
}
