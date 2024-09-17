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

func (s *Service) NewDiaryEntry(ctx context.Context, today time.Time) (DiaryWithTasks, error) {
	existingEntry, err := s.store.GetDiaryByDay(ctx, today)
	if err != nil {
		return DiaryWithTasks{}, fmt.Errorf("failed to check for existing diary: %w", err)
	}

	if existingEntry.ID != uuid.Nil {
		return DiaryWithTasks{Diary: existingEntry}, nil
	}
	
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
