package router

import (
	"asri-backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func registerChapterRoutes(r *gin.Engine, h *handler.ChapterHandler, auth, teacher gin.HandlerFunc) {
	chapters := r.Group("/chapters")
	{
		chapters.GET("", auth, h.List)
		chapters.GET("/:id", auth, h.GetByID)
		chapters.POST("", auth, teacher, h.Create)
		chapters.PUT("/:id", auth, teacher, h.Update)
		chapters.DELETE("/:id", auth, teacher, h.Delete)
	}
}
