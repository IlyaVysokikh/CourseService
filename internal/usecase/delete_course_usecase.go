package usecase

import (
	"CourseService/internal/services"
	ierrors "CourseService/pkg/errors"
	"context"
	"errors"
	"github.com/google/uuid"
)

type DeleteCourseUseCase struct {
	courseService services.CourseService
}

func NewDeleteCourseUseCase(courseService services.CourseService) *DeleteCourseUseCase {
	return &DeleteCourseUseCase{
		courseService: courseService,
	}
}

func (uc *DeleteCourseUseCase) Handle(ctx context.Context, id uuid.UUID) error {
	if err := uc.courseService.DeleteCourse(ctx, id); err != nil {
		if errors.Is(err, ierrors.ErrNotFound) {
			return ierrors.ErrNotFound
		}

		return ierrors.ErrInternal
	}

	return nil
}
