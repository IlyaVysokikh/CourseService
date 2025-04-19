package usecase

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/services"
	ierrors "CourseService/pkg/errors"

	"context"
	"log/slog"

)


type CloneCourseUsecaseImpl struct {
	cs services.CourseService
}

func NewCloneCourseUsecase(cs services.CourseService) CloneCourseUsecase {
	return &CloneCourseUsecaseImpl{
		cs: cs,
	}
}

func (u *CloneCourseUsecaseImpl) Handle(ctx context.Context, course *dto.CloneCourseRequest) (dto.CreateCourseResponse, error) {
	newCourseID, err := u.cs.CloneCourse(ctx, course)
	if err != nil {
		if err == ierrors.ErrInternal {
			slog.Error("Error cloning course", "error", err)
			return dto.CreateCourseResponse{}, ierrors.New(ierrors.ErrInternal, "failed to clone course", err)
		}
		slog.Error("Unexpected error cloning course", "error", err)
		return dto.CreateCourseResponse{}, err
	}

	return dto.CreateCourseResponse{Id: newCourseID}, nil
}