package bookmarks

import (
	"time"

	"github.com/google/uuid"
)

type Bookmark struct {
	ID          uuid.UUID
	URL         string
	Description string
	Tags        []string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
