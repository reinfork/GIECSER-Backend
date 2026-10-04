package repository

import (
	"asri-backend/internal/model"

	"gorm.io/gorm"
)

type ChapterRepository struct {
	db *gorm.DB
}

func NewChapterRepository(db *gorm.DB) *ChapterRepository {
	return &ChapterRepository{db: db}
}

func (r *ChapterRepository) FindAll() ([]model.Chapter, error) {
	var chapters []model.Chapter
	if err := r.db.Order("order_index ASC").Find(&chapters).Error; err != nil {
		return nil, err
	}
	return chapters, nil
}

func (r *ChapterRepository) FindByID(id string) (*model.Chapter, error) {
	var chapter model.Chapter
	if err := r.db.First(&chapter, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &chapter, nil
}

func (r *ChapterRepository) Create(chapter *model.Chapter) error {
	return r.db.Create(chapter).Error
}

func (r *ChapterRepository) Update(chapter *model.Chapter) error {
	return r.db.Save(chapter).Error
}

func (r *ChapterRepository) Delete(id string) error {
	return r.db.Delete(&model.Chapter{}, "id = ?", id).Error
}
