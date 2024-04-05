package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Server struct {
	mux *chi.Mux
}

func NewServer() *Server {
	mux := chi.NewRouter()
	mux.Use(middleware.AllowContentType("application/json"))

	// TODO add better place
	mux.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	})

	return &Server{
		mux: chi.NewRouter(),
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *Server) AddRoute(method, pattern string, handler http.Handler) {
	s.mux.Method(method, pattern, handler)
}
