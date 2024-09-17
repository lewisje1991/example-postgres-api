-- name: CreateDiary :one
INSERT INTO diary (id, day, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetDiaryByDay :
SELECT * FROM diary WHERE day = $1;

-- name: GetDiary :one
SELECT * FROM diary WHERE id = $1;

-- name: CreateTask :one
INSERT INTO tasks (id, title, content, status, tags, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *;

-- name: UpdateTaskContent :one
UPDATE tasks SET content = $2, updated_at = $3 WHERE id = $1 RETURNING *;

-- name: UpdateTaskStatus :one
UPDATE tasks SET status = $2, updated_at = $3 WHERE id = $1 RETURNING *;

-- name: GetTask :one
SELECT * FROM tasks WHERE id = $1;

-- name: AddTaskToDiary :one
INSERT INTO task_diary (task_id, diary_id) VALUES ($1, $2) RETURNING *;

-- name: GetTasksByDiary :many
SELECT tasks.* FROM tasks JOIN task_diary ON tasks.id = task_diary.task_id WHERE task_diary.diary_id = $1;

