package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/lewisje1991/code-bookmarks/internal/domain/bookmarks"
	"golang.org/x/exp/slog"
)

type BookmarkHandler struct {
	service *bookmarks.Service
	logger  *slog.Logger
}

func NewBookmarkHandler(logger *slog.Logger, s *bookmarks.Service) *BookmarkHandler {
	return &BookmarkHandler{
		service: s,
		logger:  logger,
	}
}

func (h *BookmarkHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		if idParam == "" {
			h.logger.Error("id is required")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("id is required"))
			return
		}

		bookmarkID, err := uuid.Parse(idParam)
		if err != nil {
			h.logger.Error(fmt.Sprintf("error parsing id: %v", err))
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid id"))
			return
		}

		bookmark, err := h.service.GetBookmark(bookmarkID)
		if err != nil {
			h.logger.Error(fmt.Sprintf("error getting bookmark: %v", err))
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error getting bookmark"))
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, bookmark)
	}
}

func (h *BookmarkHandler) Post() http.HandlerFunc {
	type request struct {
		URL         string   `json:"url"`
		Description string   `json:"description"`
		Tags        []string `json:"tags"`
	}

	validate := func(r request) error {
		if r.URL == "" {
			return fmt.Errorf("url is required")
		}
		return nil
	}

	type response struct {
		ID          string   `json:"id,omitempty"`
		URL         string   `json:"url,omitempty"`
		Description string   `json:"description,omitempty"`
		Tags        []string `json:"tags,omitempty"`
		CreatedAt   string   `json:"createdAt,omitempty"`
		UpdatedAt   string   `json:"updatedAt,omitempty"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var req request
		if err := render.DecodeJSON(r.Body, &req); err != nil {
			h.logger.Error(fmt.Sprintf("error decoding request json: %v", err))
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid json"))
			return
		}

		if err := validate(req); err != nil {
			h.logger.Error(fmt.Sprintf("error validating bookmark request: %v", err))
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid request"))
			return
		}

		internalBookmark := &bookmarks.Bookmark{
			URL:         req.URL,
			Tags:        req.Tags,
			Description: req.Description,
		}

		bookmark, err := h.service.PostBookmark(internalBookmark)
		if err != nil {
			h.logger.Error(fmt.Sprintf("error creating bookmark: %v", err))
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error creating bookmark"))
			return
		}

		render.JSON(w, r, response{
			ID:          bookmark.ID.String(),
			URL:         bookmark.URL,
			Description: bookmark.Description,
			Tags:        bookmark.Tags,
			CreatedAt:   bookmark.CreatedAt.String(),
			UpdatedAt:   bookmark.UpdatedAt.String(),
		})
	}
}
