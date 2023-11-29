package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/lewisje1991/code-bookmarks/internal/api/handlers"
)

func Routes(bh *handlers.BookmarkHandler, nh *handlers.NotesHandler) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.AllowContentType("application/json"))

	// Bookmarks
	r.Post("/bookmark", bh.PostHandler())
	r.Get("/bookmark/{id}", bh.GetHandler())

	// Notes
	r.Post("/note", nh.PostHandler())

	// Misc
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
	})

	return r
}
