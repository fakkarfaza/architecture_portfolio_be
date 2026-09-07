package router

import (
	"net/http"

	"architecture_portfolio_api/internal/handler"
	"architecture_portfolio_api/internal/repository"
	"architecture_portfolio_api/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewRouter(db *pgxpool.Pool) *gin.Engine {
	router := gin.Default()

	// Repository
	profileRepository := repository.NewProfileRepository(db)
	projectRepository := repository.NewProjectRepository(db)
	contactRepository := repository.NewContactRepository(db)

	// Service
	profileService := service.NewProfileService(profileRepository)
	projectService := service.NewProjectService(projectRepository)
	contactService := service.NewContactService(contactRepository)

	// Handler
	profileHandler := handler.NewProfileHandler(profileService)
	projectHandler := handler.NewProjectHandler(projectService)
	contactHandler := handler.NewContactHandler(contactService)

	// Routes
	router.GET("/health", healthCheck)
	router.GET("/api/profile", profileHandler.GetProfile)
	router.GET("api/projects", projectHandler.GetProjects)
	router.GET("/api/projects/:slug", projectHandler.GetProjectBySlug)
	router.POST("/api/contact", contactHandler.CreateContactMessage)

	return router
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Architecture Portfolio API is running",
	})
}
