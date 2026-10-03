package main

import (
	"log"

	"asri-backend/internal/config"

	"asri-backend/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	gin.SetMode(cfg.GinMode)

	if cfg.PGHost == "" || cfg.PGPort == "" || cfg.PGUser == "" || cfg.PGDatabase == "" {
		log.Fatalf("missing database env vars (PGHOST, PGPORT, PGUSER, PGDATABASE)")
	}

	db, err := config.InitDB(cfg.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	r := router.New(db)

	addr := ":" + cfg.ServerPort
	if err := r.Run(addr); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
