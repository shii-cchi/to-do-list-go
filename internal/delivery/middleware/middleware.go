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

// CheckTodoInput validates the request body against the TodoInputDto schema and adds it to the request context.
func CheckTodoInput(validate *validator.Validate) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			todoInput := dto.TodoInputDto{}
			if err := json.NewDecoder(r.Body).Decode(&todoInput); err != nil {
				log.Printf(delivery.ErrInvalidInput+": %s\n", err)
				delivery.RespondWithError(w, http.StatusBadRequest, delivery.ErrInvalidInput)
				return
			}

			if err := validate.Struct(&todoInput); err != nil {
				log.Printf(delivery.ErrInvalidInput+": %s\n", err)
				delivery.RespondWithError(w, http.StatusBadRequest, delivery.ErrInvalidInput)
				return
			}

			ctx := context.WithValue(r.Context(), delivery.TodoInputKey, todoInput)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// GetTodoID extracts the todos ID from the request URL and adds it to the request context.
func GetTodoID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		todoIDStr := chi.URLParam(r, "id")
		todoID, err := strconv.Atoi(todoIDStr)
		if err != nil || todoID <= 0 {
			if err != nil {
				log.Printf(delivery.ErrInvalidTodoID+": %s\n", err)
			} else {
				log.Printf(delivery.ErrInvalidTodoID+": %d\n", todoID)
			}

			delivery.RespondWithError(w, http.StatusBadRequest, delivery.ErrInvalidTodoID)
			return
		}

		ctx := context.WithValue(r.Context(), delivery.TodoIDKey, todoID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
