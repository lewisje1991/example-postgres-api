package diary

import (
	"github.com/lewisje1991/code-bookmarks/internal/foundation/server"
)

func AddRoutes(server *server.Server, h *Handler) {
	server.AddRoute("POST", "/diary", h.NewDiaryEntryHandler())
	server.AddRoute("POST", "/diary/{diaryID}/tasks", h.CreateTaskForDiaryHandler())
}
