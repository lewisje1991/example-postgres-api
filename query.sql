-- name: CreateDiary :one
INSERT INTO diary (id, day, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING *;
