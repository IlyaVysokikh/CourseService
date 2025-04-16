package pkg

import (
	"log/slog"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

type Config struct {
	Env      string      `env:"ENV" env-default:"local"`
	Logger   string      `env:"LOGGER_LEVEL" env-default:"debug"`
	APP_PORT string      `env:"APP_PORT" env-default:"8080"`
	DB_CONNECTION string `env:"DB_CONNECTION" env-default:"postgres://user:password@localhost:5432/dbname"`
}

func MustLoad() (*Config, error) {
	_ = godotenv.Load()

	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		slog.Error("Failed to read config from environment variables", "error", err)
		return nil, errors.Wrap(err, "cannot read config from environment variables")
	}

	return &cfg, nil
}