package usecase

import (
	"CourseService/internal/services"
	"CourseService/internal/interfaces/rest/dto"
	
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
		return err
	}

	return nil
}