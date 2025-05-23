package tasks

import (
	"github.com/lewisje1991/code-bookmarks/internal/foundation/server"
)

func AddRoutes(server *server.Server, h *Handler) {
	server.AddRoute("POST", "/task", h.PostTaskHandler())
	server.AddRoute("GET", "/task/{id}", h.GetTaskHandler())
}
