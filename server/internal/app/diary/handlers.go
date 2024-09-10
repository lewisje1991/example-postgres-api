package diary

import (
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

func (h *Handler) PostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		diaryEntry := h.service.NewDiaryEntry(r.Context())

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
