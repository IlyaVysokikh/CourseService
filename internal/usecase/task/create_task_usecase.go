package task

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/services"
	"CourseService/internal/usecase/shared"
	ierrors "CourseService/pkg/errors"
	"context"
	"errors"
	"log/slog"
)

type CreateTaskUseCaseImpl struct {
	taskService   services.TaskService
	moduleService services.ModuleService
}

func NewCreateTaskUseCase(taskService services.TaskService, moduleService services.ModuleService) shared.CreateTaskUseCase {
	return &CreateTaskUseCaseImpl{
		taskService:   taskService,
		moduleService: moduleService,
	}
}

func (u *CreateTaskUseCaseImpl) Handle(ctx context.Context, request dto.CreateTaskRequest) (*dto.CreateTaskResponse, error) {

	exists, err := u.moduleService.ModuleExists(ctx, request.ModuleId)
	if err != nil {
		slog.Error("Occurred error while checking module existence", err)
		return &dto.CreateTaskResponse{}, ierrors.ErrInternal
	}

	if !exists {
		slog.Error("Can not find module for create task", "moduleId", request.ModuleId)
		return &dto.CreateTaskResponse{}, ierrors.ErrNotFound
	}

	taskId, err := u.taskService.CreateTask(ctx, request)
	if err != nil {
		if errors.Is(err, ierrors.ErrInvalidInput) {
			slog.Error("Invalid input for create task", err)
			return &dto.CreateTaskResponse{}, ierrors.ErrInvalidInput
		}

		slog.Error("Occurred internal error", "err", err)
		return &dto.CreateTaskResponse{}, ierrors.ErrInternal
	}

	return &dto.CreateTaskResponse{
		TaskId: taskId,
	}, nil
}
