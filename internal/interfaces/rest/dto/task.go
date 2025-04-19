package dto

import "github.com/google/uuid"

type Task struct {
	Id uuid.UUID `json:"id"`
	Name string `json:"name"`
}