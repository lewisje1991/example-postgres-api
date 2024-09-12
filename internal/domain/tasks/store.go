package tasks

import (
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
