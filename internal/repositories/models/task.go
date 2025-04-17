package models

import "github.com/google/uuid"

type Task struct {
    Id               uuid.UUID `db:"id"`
    ModuleId         uuid.UUID `db:"id_module"`
    Text             string    `db:"c_text"`
    Language         *string   `db:"c_language"`         // nullable string
    InitialCode      *string   `db:"c_initial_code"`     // nullable text
    MemoryLimit      *int      `db:"c_memory_limit"`     // nullable integer
    ExecutionTimeout *int      `db:"c_execution_timeout"` // nullable integer
    SequenceNumber   int       `db:"c_sequence_number"`
}