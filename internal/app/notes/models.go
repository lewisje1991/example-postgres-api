package notes

import "fmt"

type Request struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (r *Request) Validate() error {
	if r.Title == "" {
		return fmt.Errorf("title is required")
	}

	if r.Content == "" {
		return fmt.Errorf("content is required")
	}
	return nil
}

type Response struct {
	Title     string   `json:"title,omitempty"`
	Content   string   `json:"content,omitempty"`
	Tags      []string `json:"tags,omitempty"`
	CreatedAt string   `json:"createdAt,omitempty"`
	UpdatedAt string   `json:"updatedAt,omitempty"`
}
