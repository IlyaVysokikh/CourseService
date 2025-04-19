package usecase

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/services"
	ierrors "CourseService/pkg/errors"
	"context"
	"log/slog"
)

type CreateCourseUsecaseImpl struct {
	cs services.CourseService
}

func NewCreateCourseUsecase(cs services.CourseService) CreateCourseUsecase {
	return &CreateCourseUsecaseImpl{
		cs: cs,
	}
}

func (u *CreateCourseUsecaseImpl) Handle(ctx context.Context, course *dto.CreateCourse) (dto.CreateCourseResponse, error) {
	createdCourseId, err := u.cs.CreateCourse(ctx, course)
	if err != nil {
		if err == ierrors.ErrInternal {
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