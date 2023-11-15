package bookmarks

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type Storer interface {
	CreateBookmark(ctx context.Context, b *Bookmark) (*Bookmark, error)
	GetBookmark(ctx context.Context, id uuid.UUID) (*Bookmark, error)
}

type Service struct {
	store Storer
}

func NewService(s Storer) *Service {
	return &Service{
		store: s,
	}
}

func (s *Service) PostBookmark(ctx context.Context, b *Bookmark) (*Bookmark, error) {
	bmk, err := s.store.CreateBookmark(ctx, b)
	if err != nil {
		return nil, fmt.Errorf("error creating bookmark: %w", err)
	}
	return bmk, nil
}

func (s *Service) GetBookmark(ctx context.Context, id uuid.UUID) (*Bookmark, error) {
	b, err := s.store.GetBookmark(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error getting bookmark: %w", err)
	}
	return b, nil
}
