package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"log/slog"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/lewisje1991/code-bookmarks/internal/api"
	"github.com/lewisje1991/code-bookmarks/internal/bookmarks"
	_ "github.com/libsql/libsql-client-go/libsql"
	_ "modernc.org/sqlite"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.AllowContentType("application/json"))

	var logger *slog.Logger
	if env := os.Getenv("ENV"); env == "prod" {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	} else {
		logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	}

	var dbUrl = "file:file.db"
	db, err := sql.Open("libsql", dbUrl)
	if err != nil {
		logger.Error("failed to open db: %v", err)
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", dbUrl, err)
		os.Exit(1)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		logger.Error("failed to ping db: %v", err)
		fmt.Fprintf(os.Stderr, "failed to ping db %s: %s", dbUrl, err)
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

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
