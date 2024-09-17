package tasks

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	store *Store
}

func NewService(s *Store) *Service {
	return &Service{
		store: s,
	}
}

func (s *Service) GetTasksForDiary(ctx context.Context, diaryID uuid.UUID) ([]Task, error) {
	tasks, err := s.store.GetTasksByDiaryID(ctx, diaryID)
	if err != nil {
		return nil, fmt.Errorf("failed to get tasks for diary: %w", err)
	}
	return tasks, nil
}

func (s *Service) CreateTask(ctx context.Context, task Task) (Task, error) {
	now := time.Now().UTC()
	task.CreatedAt = now
	task.UpdatedAt = now
	task.ID = uuid.New()
	return s.store.CreateTask(ctx, task)
}
