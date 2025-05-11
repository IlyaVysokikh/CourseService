package repositories

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/repositories/models"
	ierrors "CourseService/pkg/errors"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type TaskRepositoryImpl struct {
	conn *sqlx.DB
}

func NewTaskRepository(conn *sqlx.DB) TaskRepository {
	return &TaskRepositoryImpl{
		conn: conn,
	}
}

func (t *TaskRepositoryImpl) GetTasks(moduleId uuid.UUID) ([]models.Task, error) {
	query := `SELECT * FROM t_task WHERE id_module = $1`
	tasks := []models.Task{}
	err := t.conn.Select(&tasks, query, moduleId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			slog.Warn("No tasks found for module", "moduleId", moduleId)
			return nil, ierrors.ErrNotFound
		}
		slog.Error("Error executing select", "query", query, "error", err)
		return nil, ierrors.ErrInternal
	}

	return tasks, nil
}

func (t *TaskRepositoryImpl) GetTasksByModule(moduleId uuid.UUID) ([]models.Task, error) {
	query := `SELECT * FROM t_task WHERE id_module = $1`
	tasks := []models.Task{}
	err := t.conn.Select(&tasks, query, moduleId)
	if err != nil {
		slog.Error("Error executing select", "query", query, "error", err)
		return nil, ierrors.ErrInternal
	}

	return tasks, nil
}

func (t *TaskRepositoryImpl) GetTask(taskId uuid.UUID) (*models.Task, error) {
	query := `SELECT * FROM t_task WHERE id = $1`
	task := models.Task{}
	err := t.conn.Get(&task, query, taskId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			slog.Warn("No task found", "taskId", taskId)
			return nil, ierrors.ErrNotFound
		}
		slog.Error("Error executing select", "query", query, "error", err)
		return nil, ierrors.ErrInternal
	}

	return &task, nil
}

func (t *TaskRepositoryImpl) DeleteTask(id uuid.UUID) error {
	const query = `DELETE FROM t_task WHERE id = $1`
	_, err := t.conn.Exec(query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ierrors.ErrNotFound
		}

		return ierrors.ErrInternal
	}

	return nil
}

func (t *TaskRepositoryImpl) Create(ctx context.Context, req dto.CreateTaskRequest) (uuid.UUID, error) {
	query := `
		INSERT INTO t_task (
			id_module, c_name, c_text, c_language, c_initial_code, 
			c_memory_limit, c_execution_timeout, c_sequence_number
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
	`

	var newTaskID uuid.UUID
	err := t.conn.QueryRowContext(ctx, query,
		req.ModuleId,
		req.Name,
		req.Text,
		req.Language,
		req.InitialCode,
		req.MemoryLimit,
		req.ExecutionTimeout,
		req.SequenceNumber,
	).Scan(&newTaskID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to insert task: %w", err)
	}

	return newTaskID, nil
}
