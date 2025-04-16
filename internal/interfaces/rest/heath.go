package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) HealthCheck(ctx *gin.Context) {
    ctx.String(http.StatusOK, "OK")
}