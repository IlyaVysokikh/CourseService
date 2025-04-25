package rest

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/usecase"
	"CourseService/internal/usecase/shared"
	ierrors "CourseService/pkg/errors"
	"errors"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CoursesHandler struct {
	BaseHandler
	GetAllCourseUseCase shared.GetAllCourseUseCase
	GetCourseUseCase    shared.GetCourseUseCase
	CreateCourseUseCase shared.CreateCourseUseCase
	CloneCourseUseCase  shared.CloneCourseUseCase
	DeleteCourseUseCase shared.DeleteCourseUseCase
	UpdateCourseUseCase shared.UpdateCourseUseCase
}

func NewCoursesHandler(useCase *usecase.UseCase) *CoursesHandler {
	return &CoursesHandler{
		BaseHandler:         BaseHandler{},
		GetAllCourseUseCase: useCase.GetAllCourseUseCase,
		GetCourseUseCase:    useCase.GetCourseUseCase,
		CreateCourseUseCase: useCase.CreateCourseUseCase,
		CloneCourseUseCase:  useCase.CloneCourseUseCase,
		DeleteCourseUseCase: useCase.DeleteCourseUseCase,
		UpdateCourseUseCase: useCase.UpdateCourseUseCase,
	}
}

func (h *CoursesHandler) GetAllCoursesHandler(ctx *gin.Context) {
	courseFilter := &dto.CourseFilter{}
	if err := ctx.ShouldBindQuery(courseFilter); err != nil {
		slog.Error("Error binding query", "error", err)
		h.badRequest(ctx, err)
		return
	}

	courses, err := h.GetAllCourseUseCase.Handle(ctx, courseFilter)
	if err != nil {
		if errors.Is(err, ierrors.ErrInternal) {
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

func (h *CoursesHandler) GetCourseHandler(ctx *gin.Context) {
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

	course, err := h.GetCourseUseCase.Handle(ctx, uuidId)
	if err != nil {
		if errors.Is(err, ierrors.ErrNotFound) {
			slog.Warn("Course not found", "courseID", id)
			h.notFound(ctx, err)
			return
		}

		if errors.Is(err, ierrors.ErrInternal) {
			slog.Error("Error getting course", "error", err)
			h.internalServerError(ctx, err)
			return
		}

		if errors.Is(err, ierrors.ErrInvalidInput) {
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

func (h *CoursesHandler) CreateCourseHandler(ctx *gin.Context) {
	course := &dto.CreateCourse{}
	if err := ctx.ShouldBindJSON(course); err != nil {
		slog.Error("Error binding json", "error", err)
		h.badRequest(ctx, err)
		return
	}

	courseResponse, err := h.CreateCourseUseCase.Handle(ctx, course)
	if err != nil {
		slog.Error("Error creating course", "error", err)
		h.badRequest(ctx, err)
		return
	}

	h.created(ctx, courseResponse)
}

func (h *CoursesHandler) CloneCourseHandler(ctx *gin.Context) {
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

	clonedID, err := h.CloneCourseUseCase.Handle(ctx, &req)
	if err != nil {
		slog.Error("Error cloning course", "error", err)
		h.internalServerError(ctx, err)
		return
	}

	h.created(ctx, clonedID)
}

func (h *CoursesHandler) DeleteCourseHandler(ctx *gin.Context) {
	idUuid, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		slog.Error("Error parsing course id uuid", "error", err)
		h.badRequest(ctx, err)
		return
	}

	err = h.DeleteCourseUseCase.Handle(ctx, idUuid)
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

func (h *CoursesHandler) UpdateCourseHandler(ctx *gin.Context) {
	idUuid, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		slog.Error("Error parsing course id uuid", "error", err)
		h.badRequest(ctx, ierrors.ErrInvalidInput)
	}
	var request dto.UpdateCourseRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		slog.Error("Error binding json", "error", err)
		h.badRequest(ctx, err)
	}

	if err = h.UpdateCourseUseCase.Handle(ctx, idUuid, request); err != nil {
		if errors.Is(err, ierrors.ErrNotFound) {
			h.notFound(ctx, err)
			return
		}

		if errors.Is(err, ierrors.ErrInvalidInput) {
			h.badRequest(ctx, err)
			return
		}

		h.internalServerError(ctx, err)
		return
	}

	h.ok(ctx, nil)
}
