package handlers

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"log/slog"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/lewisje1991/code-bookmarks/internal/domain/bookmarks"
	"github.com/lewisje1991/code-bookmarks/internal/platform/server"
)

// go:generate mockgen -source=bookmarks.go -destination=mocks/bookmarks.go -package=mocks
type BookmarkService interface {
	GetBookmark(ctx context.Context, id uuid.UUID) (*bookmarks.Bookmark, error)
	PostBookmark(ctx context.Context, bookmark *bookmarks.Bookmark) (*bookmarks.Bookmark, error)
}

type BookmarkHandler struct {
	service BookmarkService
	logger  *slog.Logger
}

type BookmarkResponse struct {
	ID          string   `json:"id,omitempty"`
	URL         string   `json:"url,omitempty"`
	Description string   `json:"description,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	CreatedAt   string   `json:"createdAt,omitempty"`
	UpdatedAt   string   `json:"updatedAt,omitempty"`
}

type BookmarkRequest struct {
	URL         string   `json:"url"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

func (b *BookmarkRequest) Validate() error {
	if b.URL == "" {
		return fmt.Errorf("url is required")
	}

	if b.Description == "" {
		return fmt.Errorf("description is required")
	}
	return nil

}

func NewBookmarkHandler(l *slog.Logger, s BookmarkService) *BookmarkHandler {
	return &BookmarkHandler{
		service: s,
		logger:  l,
	}
}

func (h *BookmarkHandler) GetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		if idParam == "" {
			h.logger.Error("id is required")
			server.EncodeError(w, http.StatusBadRequest, errors.New("id is required"))
			return
		}

		bookmarkID, err := uuid.Parse(idParam)
		if err != nil {
			h.logger.Error(fmt.Sprintf("error parsing id: %v", err))
			server.EncodeError(w, http.StatusBadRequest, fmt.Errorf("invalid id: %v", err))
			return
		}

		bookmark, err := h.service.GetBookmark(r.Context(), bookmarkID)
		if err != nil {
			h.logger.Error(fmt.Sprintf("error getting bookmark: %v", err))
			server.EncodeError(w, http.StatusInternalServerError, errors.New("error getting bookmark"))
			return
		}

		if bookmark == nil {
			server.EncodeError(w, http.StatusNotFound, errors.New("bookmark not found"))
			return
		}

		server.EncodeData(w, http.StatusOK, BookmarkResponse{
			ID:          bookmark.ID.String(),
			URL:         bookmark.URL,
			Description: bookmark.Description,
			Tags:        bookmark.Tags,
			CreatedAt:   bookmark.CreatedAt.String(),
			UpdatedAt:   bookmark.UpdatedAt.String(),
		})
	}
}

func (h *BookmarkHandler) PostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req BookmarkRequest
		if err := server.Decode(r, &req); err != nil {
			h.logger.Error(fmt.Sprintf("error decoding request: %v", err))
			server.EncodeError(w, http.StatusBadRequest, fmt.Errorf("error decoding request: %v", err))
			return
		}

		b := &bookmarks.Bookmark{
			URL:         req.URL,
			Tags:        req.Tags,
			Description: req.Description,
		}

		bookmark, err := h.service.PostBookmark(r.Context(), b)
		if err != nil {
			h.logger.Error(fmt.Sprintf("error creating bookmark: %v", err))
			server.EncodeError(w, http.StatusInternalServerError, errors.New("error creating bookmark"))
			return
		}

		server.EncodeData(w, http.StatusOK, BookmarkResponse{
			ID:          bookmark.ID.String(),
			URL:         bookmark.URL,
			Description: bookmark.Description,
			Tags:        bookmark.Tags,
			CreatedAt:   bookmark.CreatedAt.String(),
			UpdatedAt:   bookmark.UpdatedAt.String(),
		})
	}
}

func sendResponse(w http.ResponseWriter, r *http.Request, statusCode int, resp any) {
	render.Status(r, statusCode)
	render.JSON(w, r, resp)
}
