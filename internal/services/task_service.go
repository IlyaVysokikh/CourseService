package services

import (
	"CourseService/internal/interfaces/rest/dto"
	ierrors "CourseService/pkg/errors"
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
	tasks, err := t.GetTasksByModule(ctx, moduleId)
	if err != nil {
		if err == ierrors.ErrInternal {
			return 0, ierrors.New(ierrors.ErrInternal, "failed to get tasks", err)
		}
		
		return 0, err
	}

	return len(tasks), nil
}

func (t *TaskServiceImpl) GetTasksByModule(ctx context.Context, moduleId uuid.UUID) ([]dto.Task, error) {
	tasks, err := t.repo.GetTasksByModule(moduleId)
	if err != nil {
		if err == ierrors.ErrInternal { 
			return nil, ierrors.New(ierrors.ErrInternal, "failed to get tasks", err)
		}

		return nil, err
	}

	var taskList []dto.Task
	for _, task := range tasks {
		taskList = append(taskList, dto.Task{
			Id:          task.Id,
			Name:        task.Name,
		})
	}

	return taskList, nil
}