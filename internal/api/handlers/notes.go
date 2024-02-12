package handlers

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/lewisje1991/code-bookmarks/internal/domain/notes"
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

type NoteResponse struct {
	Data  NoteResponseData `json:"data,omitempty"`
	Error string           `json:"error,omitempty"`
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

func (h *NotesHandler) PostHandler() http.HandlerFunc {
	validate := func(req NoteRequest) error {
		if req.Title == "" {
			return fmt.Errorf("title is required")
		}

		if req.Content == "" {
			return fmt.Errorf("content is required")
		}
		return nil
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var req NoteRequest
		if err := render.DecodeJSON(r.Body, &req); err != nil {
			h.logger.Error(fmt.Sprintf("error decoding request json: %v", err))
			sendResponse(w, r, http.StatusBadRequest, NoteResponse{Error: "invalid json"})
			return
		}

		if err := validate(req); err != nil {
			h.logger.Error(fmt.Sprintf("error validating note request: %v", err))
			sendResponse(w, r, http.StatusBadRequest, NoteResponse{Error: err.Error()})
			return
		}

		savedNote, err := h.service.PostNote(r.Context(), &notes.Note{
			Title:   req.Title,
			Content: req.Content,
		})
		if err != nil {
			h.logger.Error(fmt.Sprintf("error saving note: %v", err))
			sendResponse(w, r, http.StatusInternalServerError, NoteResponse{Error: "error saving note"})
			return
		}

		sendResponse(w, r, http.StatusOK, NoteResponse{
			Data: NoteResponseData{
				Title:     savedNote.Title,
				Content:   savedNote.Content,
				Tags:      savedNote.Tags,
				CreatedAt: time.Now().Local().Format(time.DateTime),
				UpdatedAt: time.Now().Local().Format(time.DateTime),
			},
		})
	}
}

func (h *NotesHandler) GetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	}
}
