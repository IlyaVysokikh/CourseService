package dto

import (
	"github.com/google/uuid"
)

type CourseFilter struct {
	AuthorID     *uuid.UUID `form:"author_id" json:"author_id,omitempty" db:"id_author"`
	NameContains *string    `form:"name_contains" json:"name_contains,omitempty" db:"c_name"`
}


type CourseList struct {
	Id uuid.UUID `json:"id"`
	Name string `json:"name"`
	IsArchived bool `json:"is_archived"`
	ImagePath string `json:"image_path"`
}

type Course struct {
	Id uuid.UUID `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	DateStart string `json:"date_start"`
	DateEnd string `json:"date_end"`
	ImagePath string `json:"image_path"`
	AuthorID uuid.UUID `json:"author_id"`
	IsArchived bool `json:"is_archived"`
	Modules []ModuleList `json:"modules"`
}