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

const (
	ErrCreatingTodo = "error creating todo"
	ErrGettingTodos = "error getting todos"
	ErrTodoNotFound = "todo with this id not found"
	ErrGettingTodo  = "error getting todo"
	ErrUpdatingTodo = "error updating todo"
	ErrDeletingTodo = "error deleting todo"
)

type TodoHandler struct {
	todoService service.Todo
	validator   *validator.Validate
}

func newTodoHandler(todoService service.Todo, validator *validator.Validate) *TodoHandler {
	return &TodoHandler{
		todoService: todoService,
		validator:   validator,
	}
}

func (h TodoHandler) createTodoHandler(w http.ResponseWriter, r *http.Request) {
	todoInput := r.Context().Value("todoInput").(dto.TodoInputDto)

	todo, err := h.todoService.CreateTodo(todoInput)
	if err != nil {
		log.Printf(ErrCreatingTodo+": %s\n", err)
		delivery.RespondWithError(w, http.StatusInternalServerError, ErrCreatingTodo)
		return
	}

	delivery.RespondWithJSON(w, http.StatusCreated, todo)
}

func (h TodoHandler) getTodosHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := h.todoService.GetTodos()
	if err != nil {
		log.Printf(ErrGettingTodos+": %s\n", err)
		delivery.RespondWithError(w, http.StatusInternalServerError, ErrGettingTodos)
		return
	}

	delivery.RespondWithJSON(w, http.StatusOK, todos)
}

func (h TodoHandler) getTodoHandler(w http.ResponseWriter, r *http.Request) {
	todoID := r.Context().Value("todoID").(int)

	todo, err := h.todoService.GetTodo(todoID)
	if err != nil {
		if strings.HasPrefix(err.Error(), ErrTodoNotFound) {
			log.Println(err)
			delivery.RespondWithError(w, http.StatusNotFound, ErrTodoNotFound)
			return
		}

		log.Printf(ErrGettingTodo+": %s\n", err)
		delivery.RespondWithError(w, http.StatusInternalServerError, ErrGettingTodo)
		return
	}

	delivery.RespondWithJSON(w, http.StatusOK, todo)
}

func (h TodoHandler) updateTodoHandler(w http.ResponseWriter, r *http.Request) {
	todoID := r.Context().Value("todoID").(int)

	todoInput := r.Context().Value("todoInput").(dto.TodoInputDto)

	updatedTodo, err := h.todoService.UpdateTodo(todoID, todoInput)
	if err != nil {
		if strings.HasPrefix(err.Error(), ErrTodoNotFound) {
			log.Println(err)
			delivery.RespondWithError(w, http.StatusNotFound, ErrTodoNotFound)
			return
		}

		log.Printf(ErrUpdatingTodo+": %s\n", err)
		delivery.RespondWithError(w, http.StatusInternalServerError, ErrUpdatingTodo)
		return
	}

	delivery.RespondWithJSON(w, http.StatusOK, updatedTodo)
}

func (h TodoHandler) deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	todoID := r.Context().Value("todoID").(int)

	if err := h.todoService.DeleteTodo(todoID); err != nil {
		if strings.HasPrefix(err.Error(), ErrTodoNotFound) {
			log.Println(err)
			delivery.RespondWithError(w, http.StatusNotFound, ErrTodoNotFound)
			return
		}

		log.Printf(ErrDeletingTodo+": %s\n", err)
		delivery.RespondWithError(w, http.StatusInternalServerError, ErrDeletingTodo)
		return
	}

	delivery.RespondWithJSON(w, http.StatusNoContent, nil)
}
