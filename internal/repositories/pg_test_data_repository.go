package repositories

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/repositories/models"
	ierrors "CourseService/pkg/errors"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"log/slog"
	"strings"
)

type PgTestDataRepository struct {
	conn *sqlx.DB
}

func NewTestDataRepository(conn *sqlx.DB) TestDataRepository {
	return &PgTestDataRepository{
		conn: conn,
	}
}

func (r *PgTestDataRepository) GetAll(taskId uuid.UUID) ([]models.ProgrammingTestData, error) {
	const query = `SELECT * FROM t_programming WHERE id_task = $1`

	var testData []models.ProgrammingTestData
	err := r.conn.Select(&testData, query, taskId)
	if err != nil {
		return nil, err
	}

	return testData, nil
}

func (r *PgTestDataRepository) Get(id uuid.UUID) (models.ProgrammingTestData, error) {
	const query = `SELECT * FROM t_programming WHERE id = $1`

	var testData models.ProgrammingTestData
	err := r.conn.Get(&testData, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.ProgrammingTestData{}, ierrors.ErrNotFound
		}
		return models.ProgrammingTestData{}, err
	}

	return testData, nil
}

func (r *PgTestDataRepository) Create(ctx context.Context, request dto.CreateTestDataRequest) (uuid.UUID, error) {
	const query = `INSERT INTO t_programming (id, id_task, c_input, c_output) 
                   VALUES (:id, :id_task, :c_input, :c_output) 
                   RETURNING id`

	params := map[string]interface{}{
		"id_task":  request.TaskId,
		"c_input":  request.Input,
		"c_output": request.Output,
	}

	var returnedID uuid.UUID
	stmt, err := r.conn.PrepareNamedContext(ctx, query)
	if err != nil {
		slog.Error("Error preparing named statement", "query", query, "error", err)
		return uuid.Nil, ierrors.ErrInternal
	}
	defer stmt.Close()

	err = stmt.GetContext(ctx, &returnedID, params)
	if err != nil {
		slog.Error("Error executing named statement", "query", query, "error", err)
		return uuid.Nil, ierrors.ErrInternal
	}

	return returnedID, nil
}
func (r *PgTestDataRepository) Delete(id uuid.UUID) error {
	const query = `DELETE FROM t_programming WHERE id = $1`
	_, err := r.conn.Exec(query, id)
	if err != nil {
		slog.Error("Error executing named statement", "query", query, "error", err)
		return err
	}

	return nil
}

func (r *PgTestDataRepository) Update(ctx context.Context, id uuid.UUID, request dto.UpdateTestDataRequest) error {
	parts := []string{}
	args := []interface{}{}

	if request.TaskId != nil {
		parts = append(parts, fmt.Sprintf("id_task = $%d", len(args)+1))
		args = append(args, *request.TaskId)
	}

	if request.Input != nil {
		parts = append(parts, fmt.Sprintf("c_input = $%d", len(args)+1))
		args = append(args, *request.Input)
	}

	if request.Output != nil {
		parts = append(parts, fmt.Sprintf("c_output = $%d", len(args)+1))
		args = append(args, *request.Output)
	}

	if len(parts) == 0 {
		return nil
	}

	query := fmt.Sprintf(
		"UPDATE t_programming SET %s WHERE id = $%d",
		strings.Join(parts, ", "),
		len(args)+1,
	)
	args = append(args, id)

	_, err := r.conn.ExecContext(ctx, query, args...)
	if err != nil {
		slog.Error("Error executing update", "query", query, "args", args, "error", err)
		return ierrors.ErrInternal
	}

	return nil
}
