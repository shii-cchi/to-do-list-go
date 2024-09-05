package handlers

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"to-do-list-go/internal/database"
	mock_repo "to-do-list-go/internal/database/mocks"
	"to-do-list-go/internal/delivery/dto"
	"to-do-list-go/internal/delivery/middleware"
	"to-do-list-go/internal/service"
	"to-do-list-go/internal/validator"
)

func TestCreateTodoHandler(t *testing.T) {
	type mockBehavior func(ctx context.Context, repo *mock_repo.MockRepository)

	tests := []struct {
		name           string
		input          io.Reader
		reqMethod      string
		reqTarget      string
		expectedStatus int
		expectedBody   interface{}
		mockBehavior   mockBehavior
	}{
		// CreateHandler
		{
			name: "CreateHandler Success",
			input: bytes.NewBuffer([]byte(`{
				"title": "test",
				"description": "test",
				"due_date": "2024-09-05T12:40:16+07:00"
			}`)),
			reqMethod:      http.MethodPost,
			reqTarget:      "/tasks",
			expectedStatus: http.StatusCreated,
			expectedBody: dto.TodoResponseDto{
				ID:          1,
				Title:       "test",
				Description: "test",
				DueDate:     "2024-09-05T12:40:16+07:00",
				CreatedAt:   "2024-09-05T12:24:16+07:00",
				UpdatedAt:   "2024-09-05T12:24:16+07:00",
			},
			mockBehavior: func(ctx context.Context, repo *mock_repo.MockRepository) {
				newTodo := database.Todo{
					ID:          1,
					Title:       "test",
					Description: "test",
					DueDate:     "2024-09-05T12:40:16+07:00",
					CreatedAt:   "2024-09-05T12:24:16+07:00",
					UpdatedAt:   "2024-09-05T12:24:16+07:00",
				}
				createdAndUpdatedAt := time.Now().Format(time.RFC3339)
				repo.EXPECT().CreateTodo(ctx, database.CreateTodoParams{
					Title:       "test",
					Description: "test",
					DueDate:     "2024-09-05T12:40:16+07:00",
					CreatedAt:   createdAndUpdatedAt,
					UpdatedAt:   createdAndUpdatedAt,
				}).Return(newTodo, nil).Times(1)
			},
		},
		{
			name: "CreateHandler Invalid Input 1",
			input: bytes.NewBuffer([]byte(`{
				"title": "test,
				"description": "test",
				"due_date": "2024-09-05T12:40:16+07:00"
			}`)),
			reqMethod:      http.MethodPost,
			reqTarget:      "/tasks",
			expectedStatus: http.StatusBadRequest,
			expectedBody: struct {
				Error string `json:"error"`
			}{
				Error: middleware.ErrInvalidInput,
			},
			mockBehavior: func(ctx context.Context, repo *mock_repo.MockRepository) {},
		},
		{
			name: "CreateHandler Invalid Input 2",
			input: bytes.NewBuffer([]byte(`{
				"description": "test",
				"due_date": "2024-09-05T12:40:16+07:00"
			}`)),
			reqMethod:      http.MethodPost,
			reqTarget:      "/tasks",
			expectedStatus: http.StatusBadRequest,
			expectedBody: struct {
				Error string `json:"error"`
			}{
				Error: middleware.ErrInvalidInput,
			},
			mockBehavior: func(ctx context.Context, repo *mock_repo.MockRepository) {},
		},
		{
			name: "CreateHandler Repo Error",
			input: bytes.NewBuffer([]byte(`{
				"title": "test",
				"description": "test",
				"due_date": "2024-09-05T12:40:16+07:00"
			}`)),
			reqMethod:      http.MethodPost,
			reqTarget:      "/tasks",
			expectedStatus: http.StatusInternalServerError,
			expectedBody: struct {
				Error string `json:"error"`
			}{
				Error: errCreatingTodo,
			},
			mockBehavior: func(ctx context.Context, repo *mock_repo.MockRepository) {
				createdUpdatedAt := time.Now().Format(time.RFC3339)
				repo.EXPECT().CreateTodo(ctx, database.CreateTodoParams{
					Title:       "test",
					Description: "test",
					DueDate:     "2024-09-05T12:40:16+07:00",
					CreatedAt:   createdUpdatedAt,
					UpdatedAt:   createdUpdatedAt,
				}).Return(database.Todo{}, errors.New("some db error")).Times(1)
			},
		},

		// GetTodosHandler
		{
			name:           "GetTodosHandler Success",
			input:          nil,
			reqMethod:      http.MethodGet,
			reqTarget:      "/tasks",
			expectedStatus: http.StatusOK,
			expectedBody: []dto.TodoResponseDto{
				{
					ID:          1,
					Title:       "test",
					Description: "test",
					DueDate:     "2024-09-05T12:40:16+07:00",
					CreatedAt:   "2024-09-05T12:24:16+07:00",
					UpdatedAt:   "2024-09-05T12:24:16+07:00",
				},
			},
			mockBehavior: func(ctx context.Context, repo *mock_repo.MockRepository) {
				todos := []database.Todo{
					{
						ID:          1,
						Title:       "test",
						Description: "test",
						DueDate:     "2024-09-05T12:40:16+07:00",
						CreatedAt:   "2024-09-05T12:24:16+07:00",
						UpdatedAt:   "2024-09-05T12:24:16+07:00",
					},
				}
				repo.EXPECT().GetTodos(ctx).Return(todos, nil).Times(1)
			},
		},
		{
			name:           "GetTodosHandler Repo Error",
			input:          nil,
			reqMethod:      http.MethodGet,
			reqTarget:      "/tasks",
			expectedStatus: http.StatusInternalServerError,
			expectedBody: struct {
				Error string `json:"error"`
			}{
				Error: errGettingTodos,
			},
			mockBehavior: func(ctx context.Context, repo *mock_repo.MockRepository) {
				repo.EXPECT().GetTodos(ctx).Return(nil, errors.New("some db error")).Times(1)
			},
		},

		// GetTodoHandler
		{
			name:           "GetTodoHandler Success",
			input:          nil,
			reqMethod:      http.MethodGet,
			reqTarget:      "/tasks/1",
			expectedStatus: http.StatusOK,
			expectedBody: dto.TodoResponseDto{
				ID:          1,
				Title:       "test",
				Description: "test",
				DueDate:     "2024-09-05T12:40:16+07:00",
				CreatedAt:   "2024-09-05T12:24:16+07:00",
				UpdatedAt:   "2024-09-05T12:24:16+07:00",
			},
			mockBehavior: func(ctx context.Context, repo *mock_repo.MockRepository) {
				todo := database.Todo{
					ID:          1,
					Title:       "test",
					Description: "test",
					DueDate:     "2024-09-05T12:40:16+07:00",
					CreatedAt:   "2024-09-05T12:24:16+07:00",
					UpdatedAt:   "2024-09-05T12:24:16+07:00",
				}
				todoID := int32(1)
				repo.EXPECT().GetTodo(ctx, todoID).Return(todo, nil).Times(1)
			},
		},
		{
			name:           "GetTodoHandler Invalid ID 1",
			input:          nil,
			reqMethod:      http.MethodGet,
			reqTarget:      "/tasks/a",
			expectedStatus: http.StatusBadRequest,
			expectedBody: struct {
				Error string `json:"error"`
			}{
				Error: middleware.ErrInvalidTodoID,
			},
			mockBehavior: func(ctx context.Context, repo *mock_repo.MockRepository) {},
		},
		{
			name:           "GetTodoHandler Invalid ID 2",
			input:          nil,
			reqMethod:      http.MethodGet,
			reqTarget:      "/tasks/-1",
			expectedStatus: http.StatusBadRequest,
			expectedBody: struct {
				Error string `json:"error"`
			}{
				Error: middleware.ErrInvalidTodoID,
			},
			mockBehavior: func(ctx context.Context, repo *mock_repo.MockRepository) {},
		},
		{
			name:           "GetTodoHandler Todo Not Found",
			input:          nil,
			reqMethod:      http.MethodGet,
			reqTarget:      "/tasks/11",
			expectedStatus: http.StatusNotFound,
			expectedBody: struct {
				Error string `json:"error"`
			}{
				Error: errTodoNotFound,
			},
			mockBehavior: func(ctx context.Context, repo *mock_repo.MockRepository) {
				todoID := int32(11)
				repo.EXPECT().GetTodo(ctx, todoID).Return(database.Todo{}, sql.ErrNoRows).Times(1)
			},
		},
		{
			name:           "GetTodoHandler Repo Error",
			input:          nil,
			reqMethod:      http.MethodGet,
			reqTarget:      "/tasks/1",
			expectedStatus: http.StatusInternalServerError,
			expectedBody: struct {
				Error string `json:"error"`
			}{
				Error: errGettingTodo,
			},
			mockBehavior: func(ctx context.Context, repo *mock_repo.MockRepository) {
				todoID := int32(1)
				repo.EXPECT().GetTodo(ctx, todoID).Return(database.Todo{}, errors.New("some db error")).Times(1)
			},
		},

		// UpdateHandler
		{
			name: "UpdateHandler Success",
			input: bytes.NewBuffer([]byte(`{
				"title": "test",
				"description": "test",
				"due_date": "2024-09-05T12:40:16+07:00"
			}`)),
			reqMethod:      http.MethodPut,
			reqTarget:      "/tasks/1",
			expectedStatus: http.StatusOK,
			expectedBody: dto.TodoResponseDto{
				ID:          1,
				Title:       "test",
				Description: "test",
				DueDate:     "2024-09-05T12:40:16+07:00",
				CreatedAt:   "2024-09-05T12:24:16+07:00",
				UpdatedAt:   "2024-09-05T12:24:16+07:00",
			},
			mockBehavior: func(ctx context.Context, repo *mock_repo.MockRepository) {
				todo := database.Todo{
					ID:          1,
					Title:       "test",
					Description: "test",
					DueDate:     "2024-09-05T12:40:16+07:00",
					CreatedAt:   "2024-09-05T12:24:16+07:00",
					UpdatedAt:   "2024-09-05T12:24:16+07:00",
				}
				updatedAt := time.Now().Format(time.RFC3339)
				todoID := int32(1)
				repo.EXPECT().UpdateTodo(ctx, database.UpdateTodoParams{
					ID:          todoID,
					Title:       "test",
					Description: "test",
					DueDate:     "2024-09-05T12:40:16+07:00",
					UpdatedAt:   updatedAt,
				}).Return(todo, nil).Times(1)
			},
		},
		{
			name: "UpdateHandler Invalid Input 1",
			input: bytes.NewBuffer([]byte(`{
				"title": "test",
				"description": "test",
				"due_date": "2024-09-05T12:40:16+07:00"
			}`)),
			reqMethod:      http.MethodPut,
			reqTarget:      "/tasks/a",
			expectedStatus: http.StatusBadRequest,
			expectedBody: struct {
				Error string `json:"error"`
			}{
				Error: middleware.ErrInvalidTodoID,
			},
			mockBehavior: func(ctx context.Context, repo *mock_repo.MockRepository) {},
		},
		{
			name: "UpdateHandler Invalid Input 2",
			input: bytes.NewBuffer([]byte(`{
				"title": "test",
				"description": "test",
				"due_date": "2024-09-05T12:40:16+07:00"
			}`)),
			reqMethod:      http.MethodPut,
			reqTarget:      "/tasks/-1",
			expectedStatus: http.StatusBadRequest,
			expectedBody: struct {
				Error string `json:"error"`
			}{
				Error: middleware.ErrInvalidTodoID,
			},
			mockBehavior: func(ctx context.Context, repo *mock_repo.MockRepository) {},
		},
		{
			name: "UpdateHandler Invalid Input 3",
			input: bytes.NewBuffer([]byte(`{
				"title": "test,
				"description": "test",
				"due_date": "2024-09-05T12:40:16+07:00"
			}`)),
			reqMethod:      http.MethodPut,
			reqTarget:      "/tasks/1",
			expectedStatus: http.StatusBadRequest,
			expectedBody: struct {
				Error string `json:"error"`
			}{
				Error: middleware.ErrInvalidInput,
			},
			mockBehavior: func(ctx context.Context, repo *mock_repo.MockRepository) {},
		},
		{
			name: "UpdateHandler Invalid Input 4",
			input: bytes.NewBuffer([]byte(`{
				"description": "test",
				"due_date": "2024-09-05T12:40:16+07:00"
			}`)),
			reqMethod:      http.MethodPut,
			reqTarget:      "/tasks/1",
			expectedStatus: http.StatusBadRequest,
			expectedBody: struct {
				Error string `json:"error"`
			}{
				Error: middleware.ErrInvalidInput,
			},
			mockBehavior: func(ctx context.Context, repo *mock_repo.MockRepository) {},
		},
		{
			name: "UpdateHandler Todo Not Found",
			input: bytes.NewBuffer([]byte(`{
				"title": "test",
				"description": "test",
				"due_date": "2024-09-05T12:40:16+07:00"
			}`)),
			reqMethod:      http.MethodPut,
			reqTarget:      "/tasks/11",
			expectedStatus: http.StatusNotFound,
			expectedBody: struct {
				Error string `json:"error"`
			}{
				Error: errTodoNotFound,
			},
			mockBehavior: func(ctx context.Context, repo *mock_repo.MockRepository) {
				updated := time.Now().Format(time.RFC3339)
				todoID := int32(11)
				repo.EXPECT().UpdateTodo(ctx, database.UpdateTodoParams{
					ID:          todoID,
					Title:       "test",
					Description: "test",
					DueDate:     "2024-09-05T12:40:16+07:00",
					UpdatedAt:   updated,
				}).Return(database.Todo{}, sql.ErrNoRows).Times(1)
			},
		},
		{
			name: "UpdateHandler Repo Error",
			input: bytes.NewBuffer([]byte(`{
				"title": "test",
				"description": "test",
				"due_date": "2024-09-05T12:40:16+07:00"
			}`)),
			reqMethod:      http.MethodPut,
			reqTarget:      "/tasks/1",
			expectedStatus: http.StatusInternalServerError,
			expectedBody: struct {
				Error string `json:"error"`
			}{
				Error: errUpdatingTodo,
			},
			mockBehavior: func(ctx context.Context, repo *mock_repo.MockRepository) {
				todoID := int32(1)
				updatedAt := time.Now().Format(time.RFC3339)
				repo.EXPECT().UpdateTodo(ctx, database.UpdateTodoParams{
					ID:          todoID,
					Title:       "test",
					Description: "test",
					DueDate:     "2024-09-05T12:40:16+07:00",
					UpdatedAt:   updatedAt,
				}).Return(database.Todo{}, errors.New("some db error")).Times(1)
			},
		},

		// DeleteHandler
		{
			name:           "DeleteHandler Success",
			input:          nil,
			reqMethod:      http.MethodDelete,
			reqTarget:      "/tasks/1",
			expectedStatus: http.StatusNoContent,
			expectedBody:   nil,
			mockBehavior: func(ctx context.Context, repo *mock_repo.MockRepository) {
				todoID := int32(1)
				repo.EXPECT().DeleteTodo(ctx, todoID).Return(database.Todo{}, nil).Times(1)
			},
		},
		{
			name:           "DeleteHandler Invalid ID 1",
			input:          nil,
			reqMethod:      http.MethodDelete,
			reqTarget:      "/tasks/a",
			expectedStatus: http.StatusBadRequest,
			expectedBody: struct {
				Error string `json:"error"`
			}{
				Error: middleware.ErrInvalidTodoID,
			},
			mockBehavior: func(ctx context.Context, repo *mock_repo.MockRepository) {},
		},
		{
			name:           "DeleteHandler Invalid ID 2",
			input:          nil,
			reqMethod:      http.MethodDelete,
			reqTarget:      "/tasks/-1",
			expectedStatus: http.StatusBadRequest,
			expectedBody: struct {
				Error string `json:"error"`
			}{
				Error: middleware.ErrInvalidTodoID,
			},
			mockBehavior: func(ctx context.Context, repo *mock_repo.MockRepository) {},
		},
		{
			name:           "DeleteHandler Todo Not Found",
			input:          nil,
			reqMethod:      http.MethodDelete,
			reqTarget:      "/tasks/11",
			expectedStatus: http.StatusNotFound,
			expectedBody: struct {
				Error string `json:"error"`
			}{
				Error: errTodoNotFound,
			},
			mockBehavior: func(ctx context.Context, repo *mock_repo.MockRepository) {
				todoID := int32(11)
				repo.EXPECT().DeleteTodo(ctx, todoID).Return(database.Todo{}, sql.ErrNoRows).Times(1)
			},
		},
		{
			name:           "DeleteHandler Repo Error",
			input:          nil,
			reqMethod:      http.MethodDelete,
			reqTarget:      "/tasks/1",
			expectedStatus: http.StatusInternalServerError,
			expectedBody: struct {
				Error string `json:"error"`
			}{
				Error: errDeletingTodo,
			},
			mockBehavior: func(ctx context.Context, repo *mock_repo.MockRepository) {
				todoID := int32(1)
				repo.EXPECT().DeleteTodo(ctx, todoID).Return(database.Todo{}, errors.New("some db error")).Times(1)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()

			repo := mock_repo.NewMockRepository(ctl)
			tt.mockBehavior(context.Background(), repo)

			s := service.NewService(repo)
			v, _ := validator.InitValidator()
			h := NewHandler(s, v)
			r := chi.NewRouter()
			h.RegisterRoutes(r)

			rec := httptest.NewRecorder()
			req := httptest.NewRequest(tt.reqMethod, tt.reqTarget, tt.input)
			req.Header.Set("Content-Type", "Application/Json")

			r.ServeHTTP(rec, req)
			res := rec.Result()
			defer res.Body.Close()
			data, _ := io.ReadAll(res.Body)
			jsonExpected, _ := json.Marshal(tt.expectedBody)

			require.Equal(t, jsonExpected, data)
			require.Equal(t, tt.expectedStatus, res.StatusCode)
		})
	}
}
