package repositories

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"CourseService/internal/repositories/models"
)

type PgModuleAttachmentRepository struct {
	conn *sqlx.DB
}

func NewModuleAttachmentRepository(conn *sqlx.DB) *PgModuleAttachmentRepository {
	return &PgModuleAttachmentRepository{conn: conn}
}

func (r * PgModuleAttachmentRepository) GetAllByModule(moduleId uuid.UUID) ([]*models.ModuleAttachment, error) {
	var attachments []*models.ModuleAttachment
	query := `SELECT * FROM t_module_attachment WHERE id_module = $1`
	err := r.conn.Select(&attachments, query, moduleId)
	if err != nil {
		return nil, err
	}
	return attachments, nil
}