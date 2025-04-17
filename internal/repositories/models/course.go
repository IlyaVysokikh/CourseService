package models

import "github.com/google/uuid"

type Course struct {
	ID          uuid.UUID    `db:"id"`
	Name        string `db:"c_name"`
	Description string `db:"c_description"`
	DateStart  string `db:"c_date_start"`
	DateEnd    string `db:"c_date_end"`
	ImagePath   string `db:"c_image_path"`
	AuthorID    uuid.UUID    `db:"id_author"`
	ParentCourseID *uuid.UUID `db:"id_parent_course"`
}
