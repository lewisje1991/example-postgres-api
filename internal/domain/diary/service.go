package diary

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

func (s *Service) NewDiaryEntry(ctx context.Context) (DiaryWithTasks, error) {
	entity := Diary{
		ID:  uuid.New(),
		Day: time.Now(),
		Tasks: []Task{
			{
				ID:     uuid.NewString(),
				Name:   "Add diary endpoint",
				Status: "NEW",
			},
		},
	}

	dbEntity, err := s.store.CreateDiary(ctx, entity)
	if err != nil {
		return DiaryWithTasks{}, fmt.Errorf("failed to insert new diary entry: %w", err)
	}

	return DiaryWithTasks{
		Diary: dbEntity,
	}, nil
}
