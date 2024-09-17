package tasks

import (
	"log/slog"
	"net/http"

	domain "github.com/lewisje1991/code-bookmarks/internal/domain/tasks"
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

func (h *Handler) NewTaskHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateTaskRequest
		if err := server.Decode(r, &req); err != nil {
			h.logger.Error("failed to decode request", "error", err)
			server.EncodeError(w, http.StatusBadRequest, err)
			return
		}

		task, err := h.service.CreateTask(r.Context(), domain.Task{
			Title:   req.Title,
			Content: req.Content,
			Status:  req.Status,
			Tags:    req.Tags,
		})
		if err != nil {
			h.logger.Error("failed to create task", "error", err)
			server.EncodeError(w, http.StatusInternalServerError, err)
			return
		}

		resp := Response{
			ID:        task.ID,
			Title:     task.Title,
			Content:   task.Content,
			Status:    task.Status,
			Tags:      task.Tags,
			CreatedAt: task.CreatedAt,
			UpdatedAt: task.UpdatedAt,
		}

		server.EncodeData(w, http.StatusCreated, resp)
	}
}
