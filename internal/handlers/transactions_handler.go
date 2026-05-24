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

type TransactionHandler struct {
	transactionService services.TransactionService
	accountService     services.AccountService
}

func NewTransactionHandler(transactionService services.TransactionService, accountService services.AccountService) *TransactionHandler {
	return &TransactionHandler{transactionService: transactionService, accountService: accountService}
}

func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	var req dto.CreateTransactionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid Request",
			"error":   err.Error(),
		})
		return
	}

	userID := c.GetInt("user_id")

	if err := h.transactionService.CreateTransaction(c.Request.Context(), userID, req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Transaction Created Successfully",
	})

}

func (h *TransactionHandler) GetTransactionByID(c *gin.Context) {
	userID := c.GetInt("user_id")
	idParam := c.Param("id")
	transactionID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid transaction ID",
			"error":   err.Error(),
		})
		return
	}

	transaction, err := h.transactionService.GetTransactionByID(c.Request.Context(), userID, transactionID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "Transaction not found",
				"error":   err.Error(),
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Internal Server Error",
				"error":   err.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    transaction,
	})
}

func (h *TransactionHandler) GetAllTransactionsByUserID(c *gin.Context) {
	userID := c.GetInt("user_id")

	filter := models.FilterTransaction{
		Type:        c.Query("type"),
		Category:    c.Query("category"),
		AccountType: c.Query("account_type"),
	}

	transactions, err := h.transactionService.GetAllTransactionsByUserID(c.Request.Context(), userID, filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Get All Transactions Success",
		"data":    transactions,
	})
}
