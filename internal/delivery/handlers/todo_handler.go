package handlers

import (
	"net/http"
	"to-do-list-go/internal/service"
)

type TodoHandler struct {
	todoService service.Todo
}

func newTodoHandler(todoService service.Todo) *TodoHandler {
	return &TodoHandler{
		todoService: todoService,
	}
}

func (h TodoHandler) createTodoHandler(w http.ResponseWriter, r *http.Request) {

}

func (h TodoHandler) getTodosHandler(w http.ResponseWriter, r *http.Request) {

}

func (h TodoHandler) getTodoHandler(w http.ResponseWriter, r *http.Request) {

}

func (h TodoHandler) updateTodoHandler(w http.ResponseWriter, r *http.Request) {

}

func (h TodoHandler) deleteTodoHandler(w http.ResponseWriter, r *http.Request) {

}
