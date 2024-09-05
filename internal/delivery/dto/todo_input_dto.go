package dto

// TodoInputDto represents the input data required to create or update todos, with validation rules.
type TodoInputDto struct {
	Title       string `json:"title" validate:"required,min=1"`
	Description string `json:"description" validate:"required,min=1"`
	DueDate     string `json:"due_date" validate:"required,rfc3339"`
}
