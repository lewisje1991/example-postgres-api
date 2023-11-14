package mocks

import (
	"github.com/google/uuid"
	"github.com/lewisje1991/code-bookmarks/internal/domain/bookmarks"
	"github.com/stretchr/testify/mock"
)

type BookmarkService struct {
	mock.Mock
}

func (b *BookmarkService) GetBookmark(id uuid.UUID) (*bookmarks.Bookmark, error) {
	args := b.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*bookmarks.Bookmark), args.Error(1)
}

func (b *BookmarkService) PostBookmark(bookmark *bookmarks.Bookmark) (*bookmarks.Bookmark, error) {
	args := b.Called(bookmark)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*bookmarks.Bookmark), args.Error(1)
}
