package diary

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) NewDiaryEntry(ctx context.Context) Diary {
	return Diary{
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
}
