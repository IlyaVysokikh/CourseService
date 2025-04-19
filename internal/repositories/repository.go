package repositories

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/repositories/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type (
	CourseRepository interface{
		GetAllCourses(filter *dto.CourseFilter) ([]models.Course, error)
		GetCourse(id uuid.UUID) (*models.Course, error)
		Create(course *dto.CreateCourse) (*uuid.UUID, error)
		Clone(course *dto.CloneCourseRequest) (*uuid.UUID, error)
	}

	ModuleRepository interface{
		GetModulesByCourse(courseID uuid.UUID) ([]models.Module, error)
		CreateModules(courseID uuid.UUID, modules []dto.CreateModule) error
		UpdateModules(courseID uuid.UUID, modules []dto.CreateModule) error
		GetModule(moduleID uuid.UUID) (*models.Module, error)
	}

	TaskRepository interface {
		GetTasks(moduleId uuid.UUID) ([]models.Task, error)
		GetTasksByModule(moduleId uuid.UUID) ([]models.Task, error)
	}

	Repository struct {
		CourseRepository CourseRepository
		ModuleRepository ModuleRepository
		TaskRepository   TaskRepository
	}
)

func NewRepository(conn *sqlx.DB) *Repository {
	return &Repository{
		CourseRepository: NewCourseRepositoryImpl(conn),
		ModuleRepository: NewModuleRepositoryImpl(conn),
		TaskRepository: NewTaskRepository(conn),
	}
}	