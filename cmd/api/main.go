package main

import (
	"architecture_portfolio_api/config"
	"architecture_portfolio_api/internal/router"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	log.Println("Environment variables loaded successfully")

	//Connect to PostgreSQL
	db, err := config.NewDatabasePool()
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	defer db.Close()

	log.Println("Database connected successfully")

	// Setup Router
	app := router.NewRouter(db)

	// Start Server
	log.Println("Server running on port 8080")

	if err := app.Run(":8080"); err != nil {
		log.Fatal("Failed to start server", err)
	}
}
