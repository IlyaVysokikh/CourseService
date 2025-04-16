package rest

import (
	"CourseService/internal/interfaces/rest/dto"
	"log/slog"
	"github.com/gin-gonic/gin"
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