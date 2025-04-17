package usecase

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/services"
	"context"
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
		return dto.CreateCourseResponse{}, err
	}

	return dto.CreateCourseResponse{Id: newCourseID}, nil
}