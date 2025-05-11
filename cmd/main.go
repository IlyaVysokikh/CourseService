package main

import (
	"CourseService/pkg/postgresql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

// @title Example API
// @version 1.0
// @description This is an example API to demonstrate Swagger documentation generation
// @termsOfService http://example.com/terms

// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email support@example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

func main() {
	c, err := cfg.MustLoad()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	logger := configureLogger(c.Logger)

	conn, err := postgresql.NewPostgresConnection(c.DB_CONNECTION)
	if err != nil {
		logger.Error("Failed to connect to database", "error", err)
		panic("Failed to connect to database")
	}

	repos := repositories.NewRepository(conn)

	svc := services.NewService(repos)

	useCases := usecase.NewUseCase(svc)

	h := rest.NewHandler(useCases)

	router := gin.Default()

	configureRoutes(router, h)

	if err := router.Run(":" + c.APP_PORT); err != nil {
		logger.Error("Failed to start server", "error", err)
		return
	}
	logger.Info("Application started successfully", "port", c.APP_PORT)
}

func configureLogger(c cfg.LoggerConfig) *slog.Logger {
	var writers []io.Writer

	writers = append(writers, os.Stdout)

	if c.FileOutput != "" {
		file, err := os.OpenFile(c.FileOutput, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic("failed to open log file: " + err.Error())
		}
		writers = append(writers, file)
	}

	multiWriter := io.MultiWriter(writers...)

	opts := &slog.HandlerOptions{
		AddSource: true,
	}

	switch c.Level {
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
	switch c.Format {
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
	coursesRouter := router.Group("/courses")
	{
		coursesRouter.GET("", h.CoursesHandler.GetAllCoursesHandler)
		coursesRouter.GET("/:id", h.CoursesHandler.GetCourseHandler)
		coursesRouter.DELETE("/:id", h.CoursesHandler.DeleteCourseHandler)
		coursesRouter.POST("", h.CoursesHandler.CreateCourseHandler)
		coursesRouter.PATCH("/:id", h.CoursesHandler.UpdateCourseHandler)
		coursesRouter.POST("/:id/clone", h.CoursesHandler.CloneCourseHandler)
	}

	healthCheckRouter := router.Group("/health")
	{
		healthCheckRouter.GET("", h.HealthHandler.HealthCheck)
	}

	modulesRouter := router.Group("/modules")
	{
		modulesRouter.POST("", h.ModulesHandler.CreateModulesHandler)
		modulesRouter.GET("/:id", h.ModulesHandler.GetModuleHandler)
		modulesRouter.DELETE("/:id", h.ModulesHandler.DeleteModuleHandler)
		modulesRouter.POST("/:id", h.ModulesHandler.CreateAttachment)
	}

	tasksRouter := router.Group("/tasks")
	{
		tasksRouter.GET("/:id", h.TasksHandler.GetTaskHandler)
		tasksRouter.DELETE("/:id", h.TasksHandler.DeleteTaskHandler)
		tasksRouter.PATCH("/:id")
		tasksRouter.POST("", h.TasksHandler.CreateTask)
		tasksRouter.GET("/:id/test-data", h.TestDataHandler.GetAllTestData)
	}

	testDataRouter := router.Group("/programming")
	{
		testDataRouter.GET("/:id", h.TestDataHandler.GetTestData)
		testDataRouter.DELETE("/:id", h.TestDataHandler.DeleteTestData)
		testDataRouter.PATCH("/:id", h.TestDataHandler.UpdateTestData)
		testDataRouter.POST("", h.TestDataHandler.CreateTestData)
	}
	swaggerRouter := router.Group("")
	{
		swaggerRouter.StaticFile("/doc.json", "./docs/swagger.json")
		url := ginSwagger.URL("http://localhost:8080/doc.json")
		swaggerRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}
}
