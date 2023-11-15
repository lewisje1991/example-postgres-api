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

func (s *Store) CreateBookmark(ctx context.Context, b *Bookmark) (*Bookmark, error) {
	queries := sqlite.New(s.db)

	bmk := sqlite.CreateBookmarkParams{
		ID:          uuid.New().String(),
		Url:         b.URL,
		Description: b.Description,
		Tags:        strings.Join(b.Tags, ","),
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}

	res, err := queries.CreateBookmark(ctx, bmk)
	if err != nil {
		return nil, fmt.Errorf("error executing post bookmark query: %w", err)
	}

	var tags []string
	if res.Tags != "" {
		tags = strings.Split(res.Tags, ",")
	}

	return &Bookmark{
		ID:          uuid.MustParse(res.ID),
		URL:         res.Url,
		Description: res.Description,
		Tags:        tags,
		CreatedAt:   res.CreatedAt.In(time.Local),
		UpdatedAt:   res.UpdatedAt.In(time.Local),
	}, nil
}

func (s *Store) GetBookmark(ctx context.Context, id uuid.UUID) (*Bookmark, error) {
	queries := sqlite.New(s.db)

	res, err := queries.GetBookmark(ctx, id.String())
	if err != nil {
		return nil, fmt.Errorf("error executing get bookmark query: %w", err)
	}

	var tags []string
	if res.Tags != "" {
		tags = strings.Split(res.Tags, ",")
	}

	return &Bookmark{
		ID:          uuid.MustParse(res.ID),
		URL:         res.Url,
		Description: res.Description,
		Tags:        tags,
		CreatedAt:   res.CreatedAt.In(time.Local),
		UpdatedAt:   res.UpdatedAt.In(time.Local),
	}, nil
}
