package handler

import (
	"context"
	"net/http"
	"time"

	"asri-backend/internal/store"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HealthHandler struct {
	db    *gorm.DB
	redis *store.Redis
}

func NewHealthHandler(db *gorm.DB, redis *store.Redis) *HealthHandler {
	return &HealthHandler{db: db, redis: redis}
}

func (h *HealthHandler) Check(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()

	// ponytail: sequential pings, one shared timeout — fan out when health latency matters.
	dbStatus, redisStatus := "ok", "ok"
	if sqlDB, err := h.db.DB(); err != nil {
		dbStatus = err.Error()
	} else if err := sqlDB.PingContext(ctx); err != nil {
		dbStatus = err.Error()
	}
	if err := h.redis.Ping(ctx); err != nil {
		redisStatus = err.Error()
	}

	status, code := "ok", http.StatusOK
	if dbStatus != "ok" || redisStatus != "ok" {
		status, code = "degraded", http.StatusServiceUnavailable
	}
	c.JSON(code, gin.H{"status": status, "db": dbStatus, "redis": redisStatus})
}
