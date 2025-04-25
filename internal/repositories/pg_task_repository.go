package repositories

import (
	"CourseService/internal/repositories/models"
	ierrors "CourseService/pkg/errors"
	"database/sql"
	"errors"
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
