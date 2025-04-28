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
	CreateModulesUseCase          shared.CreateModulesUseCase
	GetModuleUseCase              shared.GetModuleUseCase
	DeleteModuleUseCase           shared.DeleteModuleUseCase
	CreateModuleAttachmentUseCase shared.CreateModuleAttachmentUseCase
}

func NewModulesHandler(useCase *usecase.UseCase) *ModulesHandler {
	return &ModulesHandler{
		BaseHandler:                   BaseHandler{},
		CreateModulesUseCase:          useCase.CreateModulesUseCase,
		GetModuleUseCase:              useCase.GetModuleUseCase,
		DeleteModuleUseCase:           useCase.DeleteModuleUseCase,
		CreateModuleAttachmentUseCase: useCase.CreateModuleAttachment,
	}
}

// CreateModulesHandler godoc
// @Summary Создание модулей
// @Description Создает один или несколько модулей курса
// @Tags Modules
// @Accept json
// @Produce json
// @Param request body dto.CreateModulesRequest true "Данные для создания модулей"
// @Success 201 {string} string "Modules created successfully"
// @Failure 400 {object} dto.ErrorResponse "Bad request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /modules [post]
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

// GetModuleHandler godoc
// @Summary Получение модуля
// @Description Получает модуль по его ID
// @Tags Modules
// @Accept json
// @Produce json
// @Param id path string true "UUID модуля"
// @Success 200 {object} dto.GetModuleResponse "Данные модуля"
// @Failure 400 {object} dto.ErrorResponse "Bad request"
// @Failure 404 {object} dto.ErrorResponse "Module not found"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /modules/{id} [get]
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

// DeleteModuleHandler godoc
// @Summary Удаление модуля
// @Description Удаляет модуль по его ID
// @Tags Modules
// @Accept json
// @Produce json
// @Param id path string true "UUID модуля"
// @Success 204 {string} string "No Content"
// @Failure 400 {object} dto.ErrorResponse "Bad request"
// @Failure 404 {object} dto.ErrorResponse "Module not found"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /modules/{id} [delete]
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

// CreateAttachment godoc
// @Summary Добавление вложения к модулю
// @Description Создает вложение для указанного модуля
// @Tags Modules
// @Accept json
// @Produce json
// @Param id path string true "UUID модуля"
// @Param request body dto.CreateModuleAttachmentRequest true "Данные вложения"
// @Success 201 {object} dto.CreateModuleAttachmentResponse "Attachment created"
// @Failure 400 {object} dto.ErrorResponse "Bad request"
// @Failure 404 {object} dto.ErrorResponse "Module not found"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /modules/{id}/attachments [post]
func (h *ModulesHandler) CreateAttachment(ctx *gin.Context) {
	//dto.CreateModuleAttachmentResponse, error)
	moduleId, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		slog.Error("Error parsing module ID", "error", err)
		h.badRequest(ctx, err)
	}

	var request dto.CreateModuleAttachmentRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		slog.Error("Error binding JSON", "error", err)
		h.badRequest(ctx, err)
	}

	response, err := h.CreateModuleAttachmentUseCase.Handle(ctx, moduleId, request)
	if err != nil {
		if errors.Is(err, ierrors.ErrNotFound) {
			slog.Warn("Module not found", "moduleID", moduleId)
			h.notFound(ctx, err)
		}

		slog.Error("Error creating module", "moduleID", moduleId, "error", err)
		h.internalServerError(ctx, err)
	}

	h.created(ctx, response)
	return

}
