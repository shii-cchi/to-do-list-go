package domain

const (
	ErrInvalidInput = "invalid todo input body(fields title, description and due_date are required and can't be empty, due_date field must be a string in RFC3339 format)"
	ErrCreatingTodo = "error creating todo"

	ErrGettingTodos  = "error getting todos"
	ErrInvalidTodoID = "invalid todo id"
	ErrTodoNotFound  = "todo with this id not found"
	ErrGettingTodo   = "error getting todo"

	ErrUpdatingTodo = "error updating todo"

	ErrDeletingTodo = "error deleting todo"
)
