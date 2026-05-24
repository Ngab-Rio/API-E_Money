package handlers

import (
	"e-money/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	profileService services.ProfileService
}

func NewProfileHandler(profileService services.ProfileService) *ProfileHandler {
	return &ProfileHandler{profileService: profileService}
}

func (h *ProfileHandler) GetProfileByUserID(c *gin.Context) {
	userID := c.GetInt("user_id")

	profile, err := h.profileService.GetProfileByUserID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Get Profile Success",
		"data":    profile,
	})
}
