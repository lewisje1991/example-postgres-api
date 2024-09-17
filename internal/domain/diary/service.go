package diary

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lewisje1991/code-bookmarks/internal/domain/tasks"
)

type Service struct {
	taskService *tasks.Service
	store       *Store
}

func NewService(s *Store, taskService *tasks.Service) *Service {
	return &Service{
		store:       s,
		taskService: taskService,
	}
}

func (s *Service) NewDiaryEntry(ctx context.Context, today time.Time) (DiaryWithTasks, error) {
	// if diary entry for today exists, return it


	existingEntry, err := s.store.GetDiaryByDay(ctx, today)
	if err != nil {
		return DiaryWithTasks{}, fmt.Errorf("failed to check for existing diary: %w", err)
	}

	existingEntryTasks, err := s.taskService.GetTasksForDiary(ctx, existingEntry.ID)
	if err != nil {
		return DiaryWithTasks{}, fmt.Errorf("failed to get tasks for diary: %w", err)
	}

	if existingEntry.ID != uuid.Nil {
		return DiaryWithTasks{
			Diary: Diary{
				ID:  existingEntry.ID,
				Day: existingEntry.Day,
			},
			Tasks: existingEntryTasks,
		}, nil
	}

	// if diary entry for today does not exist, create it
	entity := Diary{
		ID:    uuid.New(),
		Day:   time.Now(),
	}

	dbEntity, err := s.store.CreateDiary(ctx, entity)
	if err != nil {
		return DiaryWithTasks{}, fmt.Errorf("failed to insert new diary entry: %w", err)
	}

	newEntry := DiaryWithTasks{
		Diary: Diary{
			ID:  dbEntity.ID,
			Day: dbEntity.Day,
		},
	}

	// get from previous day diary id
	previousDay := today.AddDate(0, 0, -1)
	previousDayEntry, err := s.store.GetDiaryByDay(ctx, previousDay)
	if err != nil {
		return DiaryWithTasks{}, fmt.Errorf("failed to retrieve previous day diary entry: %w", err)
	}

	// get uncompleted tasks from previous day
	previousDayTasks, err := s.taskService.GetTasksForDiary(ctx, previousDayEntry.ID)
	if err != nil {
		return DiaryWithTasks{}, fmt.Errorf("failed to retrieve tasks for previous day: %w", err)
	}

	previousDayTasksIDs := make([]uuid.UUID, len(previousDayTasks))
	for i, task := range previousDayTasks {
		previousDayTasksIDs[i] = task.ID
	}

	// add uncompleted tasks to new entry in db
	err = s.AddTasksToDiary(ctx, newEntry.ID, previousDayTasksIDs)
	if err != nil {
		return DiaryWithTasks{}, fmt.Errorf("failed to create tasks for new diary entry: %w", err)
	}


	return newEntry, nil
}

func (s *Service) GetDiaryEntryWithTasks(ctx context.Context, day time.Time) (DiaryWithTasks, error) {
	diaryEntry, err := s.store.GetDiaryByDay(ctx, day)
	if err != nil {
		return DiaryWithTasks{}, fmt.Errorf("failed to retrieve diary entry: %w", err)
	}

	if diaryEntry.ID == uuid.Nil {
		return DiaryWithTasks{}, fmt.Errorf("no diary entry found for the specified day")
	}

	tasks, err := s.taskService.GetTasksForDiary(ctx, diaryEntry.ID)
	if err != nil {
		return DiaryWithTasks{}, fmt.Errorf("failed to retrieve tasks: %w", err)
	}

	return DiaryWithTasks{Diary: diaryEntry, Tasks: tasks}, nil
}

func (s *Service) AddTasksToDiary(ctx context.Context, diaryID uuid.UUID, tasks []uuid.UUID) error {
	return nil
}