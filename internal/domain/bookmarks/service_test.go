package bookmarks_test

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/lewisje1991/code-bookmarks/internal/domain/bookmarks"
	"github.com/lewisje1991/code-bookmarks/internal/domain/bookmarks/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_PostBookmark(t *testing.T) {
	t.Run("returns error on store error", func(t *testing.T) {
		mockBookmarkStore := new(mocks.BookmarkStore)
		mockBookmarkStore.On("CreateBookmark", mock.Anything, mock.Anything).Return(nil, errors.New("error creating bookmark"))
		service := bookmarks.NewService(mockBookmarkStore)
		_, err := service.PostBookmark(context.Background(), &bookmarks.Bookmark{})
		assert.Error(t, err)
	})

	t.Run("returns bookmark on success", func(t *testing.T) {
		mockBookmarkStore := new(mocks.BookmarkStore)
		mockBookmarkStore.On("CreateBookmark", mock.Anything, mock.Anything).Return(&bookmarks.Bookmark{}, nil)
		service := bookmarks.NewService(mockBookmarkStore)
		response, err := service.PostBookmark(context.Background(), &bookmarks.Bookmark{})
		assert.NoError(t, err)
		assert.NotNil(t, response)
	})
}

func TestService_GetBookmark(t *testing.T) {
	t.Run("returns error on store error", func(t *testing.T) {
		mockStore := new(mocks.BookmarkStore)
		mockStore.On("GetBookmark", mock.Anything, mock.Anything).Return(nil, errors.New("error getting bookmark"))
		service := bookmarks.NewService(mockStore)
		_, err := service.GetBookmark(context.Background(), uuid.MustParse("3b1cf807-c743-43ef-bb93-cf7834bf5ca4"))
		assert.Error(t, err)
	})

	t.Run("returns bookmark on success", func(t *testing.T) {
		mockStore := new(mocks.BookmarkStore)
		mockStore.On("CreateBookmark", mock.Anything, mock.Anything).Return(&bookmarks.Bookmark{}, nil)
		service := bookmarks.NewService(mockStore)
		response, err := service.PostBookmark(context.Background(), &bookmarks.Bookmark{})
		assert.NoError(t, err)
		assert.NotNil(t, response)
	})
}
