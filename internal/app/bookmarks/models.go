package bookmarks

import (
	"fmt"

	"github.com/lewisje1991/code-bookmarks/internal/domain/bookmarks"
)

type Request struct {
	URL         string   `json:"url"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

func (r *Request) ToBookmark() *bookmarks.Bookmark {
	return &bookmarks.Bookmark{
		URL:         r.URL,
		Description: r.Description,
		Tags:        r.Tags,
	}
}

func (r *Request) Validate() error {
	if r.URL == "" {
		return fmt.Errorf("url is required")
	}

	if r.Description == "" {
		return fmt.Errorf("description is required")
	}
	return nil

}

type Response struct {
	ID          string   `json:"id,omitempty"`
	URL         string   `json:"url,omitempty"`
	Description string   `json:"description,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	CreatedAt   string   `json:"createdAt,omitempty"`
	UpdatedAt   string   `json:"updatedAt,omitempty"`
}
