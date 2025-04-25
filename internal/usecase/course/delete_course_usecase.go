package course

import (
	"CourseService/internal/services"
	"CourseService/internal/usecase/shared"
	ierrors "CourseService/pkg/errors"
	"context"
	"errors"
	"github.com/google/uuid"
)

type DeleteCourseUseCaseImpl struct {
	courseService services.CourseService
}

func NewDeleteCourseUseCase(courseService services.CourseService) shared.DeleteCourseUseCase {
	return &DeleteCourseUseCaseImpl{
		courseService: courseService,
	}
}

func (uc *DeleteCourseUseCaseImpl) Handle(ctx context.Context, id uuid.UUID) error {
	if err := uc.courseService.DeleteCourse(ctx, id); err != nil {
		if errors.Is(err, ierrors.ErrNotFound) {
			return ierrors.ErrNotFound
		}

		return ierrors.ErrInternal
	}

	return nil
}
