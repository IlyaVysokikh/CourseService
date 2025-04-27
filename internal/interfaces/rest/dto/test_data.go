package dto

import "github.com/google/uuid"

type CreateTestDataRequest struct {
	TaskId uuid.UUID `json:"task_id"`
	Input  string    `json:"input"`
	Output string    `json:"output"`
}

type UpdateTestDataRequest struct {
	TaskId *uuid.UUID `json:"task_id"`
	Input  *string    `json:"input"`
	Output *string    `json:"output"`
}

type TestDataResponse struct {
	Id     uuid.UUID `json:"id"`
	TaskId uuid.UUID `json:"task_id"`
	Input  string    `json:"input"`
	Output string    `json:"output"`
}
