package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL     string
	RedisURL        string
	JWTSecret       string
	GinMode         string
	ServerPort      string
	GroqAPIKey      string
	SessionTTLHours string
	TeacherEmail    string
	TeacherPassword string
}

func (c *Config) DSN() string {
	return c.DatabaseURL
}

func Load() *Config {
	loaded := false
	for _, p := range findEnvPaths() {
		if err := godotenv.Load(p); err == nil {
			loaded = true
			break
		}
	}
	if !loaded {
		_ = godotenv.Load()
		if os.Getenv("DATABASE_URL") != "" {
			loaded = true
		}
	}
	if !loaded && os.Getenv("DATABASE_URL") == "" {
		log.Println("no .env file found, relying on environment variables")
	}

	cfg := &Config{
		DatabaseURL:     os.Getenv("DATABASE_URL"),
		RedisURL:        os.Getenv("REDIS_URL"),
		JWTSecret:       os.Getenv("JWT_SECRET"),
		GinMode:         os.Getenv("GIN_MODE"),
		ServerPort:      os.Getenv("PORT"),
		GroqAPIKey:      os.Getenv("GROQ_API_KEY"),
		SessionTTLHours: os.Getenv("SESSION_TTL_HOURS"),
		TeacherEmail:    os.Getenv("TEACHER_EMAIL"),
		TeacherPassword: os.Getenv("TEACHER_PASSWORD"),
	}
	if cfg.JWTSecret == "" {
		cfg.JWTSecret = "dev-secret-change-me"
	}
	if cfg.GinMode == "" {
		cfg.GinMode = "debug"
	}
	if cfg.ServerPort == "" {
		cfg.ServerPort = "5000"
	}
	if cfg.RedisURL == "" {
		cfg.RedisURL = "redis://localhost:6379/0"
	}
	if cfg.SessionTTLHours == "" {
		cfg.SessionTTLHours = "12"
	}
	return cfg
}

func findEnvPaths() []string {
	var paths []string

	// 1. Current working directory and parents
	if wd, err := os.Getwd(); err == nil {
		dir := wd
		for i := 0; i < 4; i++ {
			paths = append(paths, filepath.Join(dir, ".env"))
			dir = filepath.Dir(dir)
		}
	}

	if exe, err := os.Executable(); err == nil {
		dir := filepath.Dir(exe)
		for i := 0; i < 4; i++ {
			paths = append(paths, filepath.Join(dir, ".env"))
			// Also check project backend path
			paths = append(paths, filepath.Join(dir, "backend", ".env"))
			dir = filepath.Dir(dir)
		}
		paths = append(paths, "/home/dwipapap/Projek/asri/backend/.env")
	}

	seen := make(map[string]bool)
	uniq := make([]string, 0, len(paths))
	for _, p := range paths {
		if !seen[p] {
			seen[p] = true
			uniq = append(uniq, p)
		}
	}
	return uniq
}
