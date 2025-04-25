package postgresql

import (
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

func NewPostgresConnection(connectionString string) (*sqlx.DB, error) {
	db, err := sqlx.Open("pgx", connectionString)
	if err != nil {
		slog.Error("failed to connect to database", "error", err)
		return nil, err
	}

	slog.Info("Successfully connected to the database")

	return db, nil
}
