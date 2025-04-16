package usecase

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/services"
	"context"
)

type (
	GetAllCourseUsecase interface {
		Handle(ctx context.Context, filter *dto.CourseFilter) ([]dto.CourseList, error)
	}

	Usecase struct {
		GetAllCourseUsecase GetAllCourseUsecase
	}
)

func NewUsecase(services *services.Service) *Usecase {
	return &Usecase{
		GetAllCourseUsecase: NewGetAllCourseUsecase(services.CourseService),
	}
}