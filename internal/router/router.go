package router

import (
	"asri-backend/internal/handler"
	"asri-backend/internal/middleware"
	"asri-backend/internal/repository"
	"asri-backend/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func New(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
	}))
	r.Use(middleware.CORSManual())

	registerRoutes(r, db)

	return r
}

func registerRoutes(r *gin.Engine, db *gorm.DB) {
	// Repositories
	courseRepo := repository.NewCourseRepository(db)
	lessonRepo := repository.NewLessonRepository(db)
	userRepo := repository.NewUserRepository(db)

	// Services
	courseService := service.NewCourseService(courseRepo)
	lessonService := service.NewLessonService(lessonRepo, courseRepo)
	authService := service.NewAuthService(userRepo)

	// Handlers
	healthHandler := handler.NewHealthHandler()
	courseHandler := handler.NewCourseHandler(courseService)
	lessonHandler := handler.NewLessonHandler(lessonService)
	authHandler := handler.NewAuthHandler(authService)

	registerHealthRoutes(r, healthHandler)
	registerCourseRoutes(r, courseHandler)
	registerLessonRoutes(r, lessonHandler)
	registerAuthRoutes(r, authHandler)
}