package handler

import (
	"architecture_portfolio_api/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProfileHandler struct {
	service *service.ProfileService
}

func NewProfileHandler(service *service.ProfileService) *ProfileHandler {
	return &ProfileHandler{service: service}
}

func (h *ProfileHandler) GetProfile(c *gin.Context) {
	locale := c.DefaultQuery("locale", "en")

	profile, err := h.service.GetProfile(
		c.Request.Context(),
		locale,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to get profile",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   profile,
	})
}
