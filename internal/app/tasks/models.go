package tasks

import (
	"time"

	"github.com/google/uuid"
)

type Response struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Status    string    `json:"status"`
	Tags      string    `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateTaskRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  string `json:"status"`
	Tags    string `json:"tags"`
}

type UpdateTaskRequest struct {
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
	Status  string `json:"status,omitempty"`
	Tags    string `json:"tags,omitempty"`
}
