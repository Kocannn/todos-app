package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kocannn/todos-app/domain"
)

type TodoHandler struct {
	TodoService domain.TodoService
}

// CreateTodo implements domain.TodoHandler.
func (t *TodoHandler) CreateTodo(c *gin.Context) {
	var todo *domain.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createTodo, err := t.TodoService.CreateTodo(todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "success create todo", "data": createTodo})
}

// DeleteTodo implements domain.TodoHandler.
func (t *TodoHandler) DeleteTodo(c *gin.Context) {
	var id *int

	if err := c.ShouldBindBodyWithJSON(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := t.TodoService.DeleteTodo(*id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})

}

// GetTodos implements domain.TodoHandler.
func (t *TodoHandler) GetTodos(c *gin.Context) {
	var todos []domain.Todo
	var userId *int

	if err := c.ShouldBindBodyWithJSON(&userId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todos, err := t.TodoService.GetTodos(*userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success get todo", "data": todos})

}

// UpdateTodo implements domain.TodoHandler.
func (t *TodoHandler) UpdateTodo(c *gin.Context) {
	var todo *domain.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTodo, err := t.TodoService.UpdateTodo(todo.Id, todo)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success update todo", "data": updatedTodo})
}

func NewTodoHandler(todoService domain.TodoService) domain.TodoHandler {
	return &TodoHandler{
		TodoService: todoService,
	}
}
