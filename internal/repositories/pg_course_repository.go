package repositories

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/repositories/models"
	"log/slog"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type PgCourseRepository struct {
	conn *sqlx.DB
}

func NewCourseRepositoryImpl(conn *sqlx.DB) CourseRepository {
	return &PgCourseRepository{
		conn: conn,
	}
}

func (cr *PgCourseRepository) GetAllCourses(filter *dto.CourseFilter) ([]models.Course, error) {
	baseQuery := `SELECT * FROM t_course`
	conditions := []string{}
	args := map[string]interface{}{}

	if filter != nil {
		if filter.AuthorID != nil {
			conditions = append(conditions, "id_author = :author_id")
			args["author_id"] = filter.AuthorID
		}
		if filter.NameContains != nil {
			conditions = append(conditions, "c_name ILIKE :name_contains")
			args["name_contains"] = "%" + *filter.NameContains + "%"
		}
	}

	if len(conditions) > 0 {
		baseQuery += " WHERE " + strings.Join(conditions, " AND ")
	}

	namedStmt, err := cr.conn.PrepareNamed(baseQuery)
	if err != nil {
		slog.Error("Error preparing named statement", "query", baseQuery, "error", err)
		return nil, err
	}
	defer namedStmt.Close()

	var courses []models.Course
	err = namedStmt.Select(&courses, args)
	if err != nil {
		slog.Error("Error executing select", "query", baseQuery, "error", err)
		return nil, err
	}

	return courses, nil
}

func (cr *PgCourseRepository) GetCourse(id uuid.UUID) (*models.Course, error) {
	const query = `SELECT * FROM t_course WHERE id = $1`
	course := &models.Course{}
	err := cr.conn.Get(course, query, id)
	if err != nil {
		slog.Error("Error executing select", "query", query, "error", err)
		return nil, err
	}

	return course, nil
}


func (cr *PgCourseRepository) Create(course *dto.CreateCourse) (*uuid.UUID, error) {
	const query = `
		INSERT INTO t_course (
			c_name, c_description, c_date_start, c_date_end, c_image_path, id_author
		) VALUES (
			:name, :description, :date_start, :date_end, :image_path, :author_id
		) RETURNING id
	`

	namedStmt, err := cr.conn.PrepareNamed(query)
	if err != nil {
		slog.Error("Error preparing named statement", "query", query, "error", err)
		return nil, err
	}
	defer namedStmt.Close()

	var insertedID uuid.UUID
	err = namedStmt.Get(&insertedID, map[string]interface{}{
		"name":        course.Name,
		"description": course.Description,
		"date_start":  course.DateStart,
		"date_end":    course.DateEnd,
		"image_path":  course.ImagePath,
		"author_id":   course.AuthorID,
	})

	if err != nil {
		slog.Error("Error executing insert", "query", query, "error", err)
		return nil, err
	}

	return &insertedID, nil
}


func (cr *PgCourseRepository) Clone(course *dto.CloneCourseRequest) (*uuid.UUID, error) {
	query := `
		SELECT clone_course_with_content($1, $2, $3, $4, $5, $6)
	`

	var clonedCourseID uuid.UUID
	err := cr.conn.QueryRow(
		query,
		course.ParentCourseID,
		course.Name,
		course.AuthorID,
		course.DateStart,
		course.DateEnd,
		course.ImagePath,
	).Scan(&clonedCourseID)

	if err != nil {
		slog.Error("Error executing clone course", "query", query, "error", err)
		return nil, err
	}

	return &clonedCourseID, nil
}


