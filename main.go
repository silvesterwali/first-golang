package main

import (
	"log"

	"myproject/database"
	"myproject/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	if err := database.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run("localhost:8080")
}
