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
		CreatedAt:   time.Now().UTC().Format(time.DateTime),
		UpdatedAt:   time.Now().UTC().Format(time.DateTime),
	}

	res, err := queries.CreateBookmark(context.Background(), bmk) //TODO add context
	if err != nil {
		return nil, fmt.Errorf("error executing post bookmark query: %w", err)
	}

	createdAt, err := time.Parse(time.DateTime, res.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("error parsing created at time: %w", err)
	}

	updatedAt, err := time.Parse(time.DateTime, res.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("error parsing updated at time: %w", err)
	}

	return &Bookmark{
		ID:          res.ID,
		URL:         res.Url,
		Description: res.Description,
		Tags:        strings.Split(res.Tags, ","),
		CreatedAt:   createdAt.In(time.Local),
		UpdatedAt:   updatedAt.In(time.Local),
	}, nil
}

func (s *Store) GetBookmark(id uuid.UUID) (*Bookmark, error) {
	queries := sqlite.New(s.db)

	res, err := queries.GetBookmark(context.Background(), id)
	if err != nil {
		return nil, fmt.Errorf("error executing get bookmark query: %w", err)
	}

	createdAt, err := time.Parse(time.RFC3339, res.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("error parsing created at time: %w", err)
	}

	updatedAt, err := time.Parse(time.RFC3339, res.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("error parsing updated at time: %w", err)
	}

	return &Bookmark{
		ID:          res.ID,
		URL:         res.Url,
		Description: res.Description,
		Tags:        strings.Split(res.Tags, ","),
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}, nil
}
