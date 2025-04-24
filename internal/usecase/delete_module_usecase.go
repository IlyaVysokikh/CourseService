package usecase

import (
	"CourseService/internal/services"
	"context"
	"github.com/google/uuid"
)

type DeleteModuleUseCaseImpl struct {
	moduleService services.ModuleService
}

func NewDeleteModuleUseCase(moduleService services.ModuleService) DeleteModuleUseCase {
	return &DeleteModuleUseCaseImpl{
		moduleService: moduleService,
	}
}

func (uc *DeleteModuleUseCaseImpl) Handle(ctx context.Context, id uuid.UUID) error {
	return uc.moduleService.DeleteModule(ctx, id)
}
