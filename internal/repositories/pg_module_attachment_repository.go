package repositories

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/repositories/models"
	ierrors "CourseService/pkg/errors"
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

type PgModuleAttachmentRepository struct {
	conn *sqlx.DB
}

func NewModuleAttachmentRepository(conn *sqlx.DB) *PgModuleAttachmentRepository {
	return &PgModuleAttachmentRepository{conn: conn}
}

func (r *PgModuleAttachmentRepository) GetAllByModule(moduleId uuid.UUID) ([]*models.ModuleAttachment, error) {
	var attachments []*models.ModuleAttachment
	query := `SELECT * FROM t_module_attachment WHERE id_module = $1`
	err := r.conn.Select(&attachments, query, moduleId)
	if err != nil {
		return nil, err
	}
	return attachments, nil
}

func (r *PgModuleAttachmentRepository) Create(ctx context.Context, moduleId uuid.UUID, req dto.CreateModuleAttachmentRequest) (*models.ModuleAttachment, error) {
	const query = `
		INSERT INTO t_module_attachment (id_module, c_file_name, c_bucket, c_visible) 
		VALUES ($1, $2, $3, $4)
		RETURNING *;`

	var inserted models.ModuleAttachment

	err := r.conn.GetContext(ctx, &inserted, query, moduleId, req.FileName, req.Bucket, req.Visible)
	if err != nil {
		slog.Error("failed to insert module attachment", "error", err, "query", query)
		return nil, ierrors.ErrInternal
	}

	return &inserted, nil
}
