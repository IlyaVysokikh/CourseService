package rest

import (
	"CourseService/internal/usecase"
	"CourseService/internal/usecase/shared"
	ierrors "CourseService/pkg/errors"
	"errors"

	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TasksHandler struct {
	BaseHandler
	GetTaskUseCase    shared.GetTaskUseCase
	DeleteTaskUseCase shared.DeleteTaskUseCase
}

func NewTasksHandler(useCase *usecase.UseCase) *TasksHandler {
	return &TasksHandler{
		BaseHandler:       BaseHandler{},
		GetTaskUseCase:    useCase.GetTaskUseCase,
		DeleteTaskUseCase: useCase.DeleteTaskUseCase,
	}
}

func (h *TasksHandler) GetTaskHandler(ctx *gin.Context) {
	taskId, err := uuid.Parse(ctx.Param("taskId"))
	if err != nil {
		slog.Error("failed to parse task id", "error", err)
		h.badRequest(ctx, err)
		return
	}

	task, err := h.GetTaskUseCase.Handle(ctx, taskId)
	if err != nil {
		if errors.Is(err, ierrors.ErrNotFound) {
			h.notFound(ctx, err)
			return
		}

		slog.Error("failed to get task", "error", err)
		h.internalServerError(ctx, err)
		return
	}

	h.ok(ctx, task)
}

func (h *TasksHandler) DeleteTaskHandler(ctx *gin.Context) {
	taskId, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		slog.Error("failed to parse task id", "error", err)
		h.badRequest(ctx, err)
		return
	}

	if err := h.DeleteTaskUseCase.Handle(ctx, taskId); err != nil {
		if errors.Is(err, ierrors.ErrNotFound) {
			h.notFound(ctx, err)
		}

		h.internalServerError(ctx, err)
		return
	}

	h.noContent(ctx)
	return
}
