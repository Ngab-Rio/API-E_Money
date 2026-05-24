package middlewares

import (
	"e-money/internal/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtManager utils.JWTManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHead := c.GetHeader("Authorization")
		if authHead == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Authorization header is required",
			})
			return
		}

		parts := strings.Split(authHead, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Invalid Authorization header",
			})
			return
		}

		token := parts[1]
		claims, err := jwtManager.Validate(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Invalid token",
			})
			return
		}

		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
