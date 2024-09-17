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

func (h *Handler) NewDiaryHandler() http.HandlerFunc {
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
				ID:     task.ID.String(),
				Name:   task.Title,
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