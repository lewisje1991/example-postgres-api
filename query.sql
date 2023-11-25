-- name: CreateBookmark :one
INSERT INTO bookmarks (id, url, description, tags, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: GetBookmark :one
SELECT * FROM bookmarks WHERE id = $1;