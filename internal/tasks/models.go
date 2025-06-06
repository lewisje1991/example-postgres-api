package tasks

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID        uuid.UUID
	Title     string
	Content   string
	Status    string
	Tags      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
