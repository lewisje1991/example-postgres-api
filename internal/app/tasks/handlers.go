package tasks

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	domain "github.com/lewisje1991/code-bookmarks/internal/domain/tasks"
	"github.com/lewisje1991/code-bookmarks/internal/foundation/server"
)

type TaskService interface {
	CreateTask(ctx context.Context, task domain.Task) (*domain.Task, error)
	GetTask(ctx context.Context, id uuid.UUID) (*domain.Task, error)
}

type Handler struct {
	service TaskService
}

func NewHandler(s TaskService) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) PostTaskHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateTaskRequest
		if err := server.Decode(r, &req); err != nil {
			slog.Error("failed to decode request", "error", err)
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
			slog.Error("failed to create task", "error", err)
			server.EncodeError(w, http.StatusInternalServerError, err)
			return
		}

		resp := TaskResponse{
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

func (h *Handler) GetTaskHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		taskID := chi.URLParam(r, "id")
		if taskID == "" {
			slog.Error("task ID is required")
			server.EncodeError(w, http.StatusBadRequest, errors.New("task ID is required"))
			return
		}

		id, err := uuid.Parse(taskID)
		if err != nil {
			slog.Error("invalid task ID", "error", err)
			server.EncodeError(w, http.StatusBadRequest, errors.New("invalid task ID"))
			return
		}

		task, err := h.service.GetTask(r.Context(), id)
		if err != nil {
			slog.Error("failed to get task", "error", err)
			server.EncodeError(w, http.StatusInternalServerError, err)
			return
		}

		if task == nil {
			slog.Error("task not found")
			server.EncodeError(w, http.StatusNotFound, errors.New("task not found"))
			return
		}

		resp := TaskResponse{
			ID:        task.ID,
			Title:     task.Title,
			Content:   task.Content,
			Status:    task.Status,
			Tags:      task.Tags,
			CreatedAt: task.CreatedAt,
			UpdatedAt: task.UpdatedAt,
		}

		server.EncodeData(w, http.StatusOK, resp)
	}
}
