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

	CreateModulesUsecase interface {
		Handle(ctx context.Context, courseID uuid.UUID, module *dto.CreateModulesRequest) error
	}

	GetModuleUsecase interface {
		Handle(ctx context.Context, moduleID uuid.UUID) (dto.GetModuleResponse, error)
	}

	GetTaskUseCase interface {
		Handle(ctx context.Context, taskId uuid.UUID) (*dto.TaskExtended, error)
	}

	DeleteCourseUsecase interface {
		Handle(ctx context.Context, id uuid.UUID) error
	}

	UpdateCourseUsecase interface {
		Handle(ctx context.Context, id uuid.UUID, request dto.UpdateCourseRequest) error
	}

	DeleteModuleUseCase interface {
		Handle(ctx context.Context, id uuid.UUID) error
	}

	Usecase struct {
		GetAllCourseUseCase  GetAllCourseUsecase
		GetCourseUseCase     GetCourseUsecase
		CreateCourseUseCase  CreateCourseUsecase
		CloneCourseUseCase   CloneCourseUsecase
		CreateModulesUseCase CreateModulesUsecase
		GetModuleUseCase     GetModuleUsecase
		GetTaskUseCase       GetTaskUseCase
		DeleteCourseUseCase  DeleteCourseUsecase
		UpdateCourseUseCase  UpdateCourseUsecase
		DeleteModuleUseCase  DeleteModuleUseCase
	}
)

func NewUsecase(services *services.Service) *Usecase {
	return &Usecase{
		GetAllCourseUseCase:  NewGetAllCourseUsecase(services.CourseService),
		GetCourseUseCase:     NewGetCourseUsecase(services.CourseService, services.ModuleService, services.TaskService),
		CreateCourseUseCase:  NewCreateCourseUsecase(services.CourseService),
		CloneCourseUseCase:   NewCloneCourseUsecase(services.CourseService),
		CreateModulesUseCase: NewCreateModuleUsecase(services.ModuleService),
		GetModuleUseCase:     NewGetModuleUsecase(services.ModuleService, services.TaskService, services.ModuleAttachmentService),
		GetTaskUseCase:       NewGetTaskUseCase(services.TaskService),
		DeleteCourseUseCase:  NewDeleteCourseUseCase(services.CourseService),
		UpdateCourseUseCase:  NewUpdateCourseUsecase(services.CourseService),
		DeleteModuleUseCase:  NewDeleteModuleUseCase(services.ModuleService),
	}
}
