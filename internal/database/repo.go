package database

import "context"

type Repository interface {
	CreateTodo(ctx context.Context, arg CreateTodoParams) (Todo, error)
	GetTodos(ctx context.Context) ([]Todo, error)
	GetTodo(ctx context.Context, id int32) (Todo, error)
	UpdateTodo(ctx context.Context, arg UpdateTodoParams) (Todo, error)
	DeleteTodo(ctx context.Context, id int32) (Todo, error)
}
