package shared

import (
	"CourseService/internal/interfaces/rest/dto"
	"context"

	"github.com/google/uuid"
)

type (
	GetAllCourseUseCase interface {
		Handle(ctx context.Context, filter *dto.CourseFilter) ([]dto.CourseList, error)
	}

	GetCourseUseCase interface {
		Handle(ctx context.Context, id uuid.UUID) (*dto.Course, error)
	}

	CreateCourseUseCase interface {
		Handle(ctx context.Context, course *dto.CreateCourse) (dto.CreateCourseResponse, error)
	}

	CloneCourseUseCase interface {
		Handle(ctx context.Context, course *dto.CloneCourseRequest) (dto.CreateCourseResponse, error)
	}

	CreateModulesUseCase interface {
		Handle(ctx context.Context, module *dto.CreateModulesRequest) error
	}

	GetModuleUseCase interface {
		Handle(ctx context.Context, moduleID uuid.UUID) (dto.GetModuleResponse, error)
	}

	GetTaskUseCase interface {
		Handle(ctx context.Context, taskId uuid.UUID) (*dto.TaskExtended, error)
	}

	DeleteCourseUseCase interface {
		Handle(ctx context.Context, id uuid.UUID) error
	}

	UpdateCourseUseCase interface {
		Handle(ctx context.Context, id uuid.UUID, request dto.UpdateCourseRequest) error
	}

	DeleteModuleUseCase interface {
		Handle(ctx context.Context, id uuid.UUID) error
	}

	DeleteTaskUseCase interface {
		Handle(ctx context.Context, id uuid.UUID) error
	}

	CreateModuleAttachmentUseCase interface {
		Handle(ctx context.Context, moduleId uuid.UUID, request dto.CreateModuleAttachmentRequest) (dto.CreateModuleAttachmentResponse, error)
	}

	GetAllTestDataUseCase interface {
		Handle(ctx context.Context, taskId uuid.UUID) ([]dto.TestDataResponse, error)
	}

	GetTestDataUseCase interface {
		Handle(ctx context.Context, id uuid.UUID) (dto.TestDataResponse, error)
	}

	CreateTestDataUseCase interface {
		Handle(ctx context.Context, request dto.CreateTestDataRequest) (uuid.UUID, error)
	}

	UpdateTestDataUseCase interface {
		Handle(ctx context.Context, id uuid.UUID, request dto.UpdateTestDataRequest) error
	}

	DeleteTestDataUseCase interface {
		Handle(ctx context.Context, id uuid.UUID) error
	}
)
