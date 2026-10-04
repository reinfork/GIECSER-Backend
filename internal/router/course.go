package router

import (
	"asri-backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func registerCourseRoutes(r *gin.Engine, h *handler.CourseHandler, auth, teacher gin.HandlerFunc) {
	chapters := r.Group("/chapters")
	{
		chapters.GET("/:id/courses", auth, h.ListByChapter)
	}
	courses := r.Group("/courses")
	{
		courses.GET("/:id", auth, h.GetByID)
		courses.POST("", auth, teacher, h.Create)
		courses.PUT("/:id", auth, teacher, h.Update)
		courses.DELETE("/:id", auth, teacher, h.Delete)
	}
}
