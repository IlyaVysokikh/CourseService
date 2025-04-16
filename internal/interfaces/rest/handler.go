package rest

import (
	"CourseService/internal/usecase"
	"github.com/gin-gonic/gin"
)

type Handler struct{
	usecases *usecase.Usecase
	
}
func NewHandler(usecases *usecase.Usecase) *Handler {
	return &Handler{
		usecases: usecases,
	}
}


func (h *Handler) badRequest(ctx *gin.Context, err error) {
	ctx.JSON(400, gin.H{"error": err.Error()})
}

func (h *Handler) notFound(ctx *gin.Context, err error) {
	ctx.JSON(404, gin.H{"error": err.Error()})
}

func (h *Handler) ok(ctx *gin.Context, data interface{}) {
	ctx.JSON(200, data)
}