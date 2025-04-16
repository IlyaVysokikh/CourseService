package repositories

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/repositories/models"
	"log/slog"
	"strings"

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
