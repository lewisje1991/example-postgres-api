package api

import (
	"fmt"
	"net/http"

	"log/slog"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/lewisje1991/code-bookmarks/internal/bookmarks"
)

func GetBookmarkHandler(logger *slog.Logger, s *bookmarks.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		if idParam == "" {
			logger.Error("id is required")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("id is required"))
			return
		}

		bookmarkID, err := uuid.Parse(idParam)
		if err != nil {
			logger.Error(fmt.Sprintf("error parsing id: %v", err))
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid id"))
			return
		}

		bookmark, err := s.GetBookmark(bookmarkID)

		render.Status(r, http.StatusOK)
		render.JSON(w, r, bookmark)
	}
}

func PostBookmarkHandler(logger *slog.Logger, s *bookmarks.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req bookmarks.Bookmark
		if err := render.DecodeJSON(r.Body, &req); err != nil {
			logger.Error(fmt.Sprintf("error decoding requestjson: %v", err))
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid json"))
			return
		}

		if err := req.Validate(); err != nil {
			logger.Error(fmt.Sprintf("error validating bookmark request: %v", err))
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid request"))
			return
		}

		internalBookmark := &bookmarks.Bookmark{
			URL:         req.URL,
			Tags:        req.Tags,
			Description: req.Description,
		}

		bookmark, err := s.PostBookmark(internalBookmark)
		if err != nil {
			logger.Error(fmt.Sprintf("error creating bookmark: %v", err))
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error creating bookmark"))
			return
		}

		render.JSON(w, r, bookmark)
	}
}
