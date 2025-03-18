package handlers

import (
	"net/http"

	"github.com/Charly00019/Go-to-do-list/internal/db"
	"github.com/Charly00019/Go-to-do-list/internal/models"
	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	var todos []models.Todo
	db.Database.Find(&todos)
	c.JSON(http.StatusOK, todos)
}

func CreateTodo(c *gin.Context) {
	var newTodo models.Todo
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Database.Create(&newTodo)
	c.JSON(http.StatusCreated, newTodo)
}
