package service

import (
	"to-do-list-go/internal/database"
	"to-do-list-go/internal/delivery/dto"
)

type TodoService struct {
	repo database.Repository
}

func newTodoService(repo database.Repository) *TodoService {
	return &TodoService{
		repo: repo,
	}
}

func (t TodoService) CreateTodo(todoInput dto.TodoInputDto) (dto.TodoResponseDto, error) {

}

func (t TodoService) GetTodos() ([]dto.TodoResponseDto, error) {

}

func (t TodoService) GetTodo(todoID int) (dto.TodoResponseDto, error) {

}

func (t TodoService) UpdateTodo(todoID int, todoInput dto.TodoInputDto) (dto.TodoResponseDto, error) {

}

func (t TodoService) DeleteTodo(todoID int) error {

}
