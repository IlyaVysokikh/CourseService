package services

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/repositories"
	"context"
)

type (
	CourseService interface{
		GetAllCourses(ctx context.Context, filter *dto.CourseFilter) ([]dto.CourseList, error)
	}

	Service struct {
		CourseService CourseService
	}
)

func NewService(repos *repositories.Repository) *Service {
	return &Service{
		CourseService: NewCourseServiceImpl(repos.CourseRepository),
	}
}