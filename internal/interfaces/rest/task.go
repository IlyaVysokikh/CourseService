package rest

import (
	ierrors "CourseService/pkg/errors"
	"errors"

	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) GetTaskHandler(ctx *gin.Context) {
	taskId, err := uuid.Parse(ctx.Param("taskId"))
	if err != nil {
		slog.Error("failed to parse task id", "error", err)
		h.badRequest(ctx, err)
		return
	}

	task, err := h.useCases.GetTaskUseCase.Handle(ctx, taskId)
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
