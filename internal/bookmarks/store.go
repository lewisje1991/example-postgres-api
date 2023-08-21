package bookmarks

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/lewisje1991/code-bookmarks/internal/platform/sqlite"
)

type Store struct {
	db sqlite.DBTX
}

func NewStore(db sqlite.DBTX) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) CreateBookmark(b *Bookmark) (*Bookmark, error) {
	queries := sqlite.New(s.db)

	bmk := sqlite.CreateBookmarkParams{
		ID:          uuid.New(),
		Url:         b.URL,
		Description: b.Description,
		Tags:        strings.Join(b.Tags, ","),
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}

	res, err := queries.CreateBookmark(context.Background(), bmk)
	if err != nil {
		return nil, fmt.Errorf("error executing post bookmark query: %w", err)
	}

	return &Bookmark{
		ID:          res.ID,
		URL:         res.Url,
		Description: res.Description,
		Tags:        strings.Split(res.Tags, ","),
		CreatedAt:   res.CreatedAt,
		UpdatedAt:   res.UpdatedAt,
	}, nil
}

func (s *Store) GetBookmark(id uuid.UUID) (*Bookmark, error) {
	queries := sqlite.New(s.db)

	res, err := queries.GetBookmark(context.Background(), id)
	if err != nil {
		return nil, fmt.Errorf("error executing get bookmark query: %w", err)
	}

	return &Bookmark{
		ID:          res.ID,
		URL:         res.Url,
		Description: res.Description,
		Tags:        strings.Split(res.Tags, ","),
		CreatedAt:   res.CreatedAt,
		UpdatedAt:   res.UpdatedAt,
	}, nil
}
