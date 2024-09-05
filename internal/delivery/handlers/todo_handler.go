package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"strconv"
	"strings"
	"to-do-list-go/internal/delivery"
	"to-do-list-go/internal/delivery/dto"
	"to-do-list-go/internal/domain"
	"to-do-list-go/internal/service"
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
	todoInput := dto.TodoInputDto{}
	if err := json.NewDecoder(r.Body).Decode(&todoInput); err != nil {
		log.Printf(domain.ErrInvalidInput+": %s\n", err)
		delivery.RespondWithError(w, http.StatusBadRequest, domain.ErrInvalidInput)
		return
	}

	if err := h.validator.Struct(&todoInput); err != nil {
		log.Printf(domain.ErrInvalidInput+": %s\n", err)
		delivery.RespondWithError(w, http.StatusBadRequest, domain.ErrInvalidInput)
		return
	}

	todo, err := h.todoService.CreateTodo(todoInput)
	if err != nil {
		log.Printf(domain.ErrCreatingTodo+": %s\n", err)
		delivery.RespondWithError(w, http.StatusInternalServerError, domain.ErrCreatingTodo)
		return
	}

	delivery.RespondWithJSON(w, http.StatusCreated, todo)
}

func (h TodoHandler) getTodosHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := h.todoService.GetTodos()
	if err != nil {
		log.Printf(domain.ErrGettingTodos+": %s\n", err)
		delivery.RespondWithError(w, http.StatusInternalServerError, domain.ErrGettingTodos)
		return
	}

	delivery.RespondWithJSON(w, http.StatusOK, todos)
}

func (h TodoHandler) getTodoHandler(w http.ResponseWriter, r *http.Request) {
	todoIDStr := chi.URLParam(r, "id")
	todoID, err := strconv.Atoi(todoIDStr)
	if err != nil || todoID <= 0 {
		log.Printf(domain.ErrInvalidTodoID+":%s\n", err)
		delivery.RespondWithError(w, http.StatusBadRequest, domain.ErrInvalidTodoID)
		return
	}

	todo, err := h.todoService.GetTodo(todoID)
	if err != nil {
		if strings.HasPrefix(err.Error(), domain.ErrTodoNotFound) {
			log.Println(err)
			delivery.RespondWithError(w, http.StatusNotFound, domain.ErrTodoNotFound)
			return
		}

		log.Printf(domain.ErrGettingTodo+": %s\n", err)
		delivery.RespondWithError(w, http.StatusInternalServerError, domain.ErrGettingTodo)
		return
	}

	delivery.RespondWithJSON(w, http.StatusOK, todo)
}

func (h TodoHandler) updateTodoHandler(w http.ResponseWriter, r *http.Request) {
	todoIDStr := chi.URLParam(r, "id")
	todoID, err := strconv.Atoi(todoIDStr)
	if err != nil || todoID <= 0 {
		log.Printf(domain.ErrInvalidTodoID+":%s\n", err)
		delivery.RespondWithError(w, http.StatusBadRequest, domain.ErrInvalidTodoID)
		return
	}

	todoInput := dto.TodoInputDto{}
	if err := json.NewDecoder(r.Body).Decode(&todoInput); err != nil {
		log.Printf(domain.ErrInvalidInput+": %s\n", err)
		delivery.RespondWithError(w, http.StatusBadRequest, domain.ErrInvalidInput)
		return
	}

	if err := h.validator.Struct(&todoInput); err != nil {
		log.Printf(domain.ErrInvalidInput+": %s\n", err)
		delivery.RespondWithError(w, http.StatusBadRequest, domain.ErrInvalidInput)
		return
	}

	updatedTodo, err := h.todoService.UpdateTodo(todoID, todoInput)
	if err != nil {
		if strings.HasPrefix(err.Error(), domain.ErrTodoNotFound) {
			log.Println(err)
			delivery.RespondWithError(w, http.StatusNotFound, domain.ErrTodoNotFound)
			return
		}

		log.Printf(domain.ErrUpdatingTodo+": %s\n", err)
		delivery.RespondWithError(w, http.StatusInternalServerError, domain.ErrUpdatingTodo)
		return
	}

	delivery.RespondWithJSON(w, http.StatusOK, updatedTodo)
}

func (h TodoHandler) deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	todoIDStr := chi.URLParam(r, "id")
	todoID, err := strconv.Atoi(todoIDStr)
	if err != nil || todoID <= 0 {
		log.Printf(domain.ErrInvalidTodoID+":%s\n", err)
		delivery.RespondWithError(w, http.StatusBadRequest, domain.ErrInvalidTodoID)
		return
	}

	if err := h.todoService.DeleteTodo(todoID); err != nil {
		if strings.HasPrefix(err.Error(), domain.ErrTodoNotFound) {
			log.Println(err)
			delivery.RespondWithError(w, http.StatusNotFound, domain.ErrTodoNotFound)
			return
		}

		log.Printf(domain.ErrDeletingTodo+": %s\n", err)
		delivery.RespondWithError(w, http.StatusInternalServerError, domain.ErrDeletingTodo)
		return
	}

	delivery.RespondWithJSON(w, http.StatusNoContent, nil)
}
