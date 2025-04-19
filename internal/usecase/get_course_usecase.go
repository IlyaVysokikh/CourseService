package usecase

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/services"
	"context"
	"log/slog"
	ierrors "CourseService/pkg/errors"


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
		if err == ierrors.ErrNotFound {
			slog.Warn("course not found", "courseID", id)
			return nil, ierrors.New(ierrors.ErrNotFound, "course not found", err)
		}
		
		if err == ierrors.ErrInvalidInput {
			slog.Warn("invalid input maybe dates", "courseID", id)
			return nil, ierrors.New(ierrors.ErrInvalidInput, "invalid input", err)
		}

		if err == ierrors.ErrInternal {
			slog.Error("error getting course", "error", err)
			return nil, ierrors.New(ierrors.ErrInternal, "failed to get course", err)
		}

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
			if err == ierrors.ErrInternal {
				slog.Warn("tasks not found", "moduleID", course.Modules[i].Id)
				course.Modules[i].TaskCount = 0
				continue
			}
			slog.Error("error getting task count", "error", err)
			return nil, err
		}
		course.Modules[i].TaskCount = tasksCount  
	}

	return course, nil
}