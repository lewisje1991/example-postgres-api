package notes

import (
	"context"

	"github.com/google/uuid"
)

type Storer interface {
	CreateNote(ctx context.Context, n *Note) (*Note, error)
	GetNote(ctx context.Context, id uuid.UUID) (*Note, error)
}

type Service struct {
	store Storer
}

func NewService(s Storer) *Service {
	return &Service{
		store: s,
	}
}

func (s *Service) PostNote(ctx context.Context, n *Note) (*Note, error) {
	note, err := s.store.CreateNote(ctx, n)
	if err != nil {
		return nil, err
	}
	return note, nil
}

func (s *Service) GetNote(ctx context.Context, id uuid.UUID) (*Note, error) {
	n, err := s.store.GetNote(ctx, id)
	if err != nil {
		return nil, err
	}
	return n, nil
}
