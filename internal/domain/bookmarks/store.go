package bookmarks

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/lewisje1991/code-bookmarks/internal/foundation/postgres"
)

type Store struct {
	db postgres.DBTX
}

func NewStore(db postgres.DBTX) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) CreateBookmark(ctx context.Context, b *Bookmark) (*Bookmark, error) {
	queries := postgres.New(s.db)

	id := pgtype.UUID{
		Valid: true,
		Bytes: [16]byte(uuid.New()),
	}

	bmk := postgres.CreateBookmarkParams{
		ID:          id,
		Url:         b.URL,
		Description: b.Description,
		Tags:        strings.Join(b.Tags, ","),
		CreatedAt:   pgtype.Timestamp{Time: time.Now().UTC(), Valid: true},
		UpdatedAt:   pgtype.Timestamp{Time: time.Now().UTC(), Valid: true},
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
		ID:          uuid.UUID(res.ID.Bytes),
		URL:         res.Url,
		Description: res.Description,
		Tags:        tags,
		CreatedAt:   res.CreatedAt.Time,
		UpdatedAt:   res.UpdatedAt.Time,
	}, nil
}

func (s *Store) GetBookmark(ctx context.Context, id uuid.UUID) (*Bookmark, error) {
	queries := postgres.New(s.db)

	pgID := pgtype.UUID{
		Valid: true,
		Bytes: [16]byte(id),
	}

	res, err := queries.GetBookmark(ctx, pgID)
	if err != nil {
		return nil, fmt.Errorf("error executing get bookmark query: %w", err)
	}

	var tags []string
	if res.Tags != "" {
		tags = strings.Split(res.Tags, ",")
	}

	return &Bookmark{
		ID:          uuid.UUID(res.ID.Bytes),
		URL:         res.Url,
		Description: res.Description,
		Tags:        tags,
		CreatedAt:   res.CreatedAt.Time,
		UpdatedAt:   res.UpdatedAt.Time,
	}, nil
}
