package service

import "to-do-list-go/internal/database"

type TodoService struct {
	repo database.Repository
}

func newTodoService(repo database.Repository) *TodoService {
	return &TodoService{
		repo: repo,
	}
}
