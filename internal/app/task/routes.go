package task

import (
	"github.com/lewisje1991/code-bookmarks/internal/foundation/server"
)

func AddRoutes(server *server.Server, h *Handler) {
	// Notes
	server.AddRoute("POST", "/note", h.PostHandler())
}
