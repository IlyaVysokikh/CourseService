package services

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/repositories"
	"context"
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
