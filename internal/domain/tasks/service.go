package tasks

import (
	"context"
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

func (s *Service) CreateTask(ctx context.Context, task Task) (Task, error) {
	now := time.Now().UTC()
	task.CreatedAt = now
	task.UpdatedAt = now
	task.ID = uuid.New()
	return s.store.CreateTask(ctx, task)
}
