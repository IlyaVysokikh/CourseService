package test_data

import (
	"CourseService/internal/services"
	"CourseService/internal/usecase/shared"
	"context"
	"github.com/google/uuid"
)

type DeleteTestDataUseCaseImpl struct {
	service services.TestDataService
}

func NewDeleteTestDataUseCase(service services.TestDataService) shared.DeleteTestDataUseCase {
	return &DeleteTestDataUseCaseImpl{
		service: service,
	}
}

func (u *DeleteTestDataUseCaseImpl) Handle(ctx context.Context, id uuid.UUID) error {
	return u.service.Delete(ctx, id)
}
