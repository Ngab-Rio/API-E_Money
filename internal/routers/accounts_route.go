package routers

import (
	"e-money/internal/handlers"

	"github.com/gin-gonic/gin"
)

func AccountRoutes(r *gin.RouterGroup, accountHandler *handlers.AccountHandler) {
	r.GET("", accountHandler.GetAccountByUserID)
	r.GET("/summary", accountHandler.GetSummary)
}
