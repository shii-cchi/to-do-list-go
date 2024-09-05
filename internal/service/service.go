package service

import (
	"to-do-list-go/internal/database"
	"to-do-list-go/internal/delivery/dto"
)

// Todos defines methods for managing todos operations.
type Todos interface {
	CreateTodo(todoInput dto.TodoInputDto) (dto.TodoResponseDto, error)
	GetTodos() ([]dto.TodoResponseDto, error)
	GetTodo(todoID int) (dto.TodoResponseDto, error)
	UpdateTodo(todoID int, todoInput dto.TodoInputDto) (dto.TodoResponseDto, error)
	DeleteTodo(todoID int) error
}

// Service manages todos-related operations through the Todos interface.
type Service struct {
	Todos Todos
}

// NewService creates a new Service instance.
func NewService(repo database.Repository) *Service {
	todoService := newTodoService(repo)

	return &Service{
		Todos: todoService,
	}
}
