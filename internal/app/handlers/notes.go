package handlers

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/lewisje1991/code-bookmarks/internal/domain/notes"
	"github.com/lewisje1991/code-bookmarks/internal/platform/server"
)

type NoteService interface {
	PostNote(ctx context.Context, n *notes.Note) (*notes.Note, error)
	GetNote(ctx context.Context, id uuid.UUID) (*notes.Note, error)
}

type NotesHandler struct {
	service NoteService
	logger  *slog.Logger
}

func NewNotesHandler(s NoteService, l *slog.Logger) *NotesHandler {
	return &NotesHandler{
		service: s,
		logger:  l,
	}
}

type NoteResponseData struct {
	Title     string   `json:"title,omitempty"`
	Content   string   `json:"content,omitempty"`
	Tags      []string `json:"tags,omitempty"`
	CreatedAt string   `json:"createdAt,omitempty"`
	UpdatedAt string   `json:"updatedAt,omitempty"`
}

type NoteRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (n *NoteRequest) Validate() error {
	if n.Title == "" {
		return fmt.Errorf("title is required")
	}

	if n.Content == "" {
		return fmt.Errorf("content is required")
	}
	return nil
}

func (h *NotesHandler) PostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req NoteRequest
		if err := server.Decode(r, &req); err != nil {
			h.logger.Error(fmt.Sprintf("error decoding request json: %v", err))
			server.EncodeError(w, http.StatusBadRequest, errors.New("invalid json"))
			return
		}

		savedNote, err := h.service.PostNote(r.Context(), &notes.Note{
			Title:   req.Title,
			Content: req.Content,
		})
		if err != nil {
			h.logger.Error(fmt.Sprintf("error saving note: %v", err))
			server.EncodeError(w, http.StatusInternalServerError, errors.New("error saving note"))
			return
		}

		server.EncodeData(w, http.StatusOK, NoteResponseData{
			Title:     savedNote.Title,
			Content:   savedNote.Content,
			Tags:      savedNote.Tags,
			CreatedAt: time.Now().Local().Format(time.DateTime),
			UpdatedAt: time.Now().Local().Format(time.DateTime),
		})
	}
}

func (h *NotesHandler) GetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	}
}
