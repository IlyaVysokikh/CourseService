package usecase

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/services"
	"context"
	"log/slog"
)

type GetAllCourseUsecaseImpl struct {
	courseService services.CourseService	
}

func NewGetAllCourseUsecase(courseService services.CourseService) GetAllCourseUsecase {
	return &GetAllCourseUsecaseImpl{
		courseService: courseService,
	}
}

func (u *GetAllCourseUsecaseImpl) Handle(ctx context.Context, filter *dto.CourseFilter) ([]dto.CourseList, error) {
	courses, err := u.courseService.GetAllCourses(ctx, filter)
	if err != nil {
		slog.Error("Error getting all courses", "error", err)
		return nil, err
	}

	return courses, nil
}