package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"log/slog"

	"github.com/lewisje1991/code-bookmarks/internal/api/handlers"
	"github.com/lewisje1991/code-bookmarks/internal/api/router"
	"github.com/lewisje1991/code-bookmarks/internal/domain/bookmarks"
	"github.com/lewisje1991/code-bookmarks/internal/domain/notes"
	"github.com/lewisje1991/code-bookmarks/internal/platform/config"
	"github.com/lewisje1991/code-bookmarks/internal/platform/postgres"
)

// TODO: setup fly.io api hosting
// TODO: validate tokens with supabase
// TODO: only have access to your own data
// TODO: use expo to build a mobile app

func main() {
	ctx := context.Background()

	config := config.NewConfig()
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

	db, err := postgres.Connect(ctx, config.DBURL)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to connect to db: %v", err))
		os.Exit(1)
	}
	defer db.Close()

	bookmarksStore := bookmarks.NewStore(db)
	booksmarksService := bookmarks.NewService(bookmarksStore)
	booksmarksHandler := handlers.NewBookmarkHandler(logger, booksmarksService)

	notesStore := notes.NewStore(db)
	notesService := notes.NewService(notesStore)
	notesHandler := handlers.NewNotesHandler(notesService, logger)

	router := router.Routes(booksmarksHandler, notesHandler)

	logger.Info(fmt.Sprintf("starting server on port:%d", config.HostPort))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.HostPort), router); err != nil {
		logger.Error(fmt.Sprintf("failed to start server: %v", err))
		os.Exit(1)
	}
}
