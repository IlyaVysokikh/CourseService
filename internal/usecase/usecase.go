package usecase

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/services"
	"context"

	"github.com/google/uuid"
)

type (
	GetAllCourseUsecase interface {
		Handle(ctx context.Context, filter *dto.CourseFilter) ([]dto.CourseList, error)
	}

	GetCourseUsecase interface {
		Handle(ctx context.Context, id uuid.UUID) (*dto.Course, error)
	}

	Usecase struct {
		GetAllCourseUsecase GetAllCourseUsecase
		GetCourseUsecase GetCourseUsecase
	}
)

func NewUsecase(services *services.Service) *Usecase {
	return &Usecase{
		GetAllCourseUsecase: NewGetAllCourseUsecase(services.CourseService),
		GetCourseUsecase: NewGetCourseUsecase(services.CourseService, services.ModuleService, services.TaskService),
	}
}
