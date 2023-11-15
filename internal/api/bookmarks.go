package api

import (
	"context"
	"fmt"
	"net/http"

	"log/slog"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/lewisje1991/code-bookmarks/internal/domain/bookmarks"
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
	Data  BookmarkResponseData `json:"data,omitempty"`
	Error string               `json:"error,omitempty"`
}

type BookmarkResponseData struct {
	ID          string   `json:"id,omitempty"`
	URL         string   `json:"url,omitempty"`
	Description string   `json:"description,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	CreatedAt   string   `json:"createdAt,omitempty"`
	UpdatedAt   string   `json:"updatedAt,omitempty"`
}

func NewBookmarkHandler(l *slog.Logger, s BookmarkService) *BookmarkHandler {
	return &BookmarkHandler{
		service: s,
		logger:  l,
	}
}

func (h *BookmarkHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		if idParam == "" {
			h.logger.Error("id is required")
			sendResponse(w, r, http.StatusBadRequest, BookmarkResponse{Error: "id is required"})
			return
		}

		bookmarkID, err := uuid.Parse(idParam)
		if err != nil {
			h.logger.Error(fmt.Sprintf("error parsing id: %v", err))
			sendResponse(w, r, http.StatusBadRequest, BookmarkResponse{Error: fmt.Sprintf("invalid id: %v", err)})
			return
		}

		bookmark, err := h.service.GetBookmark(r.Context(), bookmarkID)
		if err != nil {
			h.logger.Error(fmt.Sprintf("error getting bookmark: %v", err))
			sendResponse(w, r, http.StatusInternalServerError, BookmarkResponse{Error: "error getting bookmark"})
			return
		}

		if bookmark == nil {
			sendResponse(w, r, http.StatusNotFound, BookmarkResponse{Error: "bookmark not found"})
			return
		}

		sendResponse(w, r, http.StatusOK, BookmarkResponse{
			Data: BookmarkResponseData{
				ID:          bookmark.ID.String(),
				URL:         bookmark.URL,
				Description: bookmark.Description,
				Tags:        bookmark.Tags,
				CreatedAt:   bookmark.CreatedAt.String(),
				UpdatedAt:   bookmark.UpdatedAt.String(),
			},
		})
	}
}

func (h *BookmarkHandler) Post() http.HandlerFunc {
	type request struct {
		URL         string   `validate:"required" json:"url"`
		Description string   `validate:"required" json:"description"`
		Tags        []string `json:"tags"`
	}

	validate := func(req request) error {
		if req.URL == "" {
			return fmt.Errorf("url is required")
		}

		if req.Description == "" {
			return fmt.Errorf("description is required")
		}
		return nil
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var req request
		if err := render.DecodeJSON(r.Body, &req); err != nil {
			h.logger.Error(fmt.Sprintf("error decoding request json: %v", err))
			sendResponse(w, r, http.StatusBadRequest, BookmarkResponse{Error: "invalid json"})
			return
		}

		if err := validate(req); err != nil {
			h.logger.Error(fmt.Sprintf("error validating bookmark request: %v", err))
			sendResponse(w, r, http.StatusBadRequest, BookmarkResponse{Error: err.Error()})
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
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error creating bookmark"))
			return
		}

		sendResponse(w, r, http.StatusOK, BookmarkResponse{
			Data: BookmarkResponseData{
				ID:          bookmark.ID.String(),
				URL:         bookmark.URL,
				Description: bookmark.Description,
				Tags:        bookmark.Tags,
				CreatedAt:   bookmark.CreatedAt.String(),
				UpdatedAt:   bookmark.UpdatedAt.String(),
			},
		})
	}
}
