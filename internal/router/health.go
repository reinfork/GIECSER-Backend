package router

import (
	"asri-backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func registerHealthRoutes(r gin.IRouter, h *handler.HealthHandler) {
	r.GET("/health", h.Check)
}
