package rest

import (
	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
	BaseHandler
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{
		BaseHandler: BaseHandler{},
	}
}

func (h *HealthHandler) HealthCheck(ctx *gin.Context) {
	h.ok(ctx, "ok")
}
