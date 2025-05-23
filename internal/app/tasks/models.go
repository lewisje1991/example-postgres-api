package tasks

import (
	"fmt"
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

func (r *CreateTaskRequest) Validate() error {
	if r.Title == "" {
		return fmt.Errorf("title is required")
	}
	if r.Content == "" {
		return fmt.Errorf("content is required")
	}
	if r.Status == "" {
		return fmt.Errorf("status is required")
	}
	if r.Tags == "" {
		return fmt.Errorf("tags are required")
	}
	return nil
}

type UpdateTaskRequest struct {
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
	Status  string `json:"status,omitempty"`
	Tags    string `json:"tags,omitempty"`
}
