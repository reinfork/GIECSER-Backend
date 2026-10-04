package router

import (
	"time"

	"asri-backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func registerAuthRoutes(r *gin.Engine, h *handler.AuthHandler, rl func(int, time.Duration) gin.HandlerFunc) {
	r.POST("/teachers/login", rl(10, time.Minute), h.TeacherLogin)
}
