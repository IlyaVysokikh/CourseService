package usecase

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/services"
	"context"
	"log/slog"

	"github.com/google/uuid"
)

type GetCourseUsecaseImpl struct {
	cs services.CourseService
	ms services.ModuleService
	ts services.TaskService
}

func NewGetCourseUsecase(cs services.CourseService, ms services.ModuleService, ts services.TaskService) GetCourseUsecase {
	return &GetCourseUsecaseImpl{
		cs: cs,
		ms: ms,
		ts: ts,
	}
}

func (gs *GetCourseUsecaseImpl) Handle(ctx context.Context, id uuid.UUID) (*dto.Course, error) {

	course, err := gs.cs.GetCourse(ctx, id)
	if err != nil {
		slog.Error("error getting course", "error", err)
		return nil, err
	}

	modules, err := gs.ms.GetModulesByCourse(ctx, id)
	if err != nil {
		slog.Error("error getting modules", "error", err)
		return nil, err
	}

	var modulesResponse []dto.ModuleList
	for _, module := range modules {
		modulesResponse = append(modulesResponse, dto.ModuleList{
			Id:             module.Id,
			Name:           module.Name,
			DateStart:      module.DateStart,
			SequenceNumber: module.SequenceNumber,
		})
	}


	course.Modules = modulesResponse

	for i := range course.Modules {
		tasksCount, err := gs.ts.GetTaskCount(ctx, course.Modules[i].Id)
		if err != nil {
			slog.Error("error getting task count", "error", err)
			return nil, err
		}
		course.Modules[i].TaskCount = tasksCount  
	}
	// todo добавить количество решенных задач в модуле
	return course, nil
}