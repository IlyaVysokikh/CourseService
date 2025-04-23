package services

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/repositories"
	ierrors "CourseService/pkg/errors"
	"context"
	"log/slog"

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

func (t *TaskServiceImpl) GetTask(ctx context.Context, taskId uuid.UUID) (*dto.TaskExtended, error) {
	task, err := t.repo.GetTask(taskId)
	if err != nil {
		if err == ierrors.ErrNotFound {
			slog.Warn("Task not found", "taskId", taskId)
			return nil, ierrors.New(ierrors.ErrNotFound, "task not found", err)
		}

		if err == ierrors.ErrInternal {
			slog.Error("Failed to get task", "taskId", taskId, "error", err)
			return nil, ierrors.New(ierrors.ErrInternal, "failed to get task", err)
		}
		return nil, err
	}

	return &dto.TaskExtended{
		Id:               task.Id,
		Name:             task.Name,
		Text:             task.Text,
		Language:         task.Language,
		InitialCode:      task.InitialCode,
		MemoryLimit:      task.MemoryLimit,
		ExecutionTimeout: task.ExecutionTimeout,
	}, nil

}