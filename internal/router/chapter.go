package router

import (
	"time"

	"asri-backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func registerChapterRoutes(r *gin.Engine, h *handler.ChapterHandler, submissions *handler.TaskHandler, auth, teacher gin.HandlerFunc, rl func(int, time.Duration) gin.HandlerFunc) {
	chapters := r.Group("/chapters")
	{
		chapters.GET("", auth, h.List)
		chapters.GET("/:id", auth, h.GetByID)
		chapters.POST("/:id/feedback", auth, rl(10, time.Hour), submissions.ChapterFeedback)
		chapters.POST("", auth, teacher, h.Create)
		chapters.PUT("/:id", auth, teacher, h.Update)
		chapters.DELETE("/:id", auth, teacher, h.Delete)
	}
}
