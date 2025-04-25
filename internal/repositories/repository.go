package repositories

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/repositories/models"
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type (
	CourseRepository interface {
		GetAllCourses(filter *dto.CourseFilter) ([]models.Course, error)
		GetCourse(id uuid.UUID) (*models.Course, error)
		Create(course *dto.CreateCourse) (*uuid.UUID, error)
		Clone(course *dto.CloneCourseRequest) (*uuid.UUID, error)
		Delete(id uuid.UUID) error
		Update(id uuid.UUID, request dto.UpdateCourseRequest) error
	}

	ModuleRepository interface {
		GetModulesByCourse(courseID uuid.UUID) ([]models.Module, error)
		CreateModules(courseID uuid.UUID, modules []dto.CreateModule) error
		UpdateModules(courseID uuid.UUID, modules []dto.CreateModule) error
		GetModule(moduleID uuid.UUID) (*models.Module, error)
		DeleteModule(id uuid.UUID) error
	}

	TaskRepository interface {
		GetTasks(moduleId uuid.UUID) ([]models.Task, error)
		GetTasksByModule(moduleId uuid.UUID) ([]models.Task, error)
		GetTask(taskId uuid.UUID) (*models.Task, error)
		DeleteTask(id uuid.UUID) error
	}

	ModuleAttachmentRepository interface {
		GetAllByModule(moduleId uuid.UUID) ([]*models.ModuleAttachment, error)
		Create(ctx context.Context, moduleId uuid.UUID, req dto.CreateModuleAttachmentRequest) (*models.ModuleAttachment, error)
	}

	Repository struct {
		CourseRepository           CourseRepository
		ModuleRepository           ModuleRepository
		TaskRepository             TaskRepository
		ModuleAttachmentRepository ModuleAttachmentRepository
	}
)

func NewRepository(conn *sqlx.DB) *Repository {
	return &Repository{
		CourseRepository:           NewCourseRepositoryImpl(conn),
		ModuleRepository:           NewModuleRepositoryImpl(conn),
		TaskRepository:             NewTaskRepository(conn),
		ModuleAttachmentRepository: NewModuleAttachmentRepository(conn),
	}
}
