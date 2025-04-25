package course

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/services"
	"CourseService/internal/usecase/shared"
	ierrors "CourseService/pkg/errors"
	"context"
	"errors"
	"log/slog"
)

type GetAllCourseUseCaseImpl struct {
	courseService services.CourseService
}

func NewGetAllCourseUseCase(courseService services.CourseService) shared.GetAllCourseUseCase {
	return &GetAllCourseUseCaseImpl{
		courseService: courseService,
	}
}

func (u *GetAllCourseUseCaseImpl) Handle(ctx context.Context, filter *dto.CourseFilter) ([]dto.CourseList, error) {
	courses, err := u.courseService.GetAllCourses(ctx, filter)
	if err != nil {
		if errors.Is(err, ierrors.ErrInternal) {
			slog.Error("Error getting all courses", "error", err)
			return nil, ierrors.New(ierrors.ErrInternal, "failed to get all courses", err)
		}

		slog.Error("Unexpected error while getting courses", "error", err)
		return nil, err
	}

	return courses, nil
}
