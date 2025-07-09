package web

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Router struct {
	mux *chi.Mux
}

func NewRouter() *Router {
	mux := chi.NewRouter()
	
	// Add essential middleware
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Compress(5))
	mux.Use(middleware.AllowContentType("application/json"))
	
	// Add CORS middleware
	mux.Use(middleware.SetHeader("Access-Control-Allow-Origin", "*"))
	mux.Use(middleware.SetHeader("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS"))
	mux.Use(middleware.SetHeader("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token"))

	mux.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	})

	return &Router{
		mux: mux,
	}
}

func (ro *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ro.mux.ServeHTTP(w, r)
}

func (ro *Router) AddRoute(method, pattern string, handler http.Handler) {
	ro.mux.Method(method, pattern, handler)
}
