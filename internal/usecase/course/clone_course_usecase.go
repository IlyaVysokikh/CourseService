package course

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/services"
	"CourseService/internal/usecase/shared"
	ierrors "CourseService/pkg/errors"
	"errors"

	"context"
	"log/slog"
)

type CloneCourseUseCaseImpl struct {
	cs services.CourseService
}

func NewCloneCourseUseCase(cs services.CourseService) shared.CloneCourseUseCase {
	return &CloneCourseUseCaseImpl{
		cs: cs,
	}
}

func (u *CloneCourseUseCaseImpl) Handle(ctx context.Context, course *dto.CloneCourseRequest) (dto.CreateCourseResponse, error) {
	newCourseID, err := u.cs.CloneCourse(ctx, course)
	if err != nil {
		if errors.Is(err, ierrors.ErrInternal) {
			slog.Error("Error cloning course", "error", err)
			return dto.CreateCourseResponse{}, ierrors.New(ierrors.ErrInternal, "failed to clone course", err)
		}
		slog.Error("Unexpected error cloning course", "error", err)
		return dto.CreateCourseResponse{}, err
	}

	return dto.CreateCourseResponse{Id: newCourseID}, nil
}
