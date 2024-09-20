package tasks

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/lewisje1991/code-bookmarks/internal/foundation/postgres"
)

type Store struct {
	db postgres.DBTX
}

func NewStore(db postgres.DBTX) *Store {
	return &Store{db: db}
}

func (s *Store) CreateTask(ctx context.Context, task Task) (Task, error) {
	queries := postgres.New(s.db)

	res, err := queries.CreateTask(ctx, postgres.CreateTaskParams{
		ID: pgtype.UUID{
			Bytes: task.ID,
			Valid: true,
		},
		Title: task.Title,
		Tags:  task.Tags,
		CreatedAt: pgtype.Timestamp{
			Time:  task.CreatedAt,
			Valid: true,
		},
		UpdatedAt: pgtype.Timestamp{
			Time:  task.UpdatedAt,
			Valid: true,
		},
	})

	if err != nil {
		return Task{}, fmt.Errorf("failed to execute create task query: %w", err)
	}

	return taskFromDB(res), nil
}

func (s *Store) GetTask(ctx context.Context, id uuid.UUID) (Task, error) {
	queries := postgres.New(s.db)

	res, err := queries.GetTask(ctx, pgtype.UUID{
		Bytes: id,
		Valid: true,
	})

	if err != nil {
		return Task{}, fmt.Errorf("failed to execute get task query: %w", err)
	}

	return taskFromDB(res), nil
}

func (s *Store) GetTasksByDiaryID(ctx context.Context, diaryID uuid.UUID) ([]Task, error) {
	queries := postgres.New(s.db)

	dbTasks, err := queries.GetTasksByDiary(ctx, pgtype.UUID{
		Bytes: diaryID,
		Valid: true,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to execute get tasks by diary query: %w", err)
	}

	var tasks []Task
	for _, dbTask := range dbTasks {
		tasks = append(tasks, taskFromDB(dbTask))
	}

	return tasks, nil
}

func taskFromDB(dbTask postgres.Task) Task {
	return Task{
		ID:        uuid.UUID(dbTask.ID.Bytes),
		Title:     dbTask.Title,
		Tags:      dbTask.Tags,
		CreatedAt: dbTask.CreatedAt.Time,
		UpdatedAt: dbTask.UpdatedAt.Time,
	}
}
