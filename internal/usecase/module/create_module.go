package module

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/services"
	"CourseService/internal/usecase/shared"
	ierrors "CourseService/pkg/errors"
	"errors"
	"log/slog"

	"context"
)

type CreateModulesUseCaseImpl struct {
	moduleService services.ModuleService
}

func NewCreateModuleUseCase(moduleService services.ModuleService) shared.CreateModulesUseCase {
	return CreateModulesUseCaseImpl{
		moduleService: moduleService,
	}
}

func (cmu CreateModulesUseCaseImpl) Handle(ctx context.Context, module *dto.CreateModulesRequest) error {
	err := cmu.moduleService.CreateModules(ctx, *module)
	if err != nil {
		if errors.Is(err, ierrors.ErrInternal) {
			slog.Error("Error creating modules", "error", err)
			return ierrors.New(ierrors.ErrInternal, "failed to create modules", err)
		}

		slog.Error("Unexpected error creating modules", "error", err)
		return err
	}

	return nil
}
