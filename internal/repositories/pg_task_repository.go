package repositories

import (
	"CourseService/internal/repositories/models"
	"log/slog"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type TaskRepositoryImpl struct {
	conn *sqlx.DB
}

func NewTaskRepository (conn *sqlx.DB) TaskRepository {
	return &TaskRepositoryImpl{
		conn: conn,
	}
}

func (t *TaskRepositoryImpl) GetTasks(moduleId uuid.UUID) ([]models.Task, error) {
	query := `SELECT * FROM t_task WHERE id_module = $1`
	tasks := []models.Task{}
	err := t.conn.Select(&tasks, query, moduleId)
	if err != nil {
		slog.Error("Error executing select", "query", query, "error", err)
		return nil, err
	}

	return tasks, nil
}

func (t *TaskRepositoryImpl) GetTasksByModule(moduleId uuid.UUID) ([]models.Task, error) {
	query := `SELECT * FROM t_task WHERE id_module = $1`
	tasks := []models.Task{}
	err := t.conn.Select(&tasks, query, moduleId)
	if err != nil {
		slog.Error("Error executing select", "query", query, "error", err)
		return nil, err
	}

	return tasks, nil
}