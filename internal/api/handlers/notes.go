package handlers

import (
	"log/slog"
	"net/http"
)

type NotesHandler struct {
	logger *slog.Logger
}

func NewNotesHandler(l *slog.Logger) *NotesHandler {
	return &NotesHandler{
		logger: l,
	}
}

func (h *NotesHandler) PostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
