package handlers

import (
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"strings"
	"to-do-list-go/internal/delivery"
	"to-do-list-go/internal/delivery/dto"
	"to-do-list-go/internal/service"
)

// TodoHandler manages todos-related operations.
type TodoHandler struct {
	todoService service.Todos
	validator   *validator.Validate
}

func newTodoHandler(todoService service.Todos, validator *validator.Validate) *TodoHandler {
	return &TodoHandler{
		todoService: todoService,
		validator:   validator,
	}
}

func (h TodoHandler) createTodoHandler(w http.ResponseWriter, r *http.Request) {
	todoInput := r.Context().Value(delivery.TodoInputKey).(dto.TodoInputDto)

	todo, err := h.todoService.CreateTodo(todoInput)
	if err != nil {
		log.Printf(delivery.ErrCreatingTodo+": %s\n", err)
		delivery.RespondWithError(w, http.StatusInternalServerError, delivery.ErrCreatingTodo)
		return
	}

	delivery.RespondWithJSON(w, http.StatusCreated, todo)
}

func (h TodoHandler) getTodosHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := h.todoService.GetTodos()
	if err != nil {
		log.Printf(delivery.ErrGettingTodos+": %s\n", err)
		delivery.RespondWithError(w, http.StatusInternalServerError, delivery.ErrGettingTodos)
		return
	}

	delivery.RespondWithJSON(w, http.StatusOK, todos)
}

func (h TodoHandler) getTodoHandler(w http.ResponseWriter, r *http.Request) {
	todoID := r.Context().Value(delivery.TodoIDKey).(int)

	todo, err := h.todoService.GetTodo(todoID)
	if err != nil {
		if strings.HasPrefix(err.Error(), delivery.ErrTodoNotFound) {
			log.Println(err)
			delivery.RespondWithError(w, http.StatusNotFound, delivery.ErrTodoNotFound)
			return
		}

		log.Printf(delivery.ErrGettingTodo+": %s\n", err)
		delivery.RespondWithError(w, http.StatusInternalServerError, delivery.ErrGettingTodo)
		return
	}

	delivery.RespondWithJSON(w, http.StatusOK, todo)
}

func (h TodoHandler) updateTodoHandler(w http.ResponseWriter, r *http.Request) {
	todoID := r.Context().Value(delivery.TodoIDKey).(int)
	todoInput := r.Context().Value(delivery.TodoInputKey).(dto.TodoInputDto)

	updatedTodo, err := h.todoService.UpdateTodo(todoID, todoInput)
	if err != nil {
		if strings.HasPrefix(err.Error(), delivery.ErrTodoNotFound) {
			log.Println(err)
			delivery.RespondWithError(w, http.StatusNotFound, delivery.ErrTodoNotFound)
			return
		}

		log.Printf(delivery.ErrUpdatingTodo+": %s\n", err)
		delivery.RespondWithError(w, http.StatusInternalServerError, delivery.ErrUpdatingTodo)
		return
	}

	delivery.RespondWithJSON(w, http.StatusOK, updatedTodo)
}

func (h TodoHandler) deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	todoID := r.Context().Value(delivery.TodoIDKey).(int)

	if err := h.todoService.DeleteTodo(todoID); err != nil {
		if strings.HasPrefix(err.Error(), delivery.ErrTodoNotFound) {
			log.Println(err)
			delivery.RespondWithError(w, http.StatusNotFound, delivery.ErrTodoNotFound)
			return
		}

		log.Printf(delivery.ErrDeletingTodo+": %s\n", err)
		delivery.RespondWithError(w, http.StatusInternalServerError, delivery.ErrDeletingTodo)
		return
	}

	delivery.RespondWithJSON(w, http.StatusNoContent, nil)
}
