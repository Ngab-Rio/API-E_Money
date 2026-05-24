package routers

import (
	"e-money/internal/handlers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup, authHandler *handlers.AuthHandler) {
	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)
}
