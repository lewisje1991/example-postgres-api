package bookmarks

import (
	"github.com/lewisje1991/code-bookmarks/internal/foundation/server"
)

func AddRoutes(server *server.Server, bh *Handler) {
	server.AddRoute("POST", "/bookmark", bh.PostHandler())
	server.AddRoute("GET", "/bookmark/{id}", bh.GetHandler())
}
