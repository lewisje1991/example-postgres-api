package api

import (
	"net/http"

	"github.com/go-chi/render"
)

func sendResponse(w http.ResponseWriter, r *http.Request, statusCode int, resp any) {
	render.Status(r, statusCode)
	render.JSON(w, r, resp)
}
