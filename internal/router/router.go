package router

import (
	"github.com/Charly00019/Go-to-do-list/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/todos", handlers.GetTodos)
	r.POST("/todos", handlers.CreateTodo)
}
