package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/lewisje1991/code-bookmarks/internal/api/handlers"
)

func Routes(booksmarksHandler handlers.BookmarkHandler) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.AllowContentType("application/json"))
	r.Post("/bookmark", booksmarksHandler.PostHandler())
	r.Get("/bookmark/{id}", booksmarksHandler.GetHandler())
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
	})
	return r
}
