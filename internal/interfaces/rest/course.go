package rest

import (
	"CourseService/internal/interfaces/rest/dto"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) GetAllCourses(ctx *gin.Context) {
	courseFilter := &dto.CourseFilter{}
	if err := ctx.ShouldBindQuery(courseFilter); err != nil {
		slog.Error("Error binding query", "error", err)
		h.badRequest(ctx, err)
		return
	}
	slog.Info("Course filter", "filter", courseFilter)

	courses, err := h.usecases.GetAllCourseUsecase.Handle(ctx, courseFilter)
	if err != nil {
		slog.Error("Error getting all courses", "error", err)
		h.badRequest(ctx, err)
		return
	}

	h.ok(ctx, courses)
}

func (h *Handler) GetCourse(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		slog.Error("Error getting course id from params")
		// h.badRequest(ctx, "id is empty")
		return
	}

	uuidId, err := uuid.Parse(id)
	if err != nil {
		slog.Error("Error converting id to uuid", "error", err)
		h.badRequest(ctx, err)
	}

	course, err := h.usecases.GetCourseUsecase.Handle(ctx, uuidId)
	if err != nil {
		slog.Error("Error getting course", "error", err)
		h.badRequest(ctx, err)
		return
	}

	h.ok(ctx, course)
}