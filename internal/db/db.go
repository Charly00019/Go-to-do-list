package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/Charly00019/Go-to-do-list/internal/models" // Import models package
)

var Database *gorm.DB

func InitDB() {
	var err error
	Database, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migrate tables
	Database.AutoMigrate(&models.Todo{}) // Use models.Todo
}
