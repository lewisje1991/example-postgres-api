-- name: CreateTask :one
INSERT INTO tasks (id, title, description, tags, created_at, updated_at) 
VALUES ($1, $2, $3, $4, $5, $6) 
RETURNING *;

-- name: GetTask :one
SELECT * 
FROM tasks 
WHERE id = $1;