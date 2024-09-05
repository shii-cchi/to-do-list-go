package service

import (
	"to-do-list-go/internal/database"
	"to-do-list-go/internal/delivery/dto"
)

type Todo interface {
	CreateTodo(todoInput dto.TodoInputDto) (dto.TodoResponseDto, error)
	GetTodos() ([]dto.TodoResponseDto, error)
	GetTodo(todoID int) (dto.TodoResponseDto, error)
	UpdateTodo(todoID int, todoInput dto.TodoInputDto) (dto.TodoResponseDto, error)
	DeleteTodo(todoID int) error
}

type Service struct {
	Todo Todo
}

func NewService(repo database.Repository) *Service {
	todoService := newTodoService(repo)

	return &Service{
		Todo: todoService,
	}
}
