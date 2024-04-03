package notes

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	domain "github.com/lewisje1991/code-bookmarks/internal/domain/notes"
	"github.com/lewisje1991/code-bookmarks/internal/platform/server"
)

type Handler struct {
	service *domain.Service
	logger  *slog.Logger
}

func NewHandler(s *domain.Service, l *slog.Logger) *Handler {
	return &Handler{
		service: s,
		logger:  l,
	}
}

func (h *Handler) PostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req Request
		if err := server.Decode(r, &req); err != nil {
			h.logger.Error(fmt.Sprintf("error decoding request json: %v", err))
			server.EncodeError(w, http.StatusBadRequest, errors.New("invalid json"))
			return
		}

		savedNote, err := h.service.PostNote(r.Context(), &domain.Note{
			Title:   req.Title,
			Content: req.Content,
		})
		if err != nil {
			h.logger.Error(fmt.Sprintf("error saving note: %v", err))
			server.EncodeError(w, http.StatusInternalServerError, errors.New("error saving note"))
			return
		}

		server.EncodeData(w, http.StatusOK, Response{
			Title:     savedNote.Title,
			Content:   savedNote.Content,
			Tags:      savedNote.Tags,
			CreatedAt: time.Now().Local().Format(time.DateTime),
			UpdatedAt: time.Now().Local().Format(time.DateTime),
		})
	}
}

func (h *Handler) GetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	}
}
