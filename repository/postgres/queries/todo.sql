-- name: GetTodos :many
SELECT * FROM todos;

-- name: GetTodoById :one
SELECT * FROM todos WHERE id = $1;

-- name: CreateTodo :one
INSERT INTO todos (title, description) VALUES ($1, $2) RETURNING *;

-- name: UpdateTodoById :exec
UPDATE todos SET title = $1, description = $2, completed = $3 WHERE id = $4;

-- name: DeleteTodoById :exec
DELETE FROM todos WHERE id = $1;