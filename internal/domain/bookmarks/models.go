package bookmarks

import (
	"errors"
	"net/url"
	"time"

	"github.com/google/uuid"
)

type Bookmark struct {
	ID          uuid.UUID `json:"id"`
	URL         string    `json:"url"`
	Description string    `json:"description"`
	Tags        []string  `json:"tags"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (b *Bookmark) Validate() error {
	_, err := url.ParseRequestURI(b.URL)
	if err != nil {
		return errors.New("invalid url")
	}
	return nil
}
