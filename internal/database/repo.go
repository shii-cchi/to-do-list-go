package database

import "context"

type Repository interface {
	CreateTodo(ctx context.Context, arg CreateTodoParams) (Todo, error)
}
