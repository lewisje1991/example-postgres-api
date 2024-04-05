package middleware

import (
	"fmt"
	"net/http"

	"github.com/lewisje1991/code-bookmarks/internal/foundation/server"
)

func IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "POST" {
			server.EncodeError(w, http.StatusUnauthorized, fmt.Errorf("unauthorized"))
			return
		}

		next.ServeHTTP(w, r)
	})
}
