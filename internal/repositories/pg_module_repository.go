package repositories

import (
	"CourseService/internal/repositories/models"
	"log/slog"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type PgModuleRepository struct {
	conn *sqlx.DB
}

func NewModuleRepositoryImpl(conn *sqlx.DB) ModuleRepository {
	return &PgModuleRepository{
		conn: conn,
	}
}

func (p *PgModuleRepository) GetModulesByCourse(courseID uuid.UUID) ([]models.Module, error) {
	query := `SELECT * FROM t_module WHERE id_course = $1`
	modules := []models.Module{}
	err := p.conn.Select(&modules, query, courseID)
	if err != nil {
		slog.Error("Error executing select", "query", query, "error", err)
		return nil, err
	}

	return modules, nil
}
