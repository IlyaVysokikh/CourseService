package services

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/repositories"
	"context"
	"log/slog"
	"errors"
	ierrors "CourseService/pkg/errors"

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
		slog.Error("error getting modules by course", slog.Any("err", err))
		
		if errors.Is(err, ierrors.ErrNotFound) {
			return nil, ierrors.New(ierrors.ErrNotFound, "modules not found", err)
		}

		return nil, ierrors.New(ierrors.ErrInternal, "failed to get modules", err)
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

	if err := ms.repo.UpdateModules(courseID, newModules); err != nil {
		slog.Error("error updating modules", slog.Any("err", err))
		return ierrors.New(ierrors.ErrInternal, "failed to update modules", err)
	}

	if err := ms.repo.CreateModules(courseID, incomingModules); err != nil {
		slog.Error("error creating modules", slog.Any("err", err))
		return ierrors.New(ierrors.ErrInternal, "failed to create modules", err)
	}


	return nil
}

func (ms ModuleServiceImpl) GetModule(ctx context.Context, moduleID uuid.UUID) (dto.GetModule, error) {
	module, err := ms.repo.GetModule(moduleID)
	if err != nil {
		slog.Error("error getting module", slog.Any("err", err))

		if errors.Is(err, ierrors.ErrNotFound) {
			return dto.GetModule{}, ierrors.New(ierrors.ErrNotFound, "module not found", err)
		}

		return dto.GetModule{}, ierrors.New(ierrors.ErrInternal, "failed to get module", err)
	}

	return dto.GetModule{
		Id:             module.ID,
		Name:           module.Name,
	}, nil
}