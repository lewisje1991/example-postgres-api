package bookmarks

import (
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/exp/slog"
)

type Service struct {
	store *Store
}

func NewService(s *Store, l *slog.Logger) *Service {
	return &Service{
		store: s,
	}
}

func (s *Service) PostBookmark(b *Bookmark) (*Bookmark, error) {
	bmk, err := s.store.CreateBookmark(b)
	if err != nil {
		return nil, fmt.Errorf("error creating bookmark: %w", err)
	}
	return bmk, nil
}

func (s *Service) GetBookmark(id uuid.UUID) (*Bookmark, error) {
	b, err := s.store.GetBookmark(id)
	if err != nil {
		return nil, fmt.Errorf("error getting bookmark: %w", err)
	}
	return b, nil
}
