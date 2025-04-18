package rest

import (
	"CourseService/internal/interfaces/rest/dto"

	"github.com/gin-gonic/gin"
	"log/slog"
	"github.com/google/uuid"
)

func (h *Handler) CreateModulesHandler(ctx *gin.Context) {
	courseID := ctx.Param("id")
	var module dto.CreateModulesRequest

	if err := ctx.ShouldBindJSON(&module); err != nil {
		slog.Error("Error binding JSON", "error", err)
		h.badRequest(ctx, err)
		return
	}

	courseUuid, err := uuid.Parse(courseID)
	if err != nil {
		slog.Error("Error parsing course ID", "error", err)
		h.badRequest(ctx, err)
		return
	}
	err = h.usecases.CreateModulesUsecase.Handle(ctx, courseUuid, &module)
	if err != nil {
		slog.Error("Error creating modules", "error", err)
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	
	h.created(ctx, "Modules created successfully")
}