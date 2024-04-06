package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"log/slog"

	appbookmarks "github.com/lewisje1991/code-bookmarks/internal/app/bookmarks"
	appnotes "github.com/lewisje1991/code-bookmarks/internal/app/notes"
	domainbookmarks "github.com/lewisje1991/code-bookmarks/internal/domain/bookmarks"
	domainnotes "github.com/lewisje1991/code-bookmarks/internal/domain/notes"
	"github.com/lewisje1991/code-bookmarks/internal/foundation/config"
	"github.com/lewisje1991/code-bookmarks/internal/foundation/postgres"
	"github.com/lewisje1991/code-bookmarks/internal/foundation/server"
)

// TODO: authorization/RBAC
// TODO: use supabase cli to run locally
// TODO: openai integration for tagging
// TODO: use expo to build a mobile app

func main() {
	if err := Run(); err != nil {
		log.Fatalf("failed to run: %v", err)
	}
}

func Run() error {
	ctx := context.Background()

	config := config.NewConfig()
	if err := config.Load(".env"); err != nil {
		return fmt.Errorf("failed to load config: %v", err)
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
		return fmt.Errorf("failed to connect to db: %v", err)
	}
	defer db.Close()

	server := server.NewServer()

	bookmarksStore := domainbookmarks.NewStore(db)
	booksmarksService := domainbookmarks.NewService(bookmarksStore)
	booksmarksHandler := appbookmarks.NewHandler(logger, booksmarksService)
	appbookmarks.AddRoutes(server, booksmarksHandler, config.AuthSecret)

	notesStore := domainnotes.NewStore(db)
	notesService := domainnotes.NewService(notesStore)
	notesHandler := appnotes.NewHandler(notesService, logger)
	appnotes.AddRoutes(server, notesHandler)

	logger.Info(fmt.Sprintf("starting server on port:%d", config.HostPort))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.HostPort), server); err != nil {
		return fmt.Errorf("failed to start server: %v", err)
	}
	return nil
}
