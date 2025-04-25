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

type CreateCourseUseCaseImpl struct {
	cs services.CourseService
}

func NewCreateCourseUseCase(cs services.CourseService) shared.CreateCourseUseCase {
	return &CreateCourseUseCaseImpl{
		cs: cs,
	}
}

func (u *CreateCourseUseCaseImpl) Handle(ctx context.Context, course *dto.CreateCourse) (dto.CreateCourseResponse, error) {
	createdCourseId, err := u.cs.CreateCourse(ctx, course)
	if err != nil {
		if errors.Is(err, ierrors.ErrInternal) {
			slog.Error("Error creating course", "error", err)
			return dto.CreateCourseResponse{}, ierrors.New(ierrors.ErrInternal, "failed to create course", err)
		}

		slog.Error("Unexpected error creating course", "error", err)
		return dto.CreateCourseResponse{}, err
	}

	return dto.CreateCourseResponse{
		Id: createdCourseId,
	}, err
}
