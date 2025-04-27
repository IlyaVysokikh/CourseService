package test_data

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/services"
	"CourseService/internal/usecase/shared"
	"context"
	"github.com/google/uuid"
)

type UpdateTestDataUseCaseImpl struct {
	service services.TestDataService
}

func NewUpdateTestDataUseCase(service services.TestDataService) shared.UpdateTestDataUseCase {
	return &UpdateTestDataUseCaseImpl{
		service: service,
	}
}

func (u *UpdateTestDataUseCaseImpl) Handle(ctx context.Context, id uuid.UUID, request dto.UpdateTestDataRequest) error {
	return u.service.Update(ctx, id, request)
}
