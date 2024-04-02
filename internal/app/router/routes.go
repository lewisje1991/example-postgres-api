package router

import (
	"github.com/lewisje1991/code-bookmarks/internal/app/handlers"
	"github.com/lewisje1991/code-bookmarks/internal/platform/server"
)

func AddRoutes(server *server.Server, bh *handlers.BookmarkHandler, nh *handlers.NotesHandler) {
	// Bookmarks
	server.AddRoute("POST", "/bookmark", bh.PostHandler())
	server.AddRoute("GET", "/bookmark/{id}", bh.GetHandler())

	// Notes
	server.AddRoute("POST", "/note", nh.PostHandler())
}
