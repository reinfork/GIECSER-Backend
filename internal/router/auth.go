package router

import (
	"asri-backend/internal/handler"
	"asri-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func registerAuthRoutes(rg *gin.RouterGroup, h *handler.AuthHandler) {
	auth := rg.Group("/auth")
	{
		auth.POST("/register", h.Register)
		auth.POST("/login", h.Login)
		// Current user profile — requires JWT
		auth.GET("/me", middleware.AuthMiddleware(), h.Me)
	}

	users := rg.Group("/users")
	{
		// Public read for now, admin-only delete
		users.GET("", h.ListUsers)
		users.GET("/:id", h.GetUserByID)
		users.DELETE("/:id", middleware.AuthMiddleware(), middleware.RequireRole("admin"), h.DeleteUser)
	}
}
