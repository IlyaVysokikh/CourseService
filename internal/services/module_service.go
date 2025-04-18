package services

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/repositories"
	"context"
	"log/slog"

	"github.com/google/uuid"
)

type ModuleServiceImpl struct {
	repo repositories.ModuleRepository
}

func NewModuleService(repository repositories.ModuleRepository) ModuleService {
	return ModuleServiceImpl{
		repo: repository,
	}
}

func (ms ModuleServiceImpl) GetModulesByCourse(ctx context.Context, courseID uuid.UUID) ([]dto.ModuleList, error) {
	modules, err := ms.repo.GetModulesByCourse(courseID)
	if err != nil {
		return nil, err
	}

	var moduleList []dto.ModuleList
	for _, module := range modules {
		moduleList = append(moduleList, dto.ModuleList{
			Id:   module.ID,
			Name: module.Name,
			DateStart: module.DateStart,
			SequenceNumber: module.SequenceNumber,
		})
	}

	return moduleList, nil
}

func (ms ModuleServiceImpl) CreateModules(ctx context.Context, courseID uuid.UUID, modules dto.CreateModulesRequest) error {
	var newModules []dto.CreateModule
	var incomingModules []dto.CreateModule
	
	for i := 0; i < len(modules.Modules); i++ {
		if modules.Modules[i].Id == nil || *modules.Modules[i].Id == uuid.Nil {
			incomingModules = append(incomingModules, modules.Modules[i]) 
		} else {
			newModules = append(newModules, modules.Modules[i]) 
		}
	}

	err := ms.repo.UpdateModules(courseID, newModules)
	if err != nil {
		slog.Error("Error updating modules", "error", err)
		return err
	}

	err = ms.repo.CreateModules(courseID, incomingModules)
	if err != nil {
		slog.Error("Error creating modules", "error", err)
	}

	return nil
}
