package dto

import (
	"CourseService/internal/repositories/models"

	"github.com/google/uuid"
)

type Module struct {
	Id             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	DateStart      string    `json:"date_start"`
	SequenceNumber int       `json:"sequence_number"`
}

type ModuleList struct {
	Id             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	DateStart      string    `json:"date_start"`
	SequenceNumber int       `json:"sequence_number"`
	TaskCount      int       `json:"task_count"`
}

type CreateModulesRequest struct {
	CourseId uuid.UUID      `json:"course_id"`
	Modules  []CreateModule `json:"modules"`
}

type CreateModule struct {
	Id             *uuid.UUID `json:"id"`
	Name           string     `json:"name"`
	DateStart      string     `json:"date_start"`
	SequenceNumber int        `json:"sequence_number"`
}

type CreateModulesResponse struct {
}

type GetModuleResponse struct {
	Module     GetModule                  `json:"module"`
	Tasks      []Task                     `json:"tasks"`
	Attachment []*models.ModuleAttachment `json:"attachment"`
}

type GetModule struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type CreateModuleAttachmentRequest struct {
	FileName string `json:"file_name"`
	Bucket   string `json:"bucket"`
	Visible  bool   `json:"visible"`
}

type CreateModuleAttachmentResponse struct {
	Success bool                     `json:"status"`
	Message string                   `json:"message"`
	Data    *models.ModuleAttachment `json:"data,omitempty"`
}
