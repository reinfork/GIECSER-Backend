package repository

import (
	"asri-backend/internal/model"

	"gorm.io/gorm"
)

type SubmissionRepository struct {
	db *gorm.DB
}

func NewSubmissionRepository(db *gorm.DB) *SubmissionRepository {
	return &SubmissionRepository{db: db}
}

func (r *SubmissionRepository) Create(s *model.AudioSubmission) error {
	return r.db.Create(s).Error
}
