package services

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/repositories"
	"CourseService/internal/repositories/models"
	"context"

	"github.com/google/uuid"
)

type (
	CourseService interface {
		GetAllCourses(ctx context.Context, filter *dto.CourseFilter) ([]dto.CourseList, error)
		GetCourse(ctx context.Context, id uuid.UUID) (*dto.Course, error)
		CreateCourse(ctx context.Context, course *dto.CreateCourse) (uuid.UUID, error)
		CloneCourse(ctx context.Context, course *dto.CloneCourseRequest) (uuid.UUID, error)
		DeleteCourse(ctx context.Context, id uuid.UUID) error
		UpdateCourse(ctx context.Context, id uuid.UUID, request dto.UpdateCourseRequest) error
	}

	ModuleService interface {
		GetModulesByCourse(ctx context.Context, courseID uuid.UUID) ([]dto.ModuleList, error)
		CreateModules(ctx context.Context, modules dto.CreateModulesRequest) error
		GetModule(ctx context.Context, moduleID uuid.UUID) (dto.GetModule, error)
		DeleteModule(ctx context.Context, id uuid.UUID) error
	}

	TaskService interface {
		GetTaskCount(ctx context.Context, moduleId uuid.UUID) (int, error)
		GetTasksByModule(ctx context.Context, moduleId uuid.UUID) ([]dto.Task, error)
		GetTask(ctx context.Context, taskId uuid.UUID) (*dto.TaskExtended, error)
		DeleteTask(ctx context.Context, id uuid.UUID) error
	}

	ModuleAttachmentService interface {
		GetAllByModule(ctx context.Context, moduleId uuid.UUID) ([]*models.ModuleAttachment, error)
		CreateAttachment(ctx context.Context, moduleId uuid.UUID, data dto.CreateModuleAttachmentRequest) (*models.ModuleAttachment, error)
	}

	TestDataService interface {
		GetAll(ctx context.Context, taskId uuid.UUID) ([]dto.TestDataResponse, error)
		Get(ctx context.Context, id uuid.UUID) (dto.TestDataResponse, error)
		Create(ctx context.Context, request dto.CreateTestDataRequest) (uuid.UUID, error)
		Update(ctx context.Context, id uuid.UUID, request dto.UpdateTestDataRequest) error
		Delete(ctx context.Context, id uuid.UUID) error
	}

	Service struct {
		CourseService           CourseService
		ModuleService           ModuleService
		TaskService             TaskService
		ModuleAttachmentService ModuleAttachmentService
		TestDataService         TestDataService
	}
)

func NewService(repos *repositories.Repository) *Service {
	return &Service{
		CourseService:           NewCourseServiceImpl(repos.CourseRepository),
		ModuleService:           NewModuleService(repos.ModuleRepository),
		TaskService:             NewTaskService(repos.TaskRepository),
		ModuleAttachmentService: NewModuleAttachmentService(repos.ModuleAttachmentRepository),
		TestDataService:         NewTestDataService(repos.TestDataRepository),
	}
}
