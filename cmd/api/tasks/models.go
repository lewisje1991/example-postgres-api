package tasks

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	domainTasks "github.com/lewisje1991/code-bookmarks/internal/tasks"
)

type CreateTaskRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  string `json:"status"`
	Tags    string `json:"tags"`
}

type TaskResponse struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Status    string    `json:"status"`
	Tags      string    `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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

func (r *CreateTaskRequest) ToDomain() domainTasks.Task {
	return domainTasks.Task{
		Title:   r.Title,
		Content: r.Content,
		Status:  r.Status,
		Tags:    r.Tags,
	}
}

func TaskResponseFromDomain(task *domainTasks.Task) TaskResponse {
	return TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		Content:   task.Content,
		Status:    task.Status,
		Tags:      task.Tags,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
}

type UpdateTaskRequest struct {
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
	Status  string `json:"status,omitempty"`
	Tags    string `json:"tags,omitempty"`
}
