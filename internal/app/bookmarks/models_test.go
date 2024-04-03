package bookmarks_test

import (
	"testing"

	"github.com/lewisje1991/code-bookmarks/internal/app/bookmarks"
	"github.com/stretchr/testify/assert"
)

func TestRequest_Validate(t *testing.T) {
	t.Run("returns error when URL is empty", func(t *testing.T) {
		request := &bookmarks.Request{
			URL:         "",
			Description: "Sample description",
		}
		err := request.Validate()
		assert.EqualError(t, err, "url is required")
	})

	t.Run("returns error when description is empty", func(t *testing.T) {
		request := &bookmarks.Request{
			URL:         "https://example.com",
			Description: "",
		}
		err := request.Validate()
		assert.EqualError(t, err, "description is required")
	})

	t.Run("returns nil when URL and description are provided", func(t *testing.T) {
		request := &bookmarks.Request{
			URL:         "https://example.com",
			Description: "Sample description",
		}
		err := request.Validate()
		assert.NoError(t, err)
	})
}
