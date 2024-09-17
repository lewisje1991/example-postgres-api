package diary

import (
	"time"

	"github.com/google/uuid"
	"github.com/lewisje1991/code-bookmarks/internal/domain/tasks"
)

type Diary struct {
	ID  uuid.UUID
	Day time.Time
}

type DiaryWithTasks struct {
	Diary
	Tasks []tasks.Task
}
