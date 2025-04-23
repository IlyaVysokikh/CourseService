package usecase

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/services"
	ierrors "CourseService/pkg/errors"

	"context"

	"github.com/google/uuid"
)

type GetTaskUseCaseImpl struct {
	taskService services.TaskService
}

func NewGetTaskUseCase(taskService services.TaskService) *GetTaskUseCaseImpl {
	return &GetTaskUseCaseImpl{
		taskService: taskService,
	}
}

func (g *GetTaskUseCaseImpl) Handle(ctx context.Context, taskId uuid.UUID) (*dto.TaskExtended, error) {
	task, err := g.taskService.GetTask(ctx, taskId)
	if err != nil {
		if err == ierrors.ErrNotFound {
			return nil, ierrors.New(ierrors.ErrNotFound, "task not found", err)
		}

		if err == ierrors.ErrInternal {
			return nil, ierrors.New(ierrors.ErrInternal, "failed to get task", err)
		}
		return nil, err
	}

	return task, nil
}