package router

import (
	"asri-backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func registerModuleRoutes(r *gin.Engine, h *handler.ModuleHandler, auth, teacher gin.HandlerFunc) {
	courses := r.Group("/courses")
	{
		courses.GET("/:id/modules", auth, h.ListByCourse)
	}
	modules := r.Group("/modules")
	{
		modules.GET("/:id", auth, h.GetByID)
		modules.POST("", auth, teacher, h.Create)
		modules.PUT("/:id", auth, teacher, h.Update)
		modules.DELETE("/:id", auth, teacher, h.Delete)
	}
}
