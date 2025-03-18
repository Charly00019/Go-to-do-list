package main

import (
	"fmt"

	"github.com/Charly00019/Go-to-do-list/internal/db"
	"github.com/Charly00019/Go-to-do-list/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to DB
	db.InitDB()

	// Setup Routes
	router.SetupRoutes(r)

	fmt.Println("Server running on http://localhost:8080")
	r.Run(":8080")
}
