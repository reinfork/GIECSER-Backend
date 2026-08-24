package router

import (
	"asri-backend/internal/handler"
	"asri-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func registerCourseRoutes(rg *gin.RouterGroup, h *handler.CourseHandler) {
	courses := rg.Group("/courses")
	{
		courses.GET("", h.List)
		courses.GET("/:id", h.GetByID)
		// Protected — requires JWT, admin only for writes (like Express auth middleware)
		courses.POST("", middleware.AuthMiddleware(), middleware.RequireRole("admin"), h.Create)
		courses.PUT("/:id", middleware.AuthMiddleware(), middleware.RequireRole("admin"), h.Update)
		courses.DELETE("/:id", middleware.AuthMiddleware(), middleware.RequireRole("admin"), h.Delete)
	}
}
