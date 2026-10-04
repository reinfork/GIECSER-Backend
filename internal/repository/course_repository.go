package repository

import (
	"asri-backend/internal/model"

	"gorm.io/gorm"
)

type CourseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) *CourseRepository {
	return &CourseRepository{db: db}
}

func (r *CourseRepository) FindByChapter(chapterID string) ([]model.Course, error) {
	var courses []model.Course
	if err := r.db.Where("chapter_id = ?", chapterID).Order("order_index ASC").Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}

func (r *CourseRepository) FindByID(id string) (*model.Course, error) {
	var course model.Course
	if err := r.db.First(&course, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &course, nil
}

func (r *CourseRepository) Create(course *model.Course) error {
	return r.db.Create(course).Error
}

func (r *CourseRepository) Update(course *model.Course) error {
	return r.db.Save(course).Error
}

func (r *CourseRepository) Delete(id string) error {
	return r.db.Delete(&model.Course{}, "id = ?", id).Error
}
