package task

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/services"
	"CourseService/internal/usecase/shared"
	ierrors "CourseService/pkg/errors"
	"errors"

	"context"

	"github.com/google/uuid"
)

type GetTaskUseCaseImpl struct {
	taskService services.TaskService
}

func NewGetTaskUseCase(taskService services.TaskService) shared.GetTaskUseCase {
	return &GetTaskUseCaseImpl{
		taskService: taskService,
	}
}

func (g *GetTaskUseCaseImpl) Handle(ctx context.Context, taskId uuid.UUID) (*dto.TaskExtended, error) {
	task, err := g.taskService.GetTask(ctx, taskId)
	if err != nil {
		if errors.Is(err, ierrors.ErrNotFound) {
			return nil, ierrors.New(ierrors.ErrNotFound, "task not found", err)
		}

		if errors.Is(err, ierrors.ErrInternal) {
			return nil, ierrors.New(ierrors.ErrInternal, "failed to get task", err)
		}
		return nil, err
	}

	return task, nil
}
