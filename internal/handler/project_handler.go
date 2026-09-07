package handler

import (
	"architecture_portfolio_api/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProjectHandler struct {
	service *service.ProjectService
}

func NewProjectHandler(service *service.ProjectService) *ProjectHandler {
	return &ProjectHandler{service: service}
}

func (h *ProjectHandler) GetProjects(c *gin.Context) {
	locale := c.DefaultQuery("locale", "en")

	projects, err := h.service.GetProjects(
		c.Request.Context(),
		locale,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to get projects",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   projects,
	})
}

func (h *ProjectHandler) GetProjectBySlug(c *gin.Context) {
	slug := c.Param("slug")
	locale := c.DefaultQuery("locale", "en")

	project, err := h.service.GetProjectBySlug(
		c.Request.Context(),
		slug,
		locale,
	)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "Project not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   project,
	})
}
