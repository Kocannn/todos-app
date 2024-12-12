package domain

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	Id          int       `json:"id" gorm:"primary_key;autoIncrement"`
	Title       string    `json:"title" gorm:"type:varchar(100);not null"`
	Description string    `json:"description" gorm:"type:varchar(100);not null"`
	CreatedAt   time.Time `json:"-" gorm:"type:datetime;not null;autoCreateTime"`
	UpdatedAt   time.Time `json:"-" gorm:"type:datetime;not null;autoUpdateTime"`
	UserId      int       `json:"user_id"`
	User        User      `json:"-" gorm:"foreignKey:UserId"`
}

type TodoHandler interface {
	CreateTodo(c *gin.Context)
	GetTodos(c *gin.Context)
	DeleteTodo(c *gin.Context)
	UpdateTodo(c *gin.Context)
}

type TodoRepository interface {
	CreateTodo(todo *Todo) (*Todo, error)
	GetTodos(userId int) ([]Todo, error)
	DeleteTodo(todoId int) error
	UpdateTodo(todoId int, todo *Todo) (*Todo, error)
}

type TodoService interface {
	CreateTodo(todo *Todo) (*Todo, error)
	GetTodos(userId int) ([]Todo, error)
	DeleteTodo(todoId int) error
	UpdateTodo(todoId int, todo *Todo) (*Todo, error)
}
