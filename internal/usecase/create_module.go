package usecase

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/services"
	ierrors "CourseService/pkg/errors"
	"log/slog"

	"github.com/google/uuid"

	"context"
)



type CreateModulesUsecaseImpl struct {
	moduleService services.ModuleService
}

func NewCreateModuleUsecase(moduleService services.ModuleService) CreateModulesUsecase {
	return CreateModulesUsecaseImpl{
		moduleService: moduleService,
	}
}

func (cmu CreateModulesUsecaseImpl) Handle(ctx context.Context, courseID uuid.UUID, module *dto.CreateModulesRequest) (error) {
	err := cmu.moduleService.CreateModules(ctx, courseID, *module)
	if err != nil {
		if err == ierrors.ErrInternal {
			slog.Error("Error creating modules", "error", err)
			return ierrors.New(ierrors.ErrInternal, "failed to create modules", err)
		}
		
		slog.Error("Unexpected error creating modules", "error", err)
		return err
	}

	return nil
}