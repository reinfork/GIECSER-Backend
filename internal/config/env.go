package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	PGHost     string
	PGPort     string
	PGUser     string
	PGPassword string
	PGDatabase string
	JWTSecret  string
	GinMode    string
	ServerPort string
}

func (c *Config) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=require connect_timeout=5",
		c.PGHost, c.PGPort, c.PGUser, c.PGPassword, c.PGDatabase,
	)
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
		if os.Getenv("PGHOST") != "" {
			loaded = true
		}
	}
	if !loaded && os.Getenv("PGHOST") == "" {
		log.Println("no .env file found, relying on environment variables")
	}

	cfg := &Config{
		PGHost:     os.Getenv("PGHOST"),
		PGPort:     os.Getenv("PGPORT"),
		PGUser:     os.Getenv("PGUSER"),
		PGPassword: os.Getenv("PGPASSWORD"),
		PGDatabase: os.Getenv("PGDATABASE"),
		JWTSecret:  os.Getenv("JWT_SECRET"),
		GinMode:    os.Getenv("GIN_MODE"),
		ServerPort: os.Getenv("PORT"),
	}
	if cfg.JWTSecret == "" {
		cfg.JWTSecret = "dev-secret-change-me"
	}
	if cfg.GinMode == "" {
		cfg.GinMode = "debug"
	}
	if cfg.ServerPort == "" {
		cfg.ServerPort = "8080"
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
