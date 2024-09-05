package delivery

type contextKey string

// Defines context keys and error messages for todos operations.
const (
	TodoInputKey contextKey = "todoInput"
	TodoIDKey    contextKey = "todoID"

	ErrInvalidInput  = "invalid todo input body(fields title, description and due_date are required and can't be empty, due_date field must be a string in RFC3339 format)"
	ErrInvalidTodoID = "invalid todo id"

	ErrCreatingTodo = "error creating todo"
	ErrGettingTodos = "error getting todos"
	ErrTodoNotFound = "todo with this id not found"
	ErrGettingTodo  = "error getting todo"
	ErrUpdatingTodo = "error updating todo"
	ErrDeletingTodo = "error deleting todo"
)
