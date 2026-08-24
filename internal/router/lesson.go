package router

import (
	"asri-backend/internal/handler"
	"asri-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func registerLessonRoutes(rg *gin.RouterGroup, h *handler.LessonHandler) {
	lessons := rg.Group("/lessons")
	{
		lessons.GET("/:id", h.GetByID)
		lessons.POST("", middleware.AuthMiddleware(), middleware.RequireRole("admin"), h.Create)
		lessons.PUT("/:id", middleware.AuthMiddleware(), middleware.RequireRole("admin"), h.Update)
		lessons.DELETE("/:id", middleware.AuthMiddleware(), middleware.RequireRole("admin"), h.Delete)
		lessons.POST("/:id/practice", middleware.AuthMiddleware(), h.Practice)
	}

	// Nested: /api/v1/courses/:id/lessons — mirrors typical REST for Asri
	courses := rg.Group("/courses")
	{
		courses.GET("/:id/lessons", h.ListByCourse)
	}
}
