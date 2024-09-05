package middleware

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"strconv"
	"to-do-list-go/internal/delivery"
	"to-do-list-go/internal/delivery/dto"
)

const (
	ErrInvalidInput  = "invalid todo input body(fields title, description and due_date are required and can't be empty, due_date field must be a string in RFC3339 format)"
	ErrInvalidTodoID = "invalid todo id"
)

func CheckTodoInput(validate *validator.Validate) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			todoInput := dto.TodoInputDto{}
			if err := json.NewDecoder(r.Body).Decode(&todoInput); err != nil {
				log.Printf(ErrInvalidInput+": %s\n", err)
				delivery.RespondWithError(w, http.StatusBadRequest, ErrInvalidInput)
				return
			}

			if err := validate.Struct(&todoInput); err != nil {
				log.Printf(ErrInvalidInput+": %s\n", err)
				delivery.RespondWithError(w, http.StatusBadRequest, ErrInvalidInput)
				return
			}

			ctx := context.WithValue(r.Context(), "todoInput", todoInput)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetTodoID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		todoIDStr := chi.URLParam(r, "id")
		todoID, err := strconv.Atoi(todoIDStr)
		if err != nil || todoID <= 0 {
			if err != nil {
				log.Printf(ErrInvalidTodoID+": %s\n", err)
			} else {
				log.Printf(ErrInvalidTodoID+": %d\n", todoID)
			}

			delivery.RespondWithError(w, http.StatusBadRequest, ErrInvalidTodoID)
			return
		}

		ctx := context.WithValue(r.Context(), "todoID", todoID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
