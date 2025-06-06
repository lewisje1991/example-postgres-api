package middleware

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lewisje1991/code-bookmarks/internal/web"
)

func IsAuthenticated(jwtSecret string, next http.Handler) http.Handler {
	parseJWTToken := func(token string, hmacSecret []byte) (string, error) {
		t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return hmacSecret, nil
		}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}))

		if err != nil {
			return "", fmt.Errorf("error validating token: %v", err)
		}

		if !t.Valid {
			return "", errors.New("invalid token")
		}

		return t.Claims.GetSubject()
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwtToken := r.Header.Get("Authorization")
		if jwtToken == "" {
			web.EncodeError(w, http.StatusUnauthorized, errors.New("missing token"))
			return
		}

		jwtToken = strings.Split(jwtToken, "Bearer ")[1]

		userID, err := parseJWTToken(jwtToken, []byte(jwtSecret))
		if err != nil {
			log.Printf("Error parsing token: %s", err)
			web.EncodeError(w, http.StatusUnauthorized, errors.New("invalid token"))
			return
		}

		log.Printf("received request from userID:[%s]", userID)

		// Save the email in the context to use later in the handler
		ctx := context.WithValue(r.Context(), "userID", userID) // TODO use a key
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
