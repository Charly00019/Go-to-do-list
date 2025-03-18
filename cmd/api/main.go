package main

import (
	"log"

	"github.com/Charly00019/Go-to-do-list/internal/db"
	"github.com/Charly00019/Go-to-do-list/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	db.InitDB()

	// Create Gin router
	r := gin.Default()
	r.LoadHTMLGlob("templates/*") // Load HTML templates

	// Setup Routes
	router.SetupRoutes(r)

	// Start server
	log.Println("Server running on http://localhost:8080")
	r.Run(":8080")
}
