package logger

import (
	"log/slog"
	"os"
)

func InitLogger(mode string, level slog.Level) {
	options := slog.HandlerOptions{
		Level: level,
	}
	var logger *slog.Logger
	if mode == "prod" {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &options))
	} else {
		logger = slog.New(slog.NewTextHandler(os.Stdout, &options))
	}
	slog.SetDefault(logger)
}
