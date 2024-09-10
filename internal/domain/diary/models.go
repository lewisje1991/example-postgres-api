package diary

import (
	"time"

	"github.com/google/uuid"
)

type Diary struct {
	ID    uuid.UUID
	Day   time.Time
	Tasks []Task
}

type Task struct {
	ID     string
	Name   string
	Status string
}
