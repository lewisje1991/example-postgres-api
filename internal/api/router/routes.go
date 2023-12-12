package router

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/lewisje1991/code-bookmarks/internal/api/handlers"
	"github.com/lewisje1991/code-bookmarks/internal/templates/pages"
)

func Routes(bh *handlers.BookmarkHandler, nh *handlers.NotesHandler) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.AllowContentType("application/json"))

	r.Get("/", templ.Handler(pages.Homepage()).ServeHTTP)

	// Bookmarks
	r.Post("/bookmark", bh.PostHandler())
	r.Get("/bookmark/{id}", bh.GetHandler())

	// Notes
	r.Get("/note", nh.GetHandler())

	// Misc
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(pages.NotFound()).ServeHTTP(w, r)
		w.WriteHeader(404)
	})

	return r
}
