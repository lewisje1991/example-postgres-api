package diary

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	domain "github.com/lewisje1991/code-bookmarks/internal/domain/diary"
	"github.com/lewisje1991/code-bookmarks/internal/foundation/server"
)

type Handler struct {
	logger  *slog.Logger
	service *domain.Service
}

func NewHandler(l *slog.Logger, s *domain.Service) *Handler {
	return &Handler{
		service: s,
		logger:  l,
	}
}

func (h *Handler) NewDiaryEntryHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		diaryEntry, err := h.service.NewDiaryEntry(r.Context(), time.Now())
		if err != nil {
			h.logger.Error(fmt.Sprintf("error creating new entry: %v", err))
			server.EncodeError(w, http.StatusInternalServerError, fmt.Errorf("error creating new entry: %v", err))
			return
		}

		var tasks []Task
		for _, task := range diaryEntry.Tasks {
			t := Task{
				ID:     task.ID,
				Name:   task.Name,
				Status: task.Status,
			}
			tasks = append(tasks, t)
		}

		server.EncodeData(w, http.StatusOK, Response{
			ID:    diaryEntry.ID.String(),
			Date:  diaryEntry.Day.Format(time.DateOnly),
			Tasks: tasks,
		})
	}
}

func (h *Handler) CreateTaskForDiaryHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		diaryID := r.URL.Query().Get("diaryID")
		var task Task
		if err := server.Decode(r, &task); err != nil {
			h.logger.Error(fmt.Sprintf("error decoding task: %v", err))
			server.EncodeError(w, http.StatusBadRequest, fmt.Errorf("error decoding task: %v", err))
			return
		}

		task, err := h.service.CreateTaskForDiary(r.Context(), diaryID, task.Name)
		if err != nil {
			h.logger.Error(fmt.Sprintf("error creating task for diary: %v", err))
			server.EncodeError(w, http.StatusInternalServerError, fmt.Errorf("error creating task for diary: %v", err))
			return
		}
		server.EncodeData(w, http.StatusOK, task)
	}
}