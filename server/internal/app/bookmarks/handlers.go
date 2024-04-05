package bookmarks

import (
	"errors"
	"fmt"
	"net/http"

	"log/slog"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	domain "github.com/lewisje1991/code-bookmarks/internal/domain/bookmarks"
	"github.com/lewisje1991/code-bookmarks/internal/foundation/server"
)

type Handler struct {
	service *domain.Service
	logger  *slog.Logger
}

func NewHandler(l *slog.Logger, s *domain.Service) *Handler {
	return &Handler{
		service: s,
		logger:  l,
	}
}

func (h *Handler) GetHandler() http.HandlerFunc {
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

		server.EncodeData(w, http.StatusOK, Response{
			ID:          bookmark.ID.String(),
			URL:         bookmark.URL,
			Description: bookmark.Description,
			Tags:        bookmark.Tags,
			CreatedAt:   bookmark.CreatedAt.String(),
			UpdatedAt:   bookmark.UpdatedAt.String(),
		})
	}
}

func (h *Handler) PostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req Request
		if err := server.Decode(r, &req); err != nil {
			h.logger.Error(fmt.Sprintf("error decoding request: %v", err))
			server.EncodeError(w, http.StatusBadRequest, fmt.Errorf("error decoding request: %v", err))
			return
		}

		b := req.ToBookmark()

		bookmark, err := h.service.PostBookmark(r.Context(), b)
		if err != nil {
			h.logger.Error(fmt.Sprintf("error creating bookmark: %v", err))
			server.EncodeError(w, http.StatusInternalServerError, errors.New("error creating bookmark"))
			return
		}

		server.EncodeData(w, http.StatusOK, Response{
			ID:          bookmark.ID.String(),
			URL:         bookmark.URL,
			Description: bookmark.Description,
			Tags:        bookmark.Tags,
			CreatedAt:   bookmark.CreatedAt.String(),
			UpdatedAt:   bookmark.UpdatedAt.String(),
		})
	}
}
