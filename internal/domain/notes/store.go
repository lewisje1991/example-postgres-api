package notes

import (
	"context"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/lewisje1991/code-bookmarks/internal/platform/postgres"
)

type Store struct {
	db postgres.DBTX
}

func NewStore(db postgres.DBTX) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) CreateNote(ctx context.Context, n *Note) (*Note, error) {
	queries := postgres.New(s.db)

	id := pgtype.UUID{
		Valid: true,
		Bytes: [16]byte(uuid.New()),
	}

	res, err := queries.CreateNote(ctx, postgres.CreateNoteParams{
		ID:        id,
		Title:     n.Title,
		Content:   n.Content,
		Tags:      strings.Join(n.Tags, ","),
		CreatedAt: pgtype.Timestamp{Time: time.Now().UTC(), Valid: true},
		UpdatedAt: pgtype.Timestamp{Time: time.Now().UTC(), Valid: true},
	})
	if err != nil {
		return nil, err
	}

	var tags []string
	if res.Tags != "" {
		tags = strings.Split(res.Tags, ",")
	}

	return &Note{
		ID:        uuid.UUID(res.ID.Bytes),
		Title:     res.Title,
		Content:   res.Content,
		Tags:      tags,
		CreatedAt: res.CreatedAt.Time,
		UpdatedAt: res.UpdatedAt.Time,
	}, nil
}

func (s *Store) GetNote(ctx context.Context, id uuid.UUID) (*Note, error) {
	queries := postgres.New(s.db)

	res, err := queries.GetNote(ctx, pgtype.UUID{
		Valid: true,
		Bytes: [16]byte(id),
	})
	if err != nil {
		return nil, err
	}

	var tags []string
	if res.Tags != "" {
		tags = strings.Split(res.Tags, ",")
	}

	return &Note{
		ID:        uuid.UUID(res.ID.Bytes),
		Title:     res.Title,
		Content:   res.Content,
		Tags:      tags,
		CreatedAt: res.CreatedAt.Time,
		UpdatedAt: res.UpdatedAt.Time,
	}, nil
}
