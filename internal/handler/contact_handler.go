package handler

import (
	"architecture_portfolio_api/internal/model"
	"architecture_portfolio_api/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type ContactHandler struct {
	service *service.ContactService
}

func NewContactHandler(service *service.ContactService) *ContactHandler {
	return &ContactHandler{service: service}
}

func (h *ContactHandler) CreateContactMessage(c *gin.Context) {
	var request model.ContactMessage

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request body",
		})
		return
	}

	err := h.service.CreateContactMessage(
		c.Request.Context(),
		&request,
	)
	if err != nil {
		log.Println("CreateContactMessage error:", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to send contact message",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Contact message sent successfully",
	})
}
