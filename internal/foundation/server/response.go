package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error,omitempty"`
}

func EncodeError(w http.ResponseWriter, status int, err error) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()}); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}
	return nil
}

func EncodeData[T any](w http.ResponseWriter, status int, v T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	data := struct {
		Data T `json:"data,omitempty"`
	}{
		Data: v,
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}
	return nil
}
