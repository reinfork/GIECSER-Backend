package router

import (
	"asri-backend/internal/handler"
	"asri-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func registerLessonRoutes(rg *gin.Engine, h *handler.LessonHandler) {
	lessons := rg.Group("/lessons")
	{
		lessons.GET("/:id", h.GetByID)
		lessons.POST("", middleware.AuthMiddleware(), middleware.RequireRole("admin"), h.Create)
		lessons.PUT("/:id", middleware.AuthMiddleware(), middleware.RequireRole("admin"), h.Update)
		lessons.DELETE("/:id", middleware.AuthMiddleware(), middleware.RequireRole("admin"), h.Delete)
		lessons.POST("/:id/practice", middleware.AuthMiddleware(), h.Practice)
	}

	courses := rg.Group("/courses")
	{
		courses.GET("/:id/lessons", h.ListByCourse)
	}
}
