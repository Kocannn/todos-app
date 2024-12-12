package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kocannn/todos-app/handler"
	"github.com/kocannn/todos-app/repository"
	"github.com/kocannn/todos-app/service"
	"gorm.io/gorm"
)

func TodoRoutes(group *gin.RouterGroup, db *gorm.DB) {
	todoRepo := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(todoRepo)
	todoHandler := handler.NewTodoHandler(todoService)

	group.POST("/todos", todoHandler.CreateTodo)
	group.GET("/todos", todoHandler.GetTodos)
	group.PUT("/todos/:id", todoHandler.UpdateTodo)
	group.DELETE("/todos/:id", todoHandler.DeleteTodo)
}
