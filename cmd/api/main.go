package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/lewisje1991/code-bookmarks/internal/api"
	"github.com/lewisje1991/code-bookmarks/internal/bookmarks"
	"github.com/lewisje1991/code-bookmarks/internal/platform/sqlite"
	"golang.org/x/exp/slog"
)

// TODO: deployment
// TODO: tests
// TODO: htmx

func main() {
	r := chi.NewRouter()
	r.Use(middleware.AllowContentType("application/json"))

	var logger *slog.Logger
	if env := os.Getenv("ENV"); env == "prod" {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	} else {
		logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	}

	db, err := sqlite.Connect("file:file.db")
	if err != nil {
		logger.Error(fmt.Sprintf("failed to connect to db: %v", err))
		os.Exit(1)
	}

	bookmarksStore := bookmarks.NewStore(db)
	booksmarksService := bookmarks.NewService(bookmarksStore, logger)

	r.Post("/bookmark", api.PostBookmarkHandler(logger, booksmarksService))
	r.Get("/bookmark/{id}", api.GetBookmarkHandler(logger, booksmarksService))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("route does not exist"))
	})

	logger.Info("starting server on port:")
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
