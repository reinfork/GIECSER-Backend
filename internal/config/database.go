package config

import (
	"asri-backend/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(
		&model.User{},
		&model.AccessCode{},
		&model.Chapter{},
		&model.Course{},
		&model.Module{},
		&model.Task{},
		&model.AudioSubmission{},
	); err != nil {
		return nil, err
	}

	return db, nil
}
