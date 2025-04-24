package repositories

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/repositories/models"
	ierrors "CourseService/pkg/errors"
	"database/sql"
	"errors"
	"fmt"

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
		return nil, ierrors.ErrInternal
	}
	defer namedStmt.Close()

	var courses []models.Course
	err = namedStmt.Select(&courses, args)
	if err != nil {
		slog.Error("Error executing select", "query", baseQuery, "error", err)
		return nil, ierrors.ErrInternal
	}

	return courses, nil
}

func (cr *PgCourseRepository) GetCourse(id uuid.UUID) (*models.Course, error) {
	const query = `SELECT * FROM t_course WHERE id = $1`
	course := &models.Course{}
	err := cr.conn.Get(course, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			slog.Warn("Course not found", "query", query, "error", err)
			return nil, ierrors.ErrNotFound
		}
		return nil, ierrors.ErrInternal
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
		return nil, ierrors.ErrInternal
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
		return nil, ierrors.ErrInternal
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
		return nil, ierrors.ErrInternal
	}

	return &clonedCourseID, nil
}

func (cr *PgCourseRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM t_course WHERE id = $1`
	_, err := cr.conn.Exec(query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ierrors.ErrNotFound
		}
		slog.Error("Error executing delete", "query", query, "error", err)
		return ierrors.ErrInternal
	}

	return nil
}

func (cr *PgCourseRepository) Update(id uuid.UUID, request dto.UpdateCourseRequest) error {
	query := "UPDATE t_course SET "
	args := []interface{}{}
	parts := []string{}

	if request.Name != nil {
		parts = append(parts, fmt.Sprintf("c_name = $%d", len(args)+1))
		args = append(args, *request.Name)
	}

	if request.Description != nil {
		parts = append(parts, fmt.Sprintf("c_description = $%d", len(args)+1))
		args = append(args, *request.Description)
	}

	if request.DateStart != nil {
		parts = append(parts, fmt.Sprintf("c_date_start = $%d", len(args)+1))
		args = append(args, *request.DateStart)
	}

	if request.DateEnd != nil {
		parts = append(parts, fmt.Sprintf("c_date_end = $%d", len(args)+1))
		args = append(args, *request.DateEnd)
	}

	if request.ImagePath != nil {
		parts = append(parts, fmt.Sprintf("c_image_path = $%d", len(args)+1))
		args = append(args, *request.ImagePath)
	}
	if len(parts) == 0 {
		return nil
	}

	query += strings.Join(parts, ", ") + fmt.Sprintf(" WHERE id = $%d", len(args)+1)
	args = append(args, id)

	result, err := cr.conn.Exec(query, args...)
	if err != nil {
		slog.Error("Error executing update", "query", query, "error", err)
		return ierrors.ErrInternal
	}

	affected, err := result.RowsAffected()
	if err != nil {
		slog.Error("Error executing update affected rows", "query", query, "error", err)
		return ierrors.ErrInternal
	}
	if affected == 0 {
		return ierrors.ErrNotFound
	}

	return nil
}
