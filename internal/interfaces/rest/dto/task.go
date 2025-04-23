package dto

import "github.com/google/uuid"

type Task struct {
	Id uuid.UUID `json:"id"`
	Name string `json:"name"`
}

type TaskExtended struct {
	Id uuid.UUID `json:"id"`
	Name string `json:"name"`
	Text string `json:"text"`
	Language *string `json:"language"`
	InitialCode *string `json:"initial_code"`
	MemoryLimit *int `json:"memory_limit"`
	ExecutionTimeout *int `json:"execution_timeout"`
}