package services

import (
	"CourseService/internal/repositories"
	"context"

	"github.com/google/uuid"
)

type TaskServiceImpl struct {
	repo repositories.TaskRepository
}

func NewTaskService(repo repositories.TaskRepository) TaskService {
	return &TaskServiceImpl{
		repo: repo,
	}
}

func (t *TaskServiceImpl) GetTaskCount(ctx  context.Context, moduleId uuid.UUID) (int, error) {
	tasks, err := t.repo.GetTasks(moduleId)
	if err != nil {
		return 0, err
	}

	return len(tasks), nil
}