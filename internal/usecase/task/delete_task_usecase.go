package task

import (
	"CourseService/internal/services"
	"CourseService/internal/usecase/shared"
	ierrors "CourseService/pkg/errors"
	"context"
	"errors"
	"github.com/google/uuid"
	"log/slog"
)

type DeleteTaskUseCaseImpl struct {
	taskService services.TaskService
}

func NewDeleteTaskUseCase(taskService services.TaskService) shared.DeleteTaskUseCase {
	return &DeleteTaskUseCaseImpl{
		taskService: taskService,
	}
}

func (u *DeleteTaskUseCaseImpl) Handle(ctx context.Context, id uuid.UUID) error {
	if err := u.taskService.DeleteTask(ctx, id); err != nil {
		if errors.Is(err, ierrors.ErrNotFound) {
			slog.Error("task not found", "err", err, "taskId", id)
			return ierrors.ErrNotFound
		}

		slog.Error("delete task failed", "err", err, "taskId", id)
		return ierrors.ErrInternal
	}

	return nil
}
