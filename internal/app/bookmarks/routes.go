package bookmarks

import (
	"github.com/lewisje1991/code-bookmarks/internal/foundation/middleware"
	"github.com/lewisje1991/code-bookmarks/internal/foundation/server"
)

func AddRoutes(server *server.Server, bh *Handler, secret string) {
	server.AddRoute("POST", "/bookmark", middleware.IsAuthenticated(secret, bh.PostHandler()))
	server.AddRoute("GET", "/bookmark/{id}", bh.GetHandler())
}
