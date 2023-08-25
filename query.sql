-- name: CreateBookmark :one
INSERT INTO bookmarks (id, url, description, tags, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?) RETURNING *;

-- name: GetBookmark :one
SELECT * FROM bookmarks WHERE id = ?;