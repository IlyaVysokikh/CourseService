package test_data

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/services"
	"context"
	"github.com/google/uuid"
)

type GetAllTestDataUseCaseImpl struct {
	service services.TestDataService
}

func NewGetAllTestDataUseCase(service services.TestDataService) *GetAllTestDataUseCaseImpl {
	return &GetAllTestDataUseCaseImpl{
		service: service,
	}
}

func (u *GetAllTestDataUseCaseImpl) Handle(ctx context.Context, taskId uuid.UUID) ([]dto.TestDataResponse, error) {
	return u.service.GetAll(ctx, taskId)
}
