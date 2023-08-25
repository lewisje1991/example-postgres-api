package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/lewisje1991/code-bookmarks/internal/api"
	"github.com/lewisje1991/code-bookmarks/internal/bookmarks"
	"github.com/lewisje1991/code-bookmarks/internal/platform/sqlite"
	"golang.org/x/exp/slog"
)

// TODO: config
// TODO: tests
// TODO: htmx
// TODO: db errors

func main() {
	r := chi.NewRouter()
	r.Use(middleware.AllowContentType("application/json"))

	config := NewConfig()
	if err := config.Load(".env"); err != nil {
		log.Fatal("Failed to load configuration: ", err)
	}

	mode := config.Mode

	var logger *slog.Logger
	if mode == "prod" {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	} else {
		logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	}

	logger.Info(fmt.Sprintf("running in %s mode", mode))

	db, err := sqlite.Connect(sqlite.BuildURL(sqlite.DbConfig{
		DBURL:   config.DBURL,
		DBToken: config.DBToken,
	}))

	if err != nil {
		logger.Error(fmt.Sprintf("failed to connect to db: %v", err))
		os.Exit(1)
	}

	if err := db.Ping(); err != nil {
		logger.Error(fmt.Sprintf("failed to ping db: %v", err))
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

	logger.Info(fmt.Sprintf("starting server on port:%d", config.HostPort))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.HostPort), r); err != nil {
		logger.Error(fmt.Sprintf("failed to start server: %v", err))
		os.Exit(1)
	}
}
