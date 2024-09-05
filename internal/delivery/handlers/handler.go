package handlers

import (
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"to-do-list-go/internal/service"
)

type Handler struct {
	TodoHandler *TodoHandler
}

func NewHandler(service *service.Service, validator *validator.Validate) *Handler {
	todoHandler := newTodoHandler(service.Todo, validator)

	return &Handler{
		TodoHandler: todoHandler,
	}
}

func (h Handler) RegisterRoutes(r *chi.Mux) {
	r.Post("/tasks", h.TodoHandler.createTodoHandler)
	r.Get("/tasks", h.TodoHandler.getTodosHandler)
	r.Get("/tasks/{id}", h.TodoHandler.getTodoHandler)
	r.Put("/tasks/{id}", h.TodoHandler.updateTodoHandler)
	r.Delete("/tasks/{id}", h.TodoHandler.deleteTodoHandler)
}
