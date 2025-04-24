package main

import (
	"io"
	"log"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"

	"CourseService/internal/interfaces/rest"
	"CourseService/internal/repositories"
	"CourseService/internal/services"
	"CourseService/internal/usecase"
	cfg "CourseService/pkg"
)

func main() {
	cfg, err := cfg.MustLoad()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	logger := configureLogger(cfg.Logger)

	conn, err := repositories.NewPostgresConnection(cfg.DB_CONNECTION)
	if err != nil {
		logger.Error("Failed to connect to database", "error", err)
		panic("Failed to connect to database")
	}

	repositories := repositories.NewRepository(conn)

	services := services.NewService(repositories)

	usecase := usecase.NewUsecase(services)

	h := rest.NewHandler(usecase)

	router := gin.Default()

	configureRoutes(router, h)

	if err := router.Run(":" + cfg.APP_PORT); err != nil {
		logger.Error("Failed to start server", "error", err)
		return
	}
	logger.Info("Application started successfully", "port", cfg.APP_PORT)
}

func configureLogger(cfg cfg.LoggerConfig) *slog.Logger {
	var writers []io.Writer

	writers = append(writers, os.Stdout)

	if cfg.FileOutput != "" {
		file, err := os.OpenFile(cfg.FileOutput, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic("failed to open log file: " + err.Error())
		}
		writers = append(writers, file)
	}

	multiWriter := io.MultiWriter(writers...)

	opts := &slog.HandlerOptions{
		AddSource: true,
	}

	switch cfg.Level {
	case "debug":
		opts.Level = slog.LevelDebug
	case "info":
		opts.Level = slog.LevelInfo
		opts.AddSource = false
	case "warn":
		opts.Level = slog.LevelWarn
	case "error":
		opts.Level = slog.LevelError
	default:
		opts.Level = slog.LevelInfo
	}

	var handler slog.Handler
	switch cfg.Format {
	case "json":
		handler = slog.NewJSONHandler(multiWriter, opts)
	default:
		handler = slog.NewTextHandler(multiWriter, opts)
	}

	logger := slog.New(handler)
	slog.SetDefault(logger)

	return logger
}

func configureRoutes(router *gin.Engine, h *rest.Handler) {
	router.GET("/health", h.HealthCheck)

	router.GET("/courses", h.GetAllCoursesHandler)
	router.GET("/courses/:id", h.GetCourseHandler)
	router.DELETE("/courses/:id", h.DeleteCourseHandler)
	router.POST("/courses", h.CreateCourseHandler)
	router.PATCH("/courses/:id", h.UpdateCourseHandler)
	router.POST("/courses/:id/clone", h.CloneCourseHandler)

	router.POST("/courses/:id/modules", h.CreateModulesHandler)
	router.GET("/modules/:id", h.GetModuleHandler)

	router.GET("/tasks/:id", h.GetTaskHandler)
}
