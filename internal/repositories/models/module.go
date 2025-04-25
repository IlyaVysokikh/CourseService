package models

import "github.com/google/uuid"

type Module struct {
	ID             uuid.UUID `db:"id"`
	CourseID       uuid.UUID `db:"id_course"`
	Name           string    `db:"c_name"`
	DateStart      string    `db:"c_date_start"`
	SequenceNumber int       `db:"c_sequence_number"`
}

type ModuleAttachment struct {
	ID       string `db:"id"`
	ModuleID string `db:"id_module"`
	FileName string `db:"c_file_name"`
	Bucket   string `db:"c_bucket"`
	Visible  bool   `db:"c_visible"`
}

// Models used only for insert or update operations
type ModuleUpdate struct {
	Id             uuid.UUID `db:"id"`
	Name           string    `db:"c_name"`
	DateStart      string    `db:"c_date_start"`
	SequenceNumber int       `db:"c_sequence_number"`
}

type ModuleInsert struct {
	IdCourse       uuid.UUID `db:"id_course"`
	Name           string    `db:"c_name"`
	DateStart      string    `db:"c_date_start"`
	SequenceNumber int       `db:"c_sequence_number"`
}
