package repository

import (
	"asri-backend/internal/model"

	"gorm.io/gorm"
)

type LessonRepository struct {
	db *gorm.DB
}

func NewLessonRepository(db *gorm.DB) *LessonRepository {
	return &LessonRepository{db: db}
}

func (r *LessonRepository) FindByID(id uint) (*model.Lesson, error) {
	var lesson model.Lesson
	if err := r.db.First(&lesson, id).Error; err != nil {
		return nil, err
	}
	return &lesson, nil
}

func (r *LessonRepository) FindByCourseID(courseID uint) ([]model.Lesson, error) {
	var lessons []model.Lesson
	if err := r.db.Where("course_id = ?", courseID).Order("order_index ASC").Find(&lessons).Error; err != nil {
		return nil, err
	}
	return lessons, nil
}

func (r *LessonRepository) FindByCourseIDPaginated(courseID uint, limit, offset int) ([]model.Lesson, int64, error) {
	var lessons []model.Lesson
	var total int64
	if err := r.db.Model(&model.Lesson{}).Where("course_id = ?", courseID).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := r.db.Where("course_id = ?", courseID).Order("order_index ASC").Limit(limit).Offset(offset).Find(&lessons).Error; err != nil {
		return nil, 0, err
	}
	return lessons, total, nil
}

func (r *LessonRepository) Create(lesson *model.Lesson) error {
	return r.db.Create(lesson).Error
}

func (r *LessonRepository) Update(lesson *model.Lesson) error {
	return r.db.Save(lesson).Error
}

func (r *LessonRepository) Delete(id uint) error {
	return r.db.Delete(&model.Lesson{}, id).Error
}

func (r *LessonRepository) ExistsByCourseIDAndOrder(courseID uint, orderIndex int, excludeID *uint) (bool, error) {
	var count int64
	q := r.db.Model(&model.Lesson{}).Where("course_id = ? AND order_index = ?", courseID, orderIndex)
	if excludeID != nil {
		q = q.Where("id != ?", *excludeID)
	}
	if err := q.Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
