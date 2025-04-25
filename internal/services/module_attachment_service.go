package services

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/repositories"
	"CourseService/internal/repositories/models"
	"context"

	"github.com/google/uuid"
)

type ModuleAttachmentServiceImpl struct {
	ModuleAttachmentRepository repositories.ModuleAttachmentRepository
}

func NewModuleAttachmentService(moduleAttachmentRepository repositories.ModuleAttachmentRepository) *ModuleAttachmentServiceImpl {
	return &ModuleAttachmentServiceImpl{
		ModuleAttachmentRepository: moduleAttachmentRepository,
	}
}

func (s *ModuleAttachmentServiceImpl) GetAllByModule(ctx context.Context, moduleId uuid.UUID) ([]*models.ModuleAttachment, error) {
	attachments, err := s.ModuleAttachmentRepository.GetAllByModule(moduleId)
	if err != nil {
		return nil, err
	}

	var visibleAttachments []*models.ModuleAttachment
	for _, attachment := range attachments {
		if attachment.Visible {
			visibleAttachments = append(visibleAttachments, attachment)
		}
	}

	return visibleAttachments, nil
}

func (s *ModuleAttachmentServiceImpl) CreateAttachment(
	ctx context.Context, moduleId uuid.UUID, attachment dto.CreateModuleAttachmentRequest) (*models.ModuleAttachment, error) {
	return s.ModuleAttachmentRepository.Create(ctx, moduleId, attachment)
}
