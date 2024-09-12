package tasks

import (
	"log/slog"
	"net/http"

	domain "github.com/lewisje1991/code-bookmarks/internal/domain/tasks"
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

func (h *Handler) PostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
