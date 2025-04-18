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

	CreateCourseUsecase interface {
		Handle(ctx context.Context, course *dto.CreateCourse) (dto.CreateCourseResponse, error)
	}

	CloneCourseUsecase interface {
		Handle(ctx context.Context, course *dto.CloneCourseRequest) (dto.CreateCourseResponse, error)
	}

	CreateModulesUsecase interface{
		Handle(ctx context.Context, courseID uuid.UUID, module *dto.CreateModulesRequest) error
	}

	Usecase struct {
		GetAllCourseUsecase GetAllCourseUsecase
		GetCourseUsecase GetCourseUsecase
		CreateCourseUsecase CreateCourseUsecase
		CloneCourseUsecase CloneCourseUsecase
		CreateModulesUsecase CreateModulesUsecase
	}
)

func NewUsecase(services *services.Service) *Usecase {
	return &Usecase{
		GetAllCourseUsecase: NewGetAllCourseUsecase(services.CourseService),
		GetCourseUsecase: NewGetCourseUsecase(services.CourseService, services.ModuleService, services.TaskService),
		CreateCourseUsecase: NewCreateCourseUsecase(services.CourseService),
		CloneCourseUsecase: NewCloneCourseUsecase(services.CourseService),
		CreateModulesUsecase: NewCreateModuleUsecase(services.ModuleService),
	}
}
