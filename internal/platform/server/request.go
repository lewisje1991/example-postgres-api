package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
)

type validator interface {
	Validate() error
}

func Decode(r *http.Request, val any) error {
	if err := render.DecodeJSON(r.Body, val); err != nil {
		return fmt.Errorf("unable to decode payload: %w", err)
	}

	if v, ok := val.(validator); ok {
		if err := v.Validate(); err != nil {
			return err
		}
	}

	return nil
}
