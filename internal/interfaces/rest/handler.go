package rest

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	CoursesHandler  *CoursesHandler
	ModulesHandler  *ModulesHandler
	TasksHandler    *TasksHandler
	HealthHandler   *HealthHandler
	TestDataHandler *TestDataHandler
}

func NewHandler(useCases *usecase.UseCase) *Handler {
	return &Handler{
		CoursesHandler:  NewCoursesHandler(useCases),
		ModulesHandler:  NewModulesHandler(useCases),
		TasksHandler:    NewTasksHandler(useCases),
		HealthHandler:   NewHealthHandler(),
		TestDataHandler: NewTestDataHandler(useCases),
	}
}

type BaseHandler struct{}

func (h *BaseHandler) badRequest(ctx *gin.Context, err error) {
	res := dto.ErrorResponse{Message: err.Error()}
	ctx.JSON(http.StatusBadRequest, res)
}

func (h *BaseHandler) notFound(ctx *gin.Context, err error) {
	res := dto.ErrorResponse{Message: err.Error()}
	ctx.JSON(http.StatusNotFound, res)
}

func (h *BaseHandler) ok(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, data)
}

func (h *BaseHandler) created(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusCreated, data)
}

func (h *BaseHandler) internalServerError(ctx *gin.Context, err error) {
	res := dto.ErrorResponse{Message: err.Error()}
	ctx.JSON(http.StatusInternalServerError, res)
}

func (h *BaseHandler) noContent(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}
