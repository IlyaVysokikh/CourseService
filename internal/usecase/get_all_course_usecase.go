package usecase

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/services"
	ierrors "CourseService/pkg/errors"
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
		if err == ierrors.ErrInternal {
			slog.Error("Error getting all courses", "error", err)
			return nil, ierrors.New(ierrors.ErrInternal, "failed to get all courses", err)
		}

		slog.Error("Unexpected error while getting courses", "error", err)
		return nil, err
	}

	return courses, nil
}