package service

import "to-do-list-go/internal/database"

type Todo interface {
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
