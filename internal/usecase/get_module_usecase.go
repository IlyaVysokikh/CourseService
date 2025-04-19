package usecase

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/services"
	ierrors "CourseService/pkg/errors"

	"context"
	"log/slog"
	"errors"

	"github.com/google/uuid"
)

type GetModuleUsecaseImpl struct {
	moduleService services.ModuleService
	taskService services.TaskService
}

func NewGetModuleUsecase(moduleService services.ModuleService, taskService services.TaskService) GetModuleUsecase {
	return GetModuleUsecaseImpl{
		moduleService: moduleService,
		taskService: taskService,
	}
}

func (gmu GetModuleUsecaseImpl) Handle(ctx context.Context, moduleID uuid.UUID) (dto.GetModuleResponse, error) {
	module, err := gmu.moduleService.GetModule(ctx, moduleID)
	if err != nil {
		if errors.Is(err, ierrors.ErrNotFound) {
			slog.Warn("module not found", "moduleID", moduleID)
			return dto.GetModuleResponse{}, ierrors.New(ierrors.ErrNotFound, "module not found", err)
		}
		slog.Error("Error getting module", "moduleID", moduleID, "error", err)
		return dto.GetModuleResponse{}, ierrors.New(ierrors.ErrInternal, "failed to get module", err)
	}

	tasks, err := gmu.taskService.GetTasksByModule(ctx, moduleID)
	if err != nil {
		slog.Error("Error getting tasks by module", "error", err)
		return dto.GetModuleResponse{}, err
	}

	return dto.GetModuleResponse{
		Module: module,
		Tasks:  tasks,
	}, nil
}