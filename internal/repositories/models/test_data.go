package models

import "github.com/google/uuid"

type ProgrammingTestData struct {
	Id     uuid.UUID `db:"id"`
	TaskId uuid.UUID `db:"id_task"`
	Input  string    `db:"input"`
	Output string    `db:"output"`
}
