package services

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/repositories"
	ierrors "CourseService/pkg/errors"
	"context"
	"errors"
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

func (t *TaskServiceImpl) GetTaskCount(ctx context.Context, moduleId uuid.UUID) (int, error) {
	tasks, err := t.GetTasksByModule(ctx, moduleId)
	if err != nil {
		if errors.Is(err, ierrors.ErrInternal) {
			return 0, ierrors.New(ierrors.ErrInternal, "failed to get tasks", err)
		}

		return 0, err
	}

	return len(tasks), nil
}

func (t *TaskServiceImpl) GetTasksByModule(ctx context.Context, moduleId uuid.UUID) ([]dto.Task, error) {
	tasks, err := t.repo.GetTasksByModule(moduleId)
	if err != nil {
		if errors.Is(err, ierrors.ErrInternal) {
			return nil, ierrors.New(ierrors.ErrInternal, "failed to get tasks", err)
		}

		return nil, err
	}

	var taskList []dto.Task
	for _, task := range tasks {
		taskList = append(taskList, dto.Task{
			Id:   task.Id,
			Name: task.Name,
		})
	}

	return taskList, nil
}

func (t *TaskServiceImpl) GetTask(ctx context.Context, taskId uuid.UUID) (*dto.TaskExtended, error) {
	task, err := t.repo.GetTask(taskId)
	if err != nil {
		if errors.Is(err, ierrors.ErrNotFound) {
			slog.Warn("Task not found", "taskId", taskId)
			return nil, ierrors.New(ierrors.ErrNotFound, "task not found", err)
		}

		if errors.Is(err, ierrors.ErrInternal) {
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

func (t *TaskServiceImpl) DeleteTask(ctx context.Context, id uuid.UUID) error {
	return t.repo.DeleteTask(id)
}

func (t *TaskServiceImpl) CreateTask(ctx context.Context, request dto.CreateTaskRequest) (uuid.UUID, error) {
	if len(request.Name) > 255 {
		slog.Error("Name is too long", "name", request.Name)
		return uuid.Nil, ierrors.ErrInvalidInput
	}

	if request.Text == "" {
		slog.Error("Text can not be blank", "text", request.Text)
		return uuid.Nil, ierrors.ErrInvalidInput
	}

	tasks, err := t.repo.GetTasksByModule(request.ModuleId)
	if err != nil {
		if errors.Is(err, ierrors.ErrNotFound) {
			slog.Error("Failed to get tasks", "moduleId", request.ModuleId, "error", err)
		}

		if errors.Is(err, ierrors.ErrInternal) {
			slog.Error("Failed to get tasks", "moduleId", request.ModuleId, "error", err)
			return uuid.Nil, ierrors.ErrInternal
		}
	}

	for _, task := range tasks {
		if task.SequenceNumber == request.SequenceNumber {
			slog.Warn("Task with that sequence number already exists", "taskId", task.Id)
			return uuid.Nil, ierrors.ErrInvalidInput
		}
	}

	taskId, err := t.repo.Create(ctx, request)
	if err != nil {
		if errors.Is(err, ierrors.ErrInternal) {
			slog.Error("Failed to create task", "error", err)
			return uuid.Nil, ierrors.ErrInternal
		}
	}

	return taskId, nil
}
