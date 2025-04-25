package module

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/services"
	"CourseService/internal/usecase/shared"
	ierrors "CourseService/pkg/errors"

	"context"
	"errors"
	"log/slog"

	"github.com/google/uuid"
)

type GetModuleUseCaseImpl struct {
	moduleService           services.ModuleService
	taskService             services.TaskService
	moduleAttachmentService services.ModuleAttachmentService
}

func NewGetModuleUseCase(
	moduleService services.ModuleService,
	taskService services.TaskService,
	moduleAttachmentService services.ModuleAttachmentService) shared.GetModuleUseCase {
	return GetModuleUseCaseImpl{
		moduleService:           moduleService,
		taskService:             taskService,
		moduleAttachmentService: moduleAttachmentService,
	}
}

func (gmu GetModuleUseCaseImpl) Handle(ctx context.Context, moduleID uuid.UUID) (dto.GetModuleResponse, error) {
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

	attachments, err := gmu.moduleAttachmentService.GetAllByModule(ctx, moduleID)
	if err != nil {
		slog.Error("Error getting module attachments", "error", err)
		return dto.GetModuleResponse{}, err
	}

	return dto.GetModuleResponse{
		Module:     module,
		Tasks:      tasks,
		Attachment: attachments,
	}, nil
}
