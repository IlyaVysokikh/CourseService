package test_data

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/services"
	"CourseService/internal/usecase/shared"
	"context"
	"github.com/google/uuid"
)

type CreateTestDataUseCaseImpl struct {
	service services.TestDataService
}

func NewCreateTestDataUseCase(service services.TestDataService) shared.CreateTestDataUseCase {
	return &CreateTestDataUseCaseImpl{
		service: service,
	}
}

func (u *CreateTestDataUseCaseImpl) Handle(ctx context.Context, request dto.CreateTestDataRequest) (uuid.UUID, error) {
	return u.service.Create(ctx, request)
}
