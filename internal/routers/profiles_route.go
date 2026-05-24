package routers

import (
	"e-money/internal/handlers"

	"github.com/gin-gonic/gin"
)

func ProfileRoutes(r *gin.RouterGroup, profileHandler *handlers.ProfileHandler) {
	r.GET("", profileHandler.GetProfileByUserID)
}
