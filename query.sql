-- name: CreateDiary :one
INSERT INTO diary (id, day, created_at, updated_at) 
VALUES ($1, $2, $3, $4) 
RETURNING *;

-- name: GetDiaryByDay :one
SELECT * 
FROM diary 
WHERE day = $1;

-- name: GetDiary :one
SELECT * 
FROM diary 
WHERE id = $1;

-- name: CreateTask :one
INSERT INTO tasks (id, title, description, tags, created_at, updated_at) 
VALUES ($1, $2, $3, $4, $5, $6) 
RETURNING *;

-- name: GetTask :one
SELECT * 
FROM tasks 
WHERE id = $1;

-- name: AddTaskToDiary :one
INSERT INTO diary_tasks (task_id, diary_id, status) 
VALUES ($1, $2, $3) 
RETURNING *;

-- name: GetTasksByDiary :many
SELECT tasks.* 
FROM tasks 
JOIN diary_tasks ON tasks.id = diary_tasks.task_id 
WHERE diary_tasks.diary_id = $1;

