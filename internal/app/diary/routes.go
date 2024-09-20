package diary

import (
	"github.com/lewisje1991/code-bookmarks/internal/foundation/server"
)

func AddRoutes(server *server.Server, h *Handler) {
	server.AddRoute("POST", "/diary", h.NewDiaryHandler())
	// server.AddRoute("POST", "/diary/{id}/task/{id}", h.NewDiaryHandler())
}
