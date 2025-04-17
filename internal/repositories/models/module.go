package models

import "github.com/google/uuid"

type Module struct {
	ID          uuid.UUID    `db:"id"`
	CourseID	uuid.UUID    `db:"id_course"`
	Name        string `db:"c_name"`
	DateStart   string `db:"c_date_start"`
	SequenceNumber int `db:"c_sequence_number"`
}