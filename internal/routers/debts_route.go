package routers

import (
	"e-money/internal/handlers"

	"github.com/gin-gonic/gin"
)

func DebtRoutes(r *gin.RouterGroup, debtHandler *handlers.DebtHandler) {
	r.POST("", debtHandler.CreateDebt)
	r.PUT("/:id", debtHandler.UpdateDebt)
	r.GET("", debtHandler.GetDebts)
	r.GET("/:id", debtHandler.GetDebtByID)
	r.DELETE("/:id", debtHandler.DeleteDebt)
}
