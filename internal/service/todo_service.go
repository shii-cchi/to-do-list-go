package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
	"to-do-list-go/internal/database"
	"to-do-list-go/internal/delivery/dto"
)

const (
	ErrTodoNotFound = "todo with this id not found"
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
	newTodo, err := t.repo.CreateTodo(context.Background(), database.CreateTodoParams{
		Title:       todoInput.Title,
		Description: todoInput.Description,
		DueDate:     todoInput.DueDate,
	})
	if err != nil {
		return dto.TodoResponseDto{}, err
	}

	return dto.TodoResponseDto{
		ID:          newTodo.ID,
		Title:       newTodo.Title,
		Description: newTodo.Description,
		DueDate:     newTodo.DueDate,
		CreatedAt:   newTodo.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   newTodo.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (t TodoService) GetTodos() ([]dto.TodoResponseDto, error) {
	todos, err := t.repo.GetTodos(context.Background())
	if err != nil {
		return nil, err
	}

	return t.makeTodosResponseDto(todos), nil
}

func (t TodoService) GetTodo(todoID int) (dto.TodoResponseDto, error) {
	todo, err := t.repo.GetTodo(context.Background(), int32(todoID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return dto.TodoResponseDto{}, fmt.Errorf(ErrTodoNotFound+": %s\n", err)
		}

		return dto.TodoResponseDto{}, err
	}

	return dto.TodoResponseDto{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		DueDate:     todo.DueDate,
		CreatedAt:   todo.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   todo.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (t TodoService) UpdateTodo(todoID int, todoInput dto.TodoInputDto) (dto.TodoResponseDto, error) {
	updatedTodo, err := t.repo.UpdateTodo(context.Background(), database.UpdateTodoParams{
		ID:          int32(todoID),
		Title:       todoInput.Title,
		Description: todoInput.Description,
		DueDate:     todoInput.DueDate,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return dto.TodoResponseDto{}, fmt.Errorf(ErrTodoNotFound+": %s\n", err)
		}

		return dto.TodoResponseDto{}, err
	}

	return dto.TodoResponseDto{
		ID:          updatedTodo.ID,
		Title:       updatedTodo.Title,
		Description: updatedTodo.Description,
		DueDate:     updatedTodo.DueDate,
		CreatedAt:   updatedTodo.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   updatedTodo.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (t TodoService) DeleteTodo(todoID int) error {
	_, err := t.repo.DeleteTodo(context.Background(), int32(todoID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf(ErrTodoNotFound+": %s\n", err)
		}

		return err
	}

	return nil
}

func (t TodoService) makeTodosResponseDto(todos []database.Todo) []dto.TodoResponseDto {
	todosResponseDto := make([]dto.TodoResponseDto, len(todos))
	for i, todo := range todos {
		todosResponseDto[i] = dto.TodoResponseDto{
			ID:          todo.ID,
			Title:       todo.Title,
			Description: todo.Description,
			DueDate:     todo.DueDate,
			CreatedAt:   todo.CreatedAt.Format(time.RFC3339),
			UpdatedAt:   todo.UpdatedAt.Format(time.RFC3339),
		}
	}
	return todosResponseDto
}
