package handlers

import (
	"e-money/internal/dto"
	"e-money/internal/models"
	"e-money/internal/services"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DebtHandler struct {
	debtService services.DebtService
}

func NewDebtHandler(debtService services.DebtService) *DebtHandler {
	return &DebtHandler{debtService: debtService}
}

func (h *DebtHandler) GetDebtByID(c *gin.Context) {
	userID := c.GetInt("user_id")
	idParam := c.Param("id")
	debtID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid debt ID",
			"error":   err.Error(),
		})
		return
	}

	debt, err := h.debtService.GetDebtByID(c.Request.Context(), userID, debtID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "Debt not found",
				"error":   err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to retrieve debt",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    debt,
	})
}

func (h *DebtHandler) GetDebts(c *gin.Context) {
	userID := c.GetInt("user_id")
	filter := models.FilterDebt{
		Type:   c.Query("type"),
		Status: c.Query("status"),
	}
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid query parameters",
			"error":   err.Error(),
		})
		return
	}

	debts, err := h.debtService.GetDebtByUserID(c.Request.Context(), userID, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to retrieve debts",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    debts,
	})
}

func (h *DebtHandler) CreateDebt(c *gin.Context) {
	var req dto.CreateDebtRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request payload",
			"error":   err.Error(),
		})
		return
	}
	userID := c.GetInt("user_id")

	if err := h.debtService.CreateDebt(c.Request.Context(), userID, req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Debt created successfully",
			"success": true,
		})
	}

}

func (h *DebtHandler) UpdateDebt(c *gin.Context) {
	var req dto.UpdateDebtRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request payload",
			"error":   err.Error(),
		})
		return
	}
	userID := c.GetInt("user_id")
	idParam := c.Param("id")
	debtID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid debt ID",
			"error":   err.Error(),
		})
		return
	}

	if err := h.debtService.UpdateDebt(c.Request.Context(), userID, debtID, req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Debt updated successfully",
		"success": true,
	})
}

func (h *DebtHandler) DeleteDebt(c *gin.Context) {
	userID := c.GetInt("user_id")
	idParam := c.Param("id")
	debtID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid debt ID format",
			"error":   err.Error(),
		})
		return
	}

	if err := h.debtService.DeleteDebt(c.Request.Context(), debtID, userID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "Debt not found",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to delete debt",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Debt deleted successfully",
	})
}
