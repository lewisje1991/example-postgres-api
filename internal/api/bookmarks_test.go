package api_test

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/lewisje1991/code-bookmarks/internal/api"
	"github.com/lewisje1991/code-bookmarks/internal/api/mocks"
	"github.com/lewisje1991/code-bookmarks/internal/domain/bookmarks"
	"github.com/stretchr/testify/assert"
)

func TestBookmarks_Get(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	t.Run("id is required", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/bookmarks/123", nil)

		rr := httptest.NewRecorder()
		bookmarkHandler := api.NewBookmarkHandler(logger, nil)
		handler := bookmarkHandler.Get()

		handler.ServeHTTP(rr, req)

		want := api.BookmarkResponse{
			Error: "id is required",
		}

		got := api.BookmarkResponse{}
		err := json.Unmarshal(rr.Body.Bytes(), &got)
		assert.NoError(t, err)
		assert.Equal(t, want, got)
		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("id is valid guid", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/bookmarks/{id}", nil)

		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "invalid-guid")

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		rr := httptest.NewRecorder()
		bookmarkHandler := api.NewBookmarkHandler(logger, nil)
		handler := bookmarkHandler.Get()

		handler.ServeHTTP(rr, req)

		want := api.BookmarkResponse{
			Error: "invalid id: invalid UUID length: 12",
		}

		got := api.BookmarkResponse{}
		err := json.Unmarshal(rr.Body.Bytes(), &got)
		assert.NoError(t, err)
		assert.Equal(t, want, got)
		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("error getting bookmark", func(t *testing.T) {
		mockBookmarkService := new(mocks.BookmarkService)
		mockBookmarkService.On("GetBookmark", uuid.MustParse("3b1cf807-c743-43ef-bb93-cf7834bf5ca4")).Return(nil, fmt.Errorf("error getting bookmark"))

		bookmarkHandler := api.NewBookmarkHandler(logger, mockBookmarkService)
		handler := bookmarkHandler.Get()

		req := httptest.NewRequest("GET", "/bookmarks/{id}", nil)
		reqCtx := chi.NewRouteContext()
		reqCtx.URLParams.Add("id", "3b1cf807-c743-43ef-bb93-cf7834bf5ca4")

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, reqCtx))

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		want := api.BookmarkResponse{
			Error: "error getting bookmark",
		}

		got := api.BookmarkResponse{}
		err := json.Unmarshal(rr.Body.Bytes(), &got)
		assert.NoError(t, err)
		assert.Equal(t, want, got)
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})

	t.Run("bookmark is not found", func(t *testing.T) {
		mockBookmarkService := new(mocks.BookmarkService)
		mockBookmarkService.On("GetBookmark", uuid.MustParse("3b1cf807-c743-43ef-bb93-cf7834bf5ca4")).Return(nil, nil)

		bookmarkHandler := api.NewBookmarkHandler(logger, mockBookmarkService)
		handler := bookmarkHandler.Get()

		req := httptest.NewRequest("GET", "/bookmarks/{id}", nil)
		reqCtx := chi.NewRouteContext()
		reqCtx.URLParams.Add("id", "3b1cf807-c743-43ef-bb93-cf7834bf5ca4")

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, reqCtx))

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		want := api.BookmarkResponse{
			Error: "bookmark not found",
		}

		got := api.BookmarkResponse{}
		err := json.Unmarshal(rr.Body.Bytes(), &got)
		assert.NoError(t, err)
		assert.Equal(t, want, got)
		assert.Equal(t, http.StatusNotFound, rr.Code)
	})

	t.Run("bookmark is found", func(t *testing.T) {
		mockBookmarkService := new(mocks.BookmarkService)
		mockBookmark := &bookmarks.Bookmark{
			ID:          uuid.MustParse("3b1cf807-c743-43ef-bb93-cf7834bf5ca4"),
			URL:         "https://example.com",
			Description: "example",
			Tags:        []string{"example"},
			CreatedAt:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		}
		mockBookmarkService.On("GetBookmark", uuid.MustParse("3b1cf807-c743-43ef-bb93-cf7834bf5ca4")).Return(mockBookmark, nil)

		bookmarkHandler := api.NewBookmarkHandler(logger, mockBookmarkService)
		handler := bookmarkHandler.Get()

		req := httptest.NewRequest("GET", "/bookmarks/{id}", nil)
		reqCtx := chi.NewRouteContext()
		reqCtx.URLParams.Add("id", "3b1cf807-c743-43ef-bb93-cf7834bf5ca4")

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, reqCtx))

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		want := api.BookmarkResponse{
			Data: api.BookmarkResponseData{
				ID:          "3b1cf807-c743-43ef-bb93-cf7834bf5ca4",
				URL:         "https://example.com",
				Description: "example",
				Tags:        []string{"example"},
				CreatedAt:   "2020-01-01 00:00:00 +0000 UTC",
				UpdatedAt:   "2020-01-01 00:00:00 +0000 UTC",
			},
		}

		got := api.BookmarkResponse{}
		err := json.Unmarshal(rr.Body.Bytes(), &got)
		assert.NoError(t, err)

		assert.Equal(t, want, got)
		assert.Equal(t, http.StatusOK, rr.Code)
	})
}
