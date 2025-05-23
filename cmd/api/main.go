package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"log/slog"

	apptasks "github.com/lewisje1991/code-bookmarks/internal/app/tasks"
	domaintasks "github.com/lewisje1991/code-bookmarks/internal/domain/tasks"

	"github.com/lewisje1991/code-bookmarks/internal/foundation/config"
	"github.com/lewisje1991/code-bookmarks/internal/foundation/logger"
	"github.com/lewisje1991/code-bookmarks/internal/foundation/postgres"
	"github.com/lewisje1991/code-bookmarks/internal/foundation/server"
)

// TODO: authorization/RBAC
// TODO: import bookmarks via export html
// TODO: openai integration for tagging

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

	logger.InitLogger(mode, slog.LevelInfo)
	slog.Info(fmt.Sprintf("running in %s mode", mode))

	db, err := postgres.Connect(ctx, config.DBURL)
	if err != nil {
		return fmt.Errorf("failed to connect to db: %v", err)
	}
	defer db.Close()

	server := server.NewServer()

	tasksStore := domaintasks.NewStore(db)
	tasksService := domaintasks.NewService(tasksStore)
	tasksHandler := apptasks.NewHandler(tasksService)
	apptasks.AddRoutes(server, tasksHandler)

	slog.Info(fmt.Sprintf("starting server on port:%d", config.HostPort))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.HostPort), server); err != nil {
		return fmt.Errorf("failed to start server: %v", err)
	}
	return nil
}
