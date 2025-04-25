package module

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/services"
	"CourseService/internal/usecase/shared"
	ierrors "CourseService/pkg/errors"
	"context"
	"errors"
	"github.com/google/uuid"
	"log/slog"
)

type CreateModuleAttachmentUseCaseImpl struct {
	attachmentService services.ModuleAttachmentService
}

func NewCreateModuleAttachmentUseCase(
	attachmentService services.ModuleAttachmentService) shared.CreateModuleAttachmentUseCase {
	return &CreateModuleAttachmentUseCaseImpl{
		attachmentService: attachmentService,
	}
}

func (u *CreateModuleAttachmentUseCaseImpl) Handle(
	ctx context.Context, moduleId uuid.UUID, request dto.CreateModuleAttachmentRequest) (dto.CreateModuleAttachmentResponse, error) {

	data, err := u.attachmentService.CreateAttachment(ctx, moduleId, request)
	response := dto.CreateModuleAttachmentResponse{
		Success: err == nil,
		Message: "",
		Data:    data,
	}

	if err == nil {
		return response, nil
	}

	response.Message = err.Error()
	if errors.Is(err, ierrors.ErrNotFound) {
		slog.Error(
			"The module for which the materials were loaded was not found",
			"err", err, "moduleId", moduleId)
		return response, ierrors.ErrNotFound
	}

	slog.Error("Failed to load module materials", "err", err, "moduleId", moduleId)

	return response, ierrors.ErrInternal
}
