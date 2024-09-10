package diary

import (
	"context"
	"fmt"
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

func (s *Store) CreateDiary(ctx context.Context, diary Diary) (Diary, error) {
	queries := postgres.New(s.db)

	res, err := queries.CreateDiary(ctx, postgres.CreateDiaryParams{
		ID: pgtype.UUID{
			Bytes: [16]byte(uuid.New()),
			Valid: true,
		},
		Day: pgtype.Date{
			Time:  diary.Day,
			Valid: true,
		},
		CreatedAt: pgtype.Timestamp{Time: time.Now().UTC(), Valid: true},
		UpdatedAt: pgtype.Timestamp{Time: time.Now().UTC(), Valid: true},
	})

	if err != nil {
		return Diary{}, fmt.Errorf("failed to execute create diary query: %w", err)
	}

	return Diary{
		ID:  uuid.UUID(res.ID.Bytes),
		Day: res.Day.Time,
	}, nil

}
