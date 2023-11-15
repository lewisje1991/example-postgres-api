package mocks

import (
	"context"

	"github.com/google/uuid"
	"github.com/lewisje1991/code-bookmarks/internal/domain/bookmarks"
	"github.com/stretchr/testify/mock"
)

type BookmarkStore struct {
	mock.Mock
}

func (b *BookmarkStore) CreateBookmark(context.Context, *bookmarks.Bookmark) (*bookmarks.Bookmark, error) {
	args := b.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*bookmarks.Bookmark), args.Error(1)
}

func (b *BookmarkStore) GetBookmark(context.Context, uuid.UUID) (*bookmarks.Bookmark, error) {
	args := b.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*bookmarks.Bookmark), args.Error(1)
}
