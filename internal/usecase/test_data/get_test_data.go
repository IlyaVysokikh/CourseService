package test_data

import (
	dto "CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/services"
	"CourseService/internal/usecase/shared"
	"context"
	"github.com/google/uuid"
)

type GetTestDataUseCaseImpl struct {
	service services.TestDataService
}

func NewGetTestDataUseCase(service services.TestDataService) shared.GetTestDataUseCase {
	return &GetTestDataUseCaseImpl{
		service: service,
	}
}

func (u *GetTestDataUseCaseImpl) Handle(ctx context.Context, id uuid.UUID) (dto.TestDataResponse, error) {
	return u.service.Get(ctx, id)
}
