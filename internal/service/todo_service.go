package service

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"to-do-list-go/internal/database"
	"to-do-list-go/internal/delivery/dto"
	"to-do-list-go/internal/domain"
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
	timeRFC3339 := time.Now().Format(time.RFC3339)

	newTodo, err := t.repo.CreateTodo(context.Background(), database.CreateTodoParams{
		Title:       todoInput.Title,
		Description: todoInput.Description,
		DueDate:     todoInput.DueDate,
		CreatedAt:   timeRFC3339,
		UpdatedAt:   timeRFC3339,
	})
	if err != nil {
		return dto.TodoResponseDto{}, err
	}

	return dto.TodoResponseDto{
		ID:          newTodo.ID,
		Title:       newTodo.Title,
		Description: newTodo.Description,
		DueDate:     newTodo.DueDate,
		CreatedAt:   newTodo.CreatedAt,
		UpdatedAt:   newTodo.UpdatedAt,
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
		if err == sql.ErrNoRows {
			return dto.TodoResponseDto{}, fmt.Errorf(domain.ErrTodoNotFound+": %s\n", err)
		}

		return dto.TodoResponseDto{}, err
	}

	return dto.TodoResponseDto{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		DueDate:     todo.DueDate,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
	}, nil
}

func (t TodoService) UpdateTodo(todoID int, todoInput dto.TodoInputDto) (dto.TodoResponseDto, error) {
	updatedTodo, err := t.repo.UpdateTodo(context.Background(), database.UpdateTodoParams{
		ID:          int32(todoID),
		Title:       todoInput.Title,
		Description: todoInput.Description,
		DueDate:     todoInput.DueDate,
		UpdatedAt:   time.Now().Format(time.RFC3339),
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return dto.TodoResponseDto{}, fmt.Errorf(domain.ErrTodoNotFound+": %s\n", err)
		}

		return dto.TodoResponseDto{}, err
	}

	return dto.TodoResponseDto{
		ID:          updatedTodo.ID,
		Title:       updatedTodo.Title,
		Description: updatedTodo.Description,
		DueDate:     updatedTodo.DueDate,
		CreatedAt:   updatedTodo.CreatedAt,
		UpdatedAt:   updatedTodo.UpdatedAt,
	}, nil
}

func (t TodoService) DeleteTodo(todoID int) error {
	_, err := t.repo.DeleteTodo(context.Background(), int32(todoID))
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf(domain.ErrTodoNotFound+": %s\n", err)
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
			CreatedAt:   todo.CreatedAt,
			UpdatedAt:   todo.UpdatedAt,
		}
	}
	return todosResponseDto
}
