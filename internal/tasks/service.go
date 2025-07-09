package tasks

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrTaskNotFound = errors.New("task not found")
	ErrInvalidStatus = errors.New("invalid task status")
)

type Service struct {
	store *Store
}

func NewService(s *Store) *Service {
	return &Service{
		store: s,
	}
}

func (s *Service) CreateTask(ctx context.Context, task Task) (*Task, error) {
	now := time.Now().UTC()
	task.CreatedAt = now
	task.UpdatedAt = now
	task.ID = uuid.New()
	return s.store.CreateTask(ctx, task)
}

func (s *Service) GetTask(ctx context.Context, id uuid.UUID) (*Task, error) {
	task, err := s.store.GetTask(ctx, id)
	if err != nil {
		return nil, err
	}

	if task == nil {
		return nil, ErrTaskNotFound
	}

	return task, nil
}
