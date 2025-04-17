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
	}

	ModuleService interface{
		GetModulesByCourse(ctx context.Context, courseID uuid.UUID) ([]dto.ModuleList, error)
	}

	TaskService interface {
		GetTaskCount(ctx context.Context, moduleId uuid.UUID) (int, error)
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