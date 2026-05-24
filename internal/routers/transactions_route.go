package routers

import (
	"e-money/internal/handlers"

	"github.com/gin-gonic/gin"
)

func TransactionRoutes(r *gin.RouterGroup, transactionHandler *handlers.TransactionHandler) {
	r.POST("", transactionHandler.CreateTransaction)
	r.GET("/:id", transactionHandler.GetTransactionByID)
	r.GET("", transactionHandler.GetAllTransactionsByUserID)
}
