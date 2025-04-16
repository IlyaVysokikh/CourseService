package repositories

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/repositories/models"

	"github.com/jmoiron/sqlx"
)

type (
	CourseRepository interface{
		GetAllCourses(filter *dto.CourseFilter) ([]models.Course, error)
	}
	Repository struct {
		CourseRepository CourseRepository
	}
)

func NewRepository(conn *sqlx.DB) *Repository {
	return &Repository{
		CourseRepository: NewCourseRepositoryImpl(conn),
	}
}	