package tasks

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	domainTasks "github.com/lewisje1991/code-bookmarks/internal/tasks"
	"github.com/lewisje1991/code-bookmarks/internal/web"
)

type TaskService interface {
	CreateTask(ctx context.Context, task domainTasks.Task) (*domainTasks.Task, error)
	GetTask(ctx context.Context, id uuid.UUID) (*domainTasks.Task, error)
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
		if err := web.Decode(r, &req); err != nil {
			slog.Error("failed to decode request", "error", err)
			web.EncodeError(w, http.StatusBadRequest, err)
			return
		}

		task, err := h.service.CreateTask(r.Context(), req.ToDomain())
		if err != nil {
			slog.Error("failed to create task", "error", err)
			web.EncodeError(w, http.StatusInternalServerError, err)
			return
		}

		resp := TaskResponseFromDomain(task)

		web.EncodeData(w, http.StatusCreated, resp)
	}
}

func (h *Handler) GetTaskHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		taskID := chi.URLParam(r, "id")
		if taskID == "" {
			slog.Error("task ID is required")
			web.EncodeError(w, http.StatusBadRequest, errors.New("task ID is required"))
			return
		}

		id, err := uuid.Parse(taskID)
		if err != nil {
			slog.Error("invalid task ID", "error", err)
			web.EncodeError(w, http.StatusBadRequest, errors.New("invalid task ID"))
			return
		}

		task, err := h.service.GetTask(r.Context(), id)
		if err != nil {
			if errors.Is(err, domainTasks.ErrTaskNotFound) {
				web.EncodeError(w, http.StatusNotFound, err)
				return
			}
			slog.Error("failed to get task", "error", err)
			web.EncodeError(w, http.StatusInternalServerError, err)
			return
		}

		resp := TaskResponseFromDomain(task)

		web.EncodeData(w, http.StatusOK, resp)
	}
}
