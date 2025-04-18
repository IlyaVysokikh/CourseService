package repositories

import (
	"CourseService/internal/interfaces/rest/dto"
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

func (p *PgModuleRepository) CreateModules(courseID uuid.UUID, modules []dto.CreateModule) error {
    query := `INSERT INTO 
	t_module (id_course, c_name, c_date_start, c_sequence_number) VALUES
	(:id_course, :c_name, :c_date_start, :c_sequence_number)`

	var data []models.ModuleInsert
	for _, m := range modules {
		data = append(data, models.ModuleInsert{
			IdCourse:       courseID,
			Name:           m.Name,
			DateStart:      m.DateStart,
			SequenceNumber: m.SequenceNumber,
		})
	}

	_, err := p.conn.NamedExec(query, data)

	if err != nil {
		slog.Error("Error executing insert", "query", query, "error", err)
		return err
	}

	return nil
}

func (p *PgModuleRepository) UpdateModules(courseID uuid.UUID, modules []dto.CreateModule) error {
	for _, m := range modules {
		_, err := p.conn.NamedExec(`
			UPDATE t_module 
			SET c_name = :c_name, c_date_start = :c_date_start, c_sequence_number = :c_sequence_number 
			WHERE id = :id
		`, models.ModuleUpdate{
			Id:             *m.Id,
			Name:           m.Name,
			DateStart:      m.DateStart,
			SequenceNumber: m.SequenceNumber,
		})

		if err != nil {
			slog.Error("Error updating module", "id", m.Id, "error", err)
			return err
		}
	}
	return nil
}
