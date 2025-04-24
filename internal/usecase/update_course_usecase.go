package usecase

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/services"
	ierrors "CourseService/pkg/errors"
	"context"
	"errors"
	"github.com/google/uuid"
	"log/slog"
)

type UpdateCourseUseCaseImpl struct {
	courseService services.CourseService
}

func NewUpdateCourseUsecase(courseService services.CourseService) UpdateCourseUsecase {
	return &UpdateCourseUseCaseImpl{
		courseService: courseService,
	}
}

func (u UpdateCourseUseCaseImpl) Handle(ctx context.Context, id uuid.UUID, request dto.UpdateCourseRequest) error {

	if err := u.courseService.UpdateCourse(ctx, id, request); err != nil {
		if errors.Is(err, ierrors.ErrNotFound) {
			slog.Error("Course not found", "id", id)
			return ierrors.ErrNotFound
		}

		if errors.Is(err, ierrors.ErrInvalidInput) {
			slog.Error("Invalid input", "request", request)
			return ierrors.ErrInvalidInput
		}

		slog.Error("Course update failed", "id", id, "updateCourse", request, "err", err)
		return ierrors.ErrInternal
	}

	return nil
}
