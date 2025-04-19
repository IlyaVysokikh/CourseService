package main

import (
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

func configureLogger(level string) *slog.Logger {
	var handler slog.Handler

	switch level {
	case "debug":
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	case "info":
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
	case "warn":
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelWarn,
		})
	case "error":
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelError,
		})
	default:
		handler =slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
	}

	logger := slog.New(handler)
	slog.SetDefault(logger)

	return logger
}

func configureRoutes(router *gin.Engine, h *rest.Handler) {
	router.GET("/health", h.HealthCheck)
	router.GET("/courses", h.GetAllCoursesHandler)
	router.GET("/courses/:id", h.GetCourseHandler)
	router.POST("/courses", h.CreateCourseHandler)
	router.POST("/courses/:id/clone", h.CloneCourseHandler)
	router.POST("/courses/:id/modules", h.CreateModulesHandler)
	router.GET("/modules/:id", h.GetModuleHandler)
}