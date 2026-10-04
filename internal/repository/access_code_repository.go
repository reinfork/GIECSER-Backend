package repository

import (
	"time"

	"asri-backend/internal/model"

	"gorm.io/gorm"
)

type AccessCodeRepository struct {
	db *gorm.DB
}

func NewAccessCodeRepository(db *gorm.DB) *AccessCodeRepository {
	return &AccessCodeRepository{db: db}
}

func (r *AccessCodeRepository) Create(code *model.AccessCode) error {
	return r.db.Create(code).Error
}

func (r *AccessCodeRepository) FindByCode(code string) (*model.AccessCode, error) {
	var ac model.AccessCode
	if err := r.db.Where("code = ?", code).First(&ac).Error; err != nil {
		return nil, err
	}
	return &ac, nil
}

func (r *AccessCodeRepository) ListActive(now time.Time) ([]model.AccessCode, error) {
	var codes []model.AccessCode
	if err := r.db.Where("expires_at > ?", now).Order("created_at DESC").Find(&codes).Error; err != nil {
		return nil, err
	}
	return codes, nil
}
