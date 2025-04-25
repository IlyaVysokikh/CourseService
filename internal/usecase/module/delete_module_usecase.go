package module

import (
	"CourseService/internal/services"
	"CourseService/internal/usecase/shared"
	"context"
	"github.com/google/uuid"
)

type DeleteModuleUseCaseImpl struct {
	moduleService services.ModuleService
}

func NewDeleteModuleUseCase(moduleService services.ModuleService) shared.DeleteModuleUseCase {
	return &DeleteModuleUseCaseImpl{
		moduleService: moduleService,
	}
}

func (uc *DeleteModuleUseCaseImpl) Handle(ctx context.Context, id uuid.UUID) error {
	return uc.moduleService.DeleteModule(ctx, id)
}
