package main

import (
	"context"
	"log"
	"strconv"

	"asri-backend/internal/config"
	"asri-backend/internal/repository"
	"asri-backend/internal/router"
	"asri-backend/internal/service"
	"asri-backend/internal/store"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	gin.SetMode(cfg.GinMode)

	if cfg.DatabaseURL == "" {
		log.Fatalf("missing DATABASE_URL env var (postgres://user:pass@host:port/db?sslmode=require)")
	}

	db, err := config.InitDB(cfg.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	sessionTTL, _ := strconv.Atoi(cfg.SessionTTLHours)
	redisStore := store.NewRedis(cfg.RedisURL)
	if err := redisStore.Ping(context.Background()); err != nil {

		log.Printf("warning: redis unreachable (%v) — PIN validation will fail closed", err)
	}

	if cfg.GroqAPIKey == "" {
		log.Println("warning: GROQ_API_KEY unset — submit-audio will fail until configured")
	}

	// First-boot teacher provisioning (env-driven, idempotent).
	authSvc := service.NewAuthService(repository.NewUserRepository(db))
	if u, err := authSvc.SeedTeacher(cfg.TeacherEmail, cfg.TeacherPassword); err != nil {
		log.Fatalf("failed to seed teacher: %v", err)
	} else if u != nil {
		log.Printf("teacher account ready: %s", u.Email)
	}

	r := router.New(db, router.Config{
		RedisURL:        cfg.RedisURL,
		SessionTTLHours: sessionTTL,
		GroqAPIKey:      cfg.GroqAPIKey,
	})

	addr := ":" + cfg.ServerPort
	if err := r.Run(addr); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
