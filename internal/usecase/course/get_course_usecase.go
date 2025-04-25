package course

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/services"
	"CourseService/internal/usecase/shared"
	ierrors "CourseService/pkg/errors"
	"context"
	"errors"
	"log/slog"

	"github.com/google/uuid"
)

type GetCourseUseCaseImpl struct {
	cs services.CourseService
	ms services.ModuleService
	ts services.TaskService
}

func NewGetCourseUseCase(cs services.CourseService, ms services.ModuleService, ts services.TaskService) shared.GetCourseUseCase {
	return &GetCourseUseCaseImpl{
		cs: cs,
		ms: ms,
		ts: ts,
	}
}

func (gs *GetCourseUseCaseImpl) Handle(ctx context.Context, id uuid.UUID) (*dto.Course, error) {

	course, err := gs.cs.GetCourse(ctx, id)
	if err != nil {
		if errors.Is(err, ierrors.ErrNotFound) {
			slog.Warn("course not found", "courseID", id)
			return nil, ierrors.New(ierrors.ErrNotFound, "course not found", err)
		}

		if errors.Is(err, ierrors.ErrInvalidInput) {
			slog.Warn("invalid input maybe dates", "courseID", id)
			return nil, ierrors.New(ierrors.ErrInvalidInput, "invalid input", err)
		}

		if errors.Is(err, ierrors.ErrInternal) {
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
			if errors.Is(err, ierrors.ErrInternal) {
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
