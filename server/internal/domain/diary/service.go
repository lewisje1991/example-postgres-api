package diary

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
func (s *Service) NewDiaryEntry(ctx context.Context) (Diary, error) {
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

	return s.store.CreateDiary(ctx, entity)
}
