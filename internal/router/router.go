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

	// Service
	profileService := service.NewProfileService(profileRepository)

	// Handler
	profileHandler := handler.NewProfileHandler(profileService)

	// Routes
	router.GET("/health", healthCheck)
	router.GET("/api/profile", profileHandler.GetProfile)

	return router
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Architecture Portfolio API is running",
	})
}
