package services

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/repositories"
	"context"

	"github.com/google/uuid"
)

type (
	CourseService interface{
		GetAllCourses(ctx context.Context, filter *dto.CourseFilter) ([]dto.CourseList, error)
		GetCourse(ctx context.Context, id uuid.UUID) (*dto.Course, error)
		CreateCourse(ctx context.Context, course *dto.CreateCourse) (uuid.UUID, error)
		CloneCourse(ctx context.Context, course *dto.CloneCourseRequest) (uuid.UUID, error)
	}

	ModuleService interface{
		GetModulesByCourse(ctx context.Context, courseID uuid.UUID) ([]dto.ModuleList, error)
		CreateModules(ctx context.Context, courseID uuid.UUID, modules dto.CreateModulesRequest) error
		GetModule(ctx context.Context, moduleID uuid.UUID) (dto.GetModule, error) 
	}

	TaskService interface {
		GetTaskCount(ctx context.Context, moduleId uuid.UUID) (int, error)
		GetTasksByModule(ctx context.Context, moduleId uuid.UUID) ([]dto.Task, error)
	}

	Service struct {
		CourseService CourseService
		ModuleService ModuleService
		TaskService  TaskService
	}
)

func NewService(repos *repositories.Repository) *Service {
	return &Service{
		CourseService: NewCourseServiceImpl(repos.CourseRepository),
		ModuleService: NewModuleService(repos.ModuleRepository),
		TaskService: NewTaskService(repos.TaskRepository),
	}
}