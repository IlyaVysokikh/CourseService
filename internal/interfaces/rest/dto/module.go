package dto

import (
	"github.com/google/uuid"
)

type Module struct {
	Id uuid.UUID `json:"id"`
	Name string `json:"name"`
	DateStart string `json:"date_start"`
	SequenceNumber int `json:"sequence_number"`
}

type ModuleList struct {
	Id uuid.UUID `json:"id"`
	Name string `json:"name"`
	DateStart string `json:"date_start"`
	SequenceNumber int `json:"sequence_number"`
	TaskCount int `json:"task_count"`
}


type CreateModulesRequest struct {
	Modules []CreateModule `json:"modules"`
}


type CreateModule struct {
	Id *uuid.UUID `json:"id"`
	Name string `json:"name"`
	DateStart string `json:"date_start"`
	SequenceNumber int `json:"sequence_number"`
}



type CreateModulesResponse struct {

}