package repository

import (
	"github.com/kocannn/todos-app/domain"
	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

// CreateTodo implements domain.TodoRepository.
func (t *TodoRepository) CreateTodo(todo *domain.Todo) (*domain.Todo, error) {
	if err := t.db.Create(&todo).Error; err != nil {
		return todo, err
	}
	return todo, nil
}

// DeleteTodo implements domain.TodoRepository.
func (t *TodoRepository) DeleteTodo(todoId int) error {
	if err := t.db.Delete(&domain.Todo{}, todoId).Error; err != nil {
		return err
	}
	return nil
}

// GetTodos implements domain.TodoRepository.
func (t *TodoRepository) GetTodos(userId int) ([]domain.Todo, error) {
	var todos []domain.Todo
	if err := t.db.Where("user_id = ?", userId).Find(&todos).Error; err != nil {
		return todos, err
	}
	return todos, nil
}

// UpdateTodo implements domain.TodoRepository.
func (t *TodoRepository) UpdateTodo(todoId int, todo *domain.Todo) (*domain.Todo, error) {
	if err := t.db.Model(&domain.Todo{}).Where("id = ?", todoId).Updates(todo).Error; err != nil {
		return todo, err
	}
	return todo, nil
}

func NewTodoRepository(db *gorm.DB) domain.TodoRepository {
	return &TodoRepository{
		db: db,
	}
}
