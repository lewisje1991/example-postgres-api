package tasks

import (
	"github.com/lewisje1991/code-bookmarks/internal/web"
)

func AddRoutes(r *web.Router, h *Handler) {
	r.AddRoute("POST", "/task", h.PostTaskHandler())
	r.AddRoute("GET", "/task/{id}", h.GetTaskHandler())
}
