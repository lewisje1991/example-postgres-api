package bookmarks

import (
	"github.com/lewisje1991/code-bookmarks/internal/foundation/middleware"
	"github.com/lewisje1991/code-bookmarks/internal/foundation/server"
)

func AddRoutes(server *server.Server, bh *Handler) {
	server.AddRoute("POST", "/bookmark", middleware.IsAuthenticated(bh.PostHandler()))
	server.AddRoute("GET", "/bookmark/{id}", bh.GetHandler())
}
