package task

import "github.com/google/uuid"

type Response struct {
	ID     uuid.UUID
	Name   string
	Status string
}
