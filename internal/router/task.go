package router

import (
	"time"

	"asri-backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func registerTaskRoutes(r *gin.Engine, h *handler.TaskHandler, auth, teacher gin.HandlerFunc, rl func(int, time.Duration) gin.HandlerFunc) {
	modules := r.Group("/modules")
	{
		modules.GET("/:id/tasks", auth, h.ListByModule)
	}
	tasks := r.Group("/tasks")
	{
		tasks.GET("/:id", auth, h.GetByID)
		tasks.POST("/:id/submit-audio", auth, rl(20, time.Minute), h.SubmitAudio)
		tasks.POST("/:id/submit-text", auth, rl(20, time.Minute), h.SubmitText)
		tasks.POST("", auth, teacher, h.Create)
		tasks.PUT("/:id", auth, teacher, h.Update)
		tasks.DELETE("/:id", auth, teacher, h.Delete)
	}
}
