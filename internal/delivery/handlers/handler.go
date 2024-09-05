package handlers

import (
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"to-do-list-go/internal/delivery/middleware"
	"to-do-list-go/internal/service"
)

// Handler manages the endpoints, including the TodoHandler for handling todos-related requests.
type Handler struct {
	TodoHandler *TodoHandler
}

// NewHandler creates a new Handler.
func NewHandler(service *service.Service, validator *validator.Validate) *Handler {
	todoHandler := newTodoHandler(service.Todos, validator)

	return &Handler{
		TodoHandler: todoHandler,
	}
}

// RegisterRoutes manages route registration for todos endpoints with associated middlewares.
func (h Handler) RegisterRoutes(r *chi.Mux) {
	r.With(middleware.CheckTodoInput(h.TodoHandler.validator)).Post("/tasks", h.TodoHandler.createTodoHandler)
	r.Get("/tasks", h.TodoHandler.getTodosHandler)
	r.With(middleware.GetTodoID).Get("/tasks/{id}", h.TodoHandler.getTodoHandler)
	r.With(middleware.CheckTodoInput(h.TodoHandler.validator), middleware.GetTodoID).Put("/tasks/{id}", h.TodoHandler.updateTodoHandler)
	r.With(middleware.GetTodoID).Delete("/tasks/{id}", h.TodoHandler.deleteTodoHandler)
}
