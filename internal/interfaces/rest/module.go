package rest

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/usecase"
	"CourseService/internal/usecase/shared"
	ierrors "CourseService/pkg/errors"

	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log/slog"
)

type ModulesHandler struct {
	BaseHandler
	CreateModulesUseCase shared.CreateModulesUseCase
	GetModuleUseCase     shared.GetModuleUseCase
	DeleteModuleUseCase  shared.DeleteModuleUseCase
}

func NewModulesHandler(useCase *usecase.UseCase) *ModulesHandler {
	return &ModulesHandler{
		BaseHandler:          BaseHandler{},
		CreateModulesUseCase: useCase.CreateModulesUseCase,
		GetModuleUseCase:     useCase.GetModuleUseCase,
		DeleteModuleUseCase:  useCase.DeleteModuleUseCase,
	}
}

func (h *ModulesHandler) CreateModulesHandler(ctx *gin.Context) {
	var module dto.CreateModulesRequest

	if err := ctx.ShouldBindJSON(&module); err != nil {
		slog.Error("Error binding JSON", "error", err)
		h.badRequest(ctx, err)
		return
	}

	if err := h.CreateModulesUseCase.Handle(ctx, &module); err != nil {
		if errors.Is(err, ierrors.ErrNotFound) {
			slog.Error("Course not found", "err", err)
			return
		}

		slog.Error("Error creating modules", "error", err)
		h.internalServerError(ctx, err)
		return
	}

	h.created(ctx, "Modules created successfully")
}

func (h *ModulesHandler) GetModuleHandler(ctx *gin.Context) {
	moduleID := ctx.Param("id")

	moduleUuid, err := uuid.Parse(moduleID)
	if err != nil {
		slog.Error("Error parsing module ID", "error", err)
		h.badRequest(ctx, err)
		return
	}

	module, err := h.GetModuleUseCase.Handle(ctx, moduleUuid)
	if err != nil {
		if errors.Is(err, ierrors.ErrNotFound) {
			slog.Warn("Module not found", "moduleID", moduleUuid)
			h.notFound(ctx, err)
		} else {
			slog.Error("Error getting module", "moduleID", moduleUuid, "error", err)
			ctx.JSON(500, gin.H{"error": "Internal server error"})
		}
		return
	}

	h.ok(ctx, module)
}

func (h *ModulesHandler) DeleteModuleHandler(ctx *gin.Context) {
	moduleID := ctx.Param("id")
	moduleUuid, err := uuid.Parse(moduleID)
	if err != nil {
		slog.Error("Error parsing module ID", "error", err)
		h.badRequest(ctx, err)
		return
	}

	if err = h.DeleteModuleUseCase.Handle(ctx, moduleUuid); err != nil {
		if errors.Is(err, ierrors.ErrNotFound) {
			slog.Warn("Module not found", "moduleID", moduleUuid)
			h.notFound(ctx, err)
			return
		}

		slog.Error("Error deleting module", "moduleID", moduleUuid, "error", err)
		h.internalServerError(ctx, err)
		return
	}

	h.noContent(ctx)
	return
}
