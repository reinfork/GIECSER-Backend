package router

import (
	"time"

	"asri-backend/internal/groq"
	"asri-backend/internal/handler"
	"asri-backend/internal/middleware"
	"asri-backend/internal/repository"
	"asri-backend/internal/service"
	"asri-backend/internal/store"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Config carries runtime wiring main assembles once.
type Config struct {
	RedisURL        string
	SessionTTLHours int
	GroqAPIKey      string
}

// New returns a configured gin.Engine with all routes registered.
// ponytail: routes intentionally unversioned — single frontend, lockstep deploys.
func New(db *gorm.DB, cfg Config) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
	}))
	r.Use(middleware.CORSManual())

	registerRoutes(r, db, cfg)

	return r
}

func registerRoutes(r *gin.Engine, db *gorm.DB, cfg Config) {
	redisStore := store.NewRedis(cfg.RedisURL)

	// Repositories
	userRepo := repository.NewUserRepository(db)
	codeRepo := repository.NewAccessCodeRepository(db)
	chapterRepo := repository.NewChapterRepository(db)
	courseRepo := repository.NewCourseRepository(db)
	moduleRepo := repository.NewModuleRepository(db)
	taskRepo := repository.NewTaskRepository(db)
	submissionRepo := repository.NewSubmissionRepository(db)

	// Services
	authService := service.NewAuthService(userRepo)
	accessService := service.NewAccessService(codeRepo, chapterRepo, redisStore, cfg.SessionTTLHours)
	chapterService := service.NewChapterService(chapterRepo)
	courseService := service.NewCourseService(courseRepo, chapterRepo)
	moduleService := service.NewModuleService(moduleRepo, courseRepo)
	taskService := service.NewTaskService(taskRepo, moduleRepo, courseRepo, chapterRepo)
	submissionService := service.NewSubmissionService(
		taskRepo, moduleRepo, submissionRepo,
		groq.NewWhisperClient(cfg.GroqAPIKey), groq.NewFeedbackClient(cfg.GroqAPIKey),
	)

	// Handlers
	healthHandler := handler.NewHealthHandler()
	authHandler := handler.NewAuthHandler(authService)
	accessHandler := handler.NewAccessHandler(accessService)
	chapterHandler := handler.NewChapterHandler(chapterService)
	courseHandler := handler.NewCourseHandler(courseService)
	moduleHandler := handler.NewModuleHandler(moduleService)
	taskHandler := handler.NewTaskHandler(taskService, submissionService)

	rl := func(limit int, window time.Duration) gin.HandlerFunc {
		return middleware.RateLimit(redisStore, limit, window)
	}
	auth := middleware.AuthMiddleware()
	teacher := middleware.RequireKind(middleware.KindTeacher)

	registerHealthRoutes(r, healthHandler)
	registerAuthRoutes(r, authHandler, rl)
	registerAccessRoutes(r, accessHandler, auth, teacher, rl)
	registerChapterRoutes(r, chapterHandler, auth, teacher)
	registerCourseRoutes(r, courseHandler, auth, teacher)
	registerModuleRoutes(r, moduleHandler, auth, teacher)
	registerTaskRoutes(r, taskHandler, auth, teacher, rl)
}
