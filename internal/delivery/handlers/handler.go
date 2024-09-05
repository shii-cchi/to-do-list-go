package handlers

import (
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"to-do-list-go/internal/delivery/middleware"
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
	r.With(middleware.CheckTodoInput(h.TodoHandler.validator)).Post("/tasks", h.TodoHandler.createTodoHandler)
	r.Get("/tasks", h.TodoHandler.getTodosHandler)
	r.With(middleware.GetTodoID).Get("/tasks/{id}", h.TodoHandler.getTodoHandler)
	r.With(middleware.CheckTodoInput(h.TodoHandler.validator), middleware.GetTodoID).Put("/tasks/{id}", h.TodoHandler.updateTodoHandler)
	r.With(middleware.GetTodoID).Delete("/tasks/{id}", h.TodoHandler.deleteTodoHandler)
}
