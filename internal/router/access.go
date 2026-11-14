package router

import (
	"time"

	"asri-backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func registerAccessRoutes(r *gin.Engine, h *handler.AccessHandler, auth, teacher gin.HandlerFunc, rl func(int, time.Duration) gin.HandlerFunc) {
	// NOTE: caller passes the strict limiter — PIN guesses stay throttled
	// even when Redis is down.
	r.POST("/access/validate", rl(10, time.Minute), h.ValidatePIN)
	codes := r.Group("/access-codes")
	{
		codes.POST("", auth, teacher, h.IssuePIN)
		codes.GET("", auth, teacher, h.ListPINs)
	}
}
