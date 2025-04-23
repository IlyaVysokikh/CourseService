package rest

import (
	"CourseService/internal/interfaces/rest/dto"
	ierrors "CourseService/pkg/errors"
	"errors"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) GetAllCoursesHandler(ctx *gin.Context) {
	courseFilter := &dto.CourseFilter{}
	if err := ctx.ShouldBindQuery(courseFilter); err != nil {
		slog.Error("Error binding query", "error", err)
		h.badRequest(ctx, err)
		return
	}

	courses, err := h.usecases.GetAllCourseUsecase.Handle(ctx, courseFilter)
	if err != nil {
		if err == ierrors.ErrInternal {
			slog.Error("Error getting all courses", "error", err)
			h.internalServerError(ctx, err)
			return
		}

		slog.Error("Error getting all courses", "error", err)
		h.badRequest(ctx, err)
		return
	}

	h.ok(ctx, courses)
}

func (h *Handler) GetCourseHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		slog.Error("Error getting course id from params")
		h.badRequest(ctx, ierrors.ErrInvalidInput)
		return
	}

	uuidId, err := uuid.Parse(id)
	if err != nil {
		slog.Error("Error converting id to uuid", "error", err)
		h.badRequest(ctx, err)
	}

	course, err := h.usecases.GetCourseUsecase.Handle(ctx, uuidId)
	if err != nil {
		if err == ierrors.ErrNotFound {
			slog.Warn("Course not found", "courseID", id)
			h.notFound(ctx, err)
			return
		}

		if err == ierrors.ErrInternal {
			slog.Error("Error getting course", "error", err)
			h.internalServerError(ctx, err)
			return
		}

		if err == ierrors.ErrInvalidInput {
			slog.Warn("Invalid input", "courseID", id)
			h.badRequest(ctx, err)
			return
		}

		slog.Error("Error getting course", "error", err)
		h.badRequest(ctx, err)
		return
	}

	h.ok(ctx, course)
}

func (h *Handler) CreateCourseHandler(ctx *gin.Context) {
	course := &dto.CreateCourse{}
	if err := ctx.ShouldBindJSON(course); err != nil {
		slog.Error("Error binding json", "error", err)
		h.badRequest(ctx, err)
		return
	}

	courseResponse, err := h.usecases.CreateCourseUsecase.Handle(ctx, course)
	if err != nil {
		slog.Error("Error creating course", "error", err)
		h.badRequest(ctx, err)
		return
	}

	h.created(ctx, courseResponse)
}

func (h *Handler) CloneCourseHandler(ctx *gin.Context) {
	parentIDStr := ctx.Param("id")
	parentID, err := uuid.Parse(parentIDStr)
	if err != nil {
		slog.Error("Error parsing parent course ID", "error", err)
		h.badRequest(ctx, err)
		return
	}

	var req dto.CloneCourseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		slog.Error("Error binding json", "error", err)
		h.badRequest(ctx, err)
		return
	}

	req.ParentCourseID = parentID

	clonedID, err := h.usecases.CloneCourseUsecase.Handle(ctx, &req)
	if err != nil {
		slog.Error("Error cloning course", "error", err)
		h.internalServerError(ctx, err)
		return
	}

	h.created(ctx, clonedID)
}

func (h *Handler) DeleteCourseHandler(ctx *gin.Context) {
	idUuid, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		slog.Error("Error parsing course id uuid", "error", err)
		h.badRequest(ctx, err)
		return
	}

	err = h.usecases.DeleteCourseUsecase.Handle(ctx, idUuid)
	if err != nil {
		if errors.Is(err, ierrors.ErrNotFound) {
			h.notFound(ctx, err)
			return
		}

		h.internalServerError(ctx, err)
		return
	}

	h.noContent(ctx)
	return
}
