package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"

	apptasks "github.com/lewisje1991/code-bookmarks/cmd/api/tasks"
	"github.com/lewisje1991/code-bookmarks/internal/config"
	"github.com/lewisje1991/code-bookmarks/internal/logger"
	"github.com/lewisje1991/code-bookmarks/internal/postgres"
	domaintasks "github.com/lewisje1991/code-bookmarks/internal/tasks"
	"github.com/lewisje1991/code-bookmarks/internal/web"
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

	router := web.NewRouter()

	tasksStore := domaintasks.NewStore(db)
	tasksService := domaintasks.NewService(tasksStore)
	tasksHandler := apptasks.NewHandler(tasksService)
	apptasks.AddRoutes(router, tasksHandler)

	// Create and start server
	server := web.NewServer(config.HostPort, router)
	return server.Start()
}
