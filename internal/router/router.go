package router

import (
	"net/http"

	"github.com/Charly00019/Go-to-do-list/internal/db"
	"github.com/Charly00019/Go-to-do-list/internal/models"
	"github.com/gin-gonic/gin"
)

// SetupRoutes initializes all routes
func SetupRoutes(r *gin.Engine) {
	r.LoadHTMLGlob("templates/*")         // Load HTML files
	r.Static("/static/js", "./static/js") // Serve static assets

	// Define routes only once
	r.GET("/", ShowIndex)
	r.GET("/todos", GetTodos)
	r.POST("/todos", AddTodo)
	r.PUT("/todos/:id", UpdateTodo)
	r.DELETE("/todos/:id", DeleteTodo)
}

// ShowIndex renders the HTML page
func ShowIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// GetTodos returns all todos
func GetTodos(c *gin.Context) {
	var todos []models.Todo
	db.Database.Find(&todos)
	c.JSON(http.StatusOK, todos)
}

// AddTodo creates a new todo
func AddTodo(c *gin.Context) {
	var input struct {
		Title string `json:"title"`
	}

	// Bind JSON data
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Validate input
	if input.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
		return
	}

	// Create a new Todo
	todo := models.Todo{
		Title:  input.Title,
		Status: "pending", // Default status
	}

	// Save to database
	db.Database.Create(&todo)
	c.JSON(http.StatusOK, todo)
}

// UpdateTodo marks a todo as completed
func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo
	if err := db.Database.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	todo.Status = "completed"
	db.Database.Save(&todo)
	c.JSON(http.StatusOK, todo)
}

// DeleteTodo removes a todo
func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	if err := db.Database.Delete(&models.Todo{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}
