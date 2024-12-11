package service

import "github.com/kocannn/todos-app/domain"

type TodoService struct {
	TodoRepo domain.TodoRepository
}

// CreateTodo implements domain.TodoService.
func (t *TodoService) CreateTodo(todo *domain.Todo) (*domain.Todo, error) {
	return t.TodoRepo.CreateTodo(todo)
}

// DeleteTodo implements domain.TodoService.
func (t *TodoService) DeleteTodo(todoId int) error {
	return t.TodoRepo.DeleteTodo(todoId)
}

// GetTodos implements domain.TodoService.
func (t *TodoService) GetTodos(userId int) ([]domain.Todo, error) {
	return t.TodoRepo.GetTodos(userId)
}

// UpdateTodo implements domain.TodoService.
func (t *TodoService) UpdateTodo(todoId int, todo *domain.Todo) (*domain.Todo, error) {
	return t.TodoRepo.UpdateTodo(todoId, todo)
}

func NewTodoService(todoRepo domain.TodoRepository) domain.TodoService {
	return &TodoService{
		TodoRepo: todoRepo,
	}
}
