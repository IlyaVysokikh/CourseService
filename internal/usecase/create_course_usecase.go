package usecase

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/services"
	"context"
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
		return dto.CreateCourseResponse{}, err
	}

	return dto.CreateCourseResponse{
		Id: createdCourseId,
	}, err
}