-- name: CreateTodo :one
INSERT INTO todos (title, description, due_date, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetTodos :many
SELECT * FROM todos;

-- name: GetTodo :one
SELECT * FROM todos
WHERE id = $1;

-- name: UpdateTodo :one
UPDATE todos
SET title = $2, description = $3, due_date = $4, updated_at = $5
WHERE id = $1
RETURNING *;

-- name: DeleteTodo :one
DELETE FROM todos
WHERE id = $1
RETURNING *;