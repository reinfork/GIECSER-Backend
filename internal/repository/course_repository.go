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

func (r *CourseRepository) FindAll() ([]model.Course, error) {
	var courses []model.Course
	if err := r.db.Preload("Lessons", func(db *gorm.DB) *gorm.DB {
		return db.Order("order_index ASC")
	}).Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}

func (r *CourseRepository) FindAllPaginated(limit, offset int) ([]model.Course, int64, error) {
	var courses []model.Course
	var total int64
	if err := r.db.Model(&model.Course{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	q := r.db.Preload("Lessons", func(db *gorm.DB) *gorm.DB {
		return db.Order("order_index ASC")
	}).Order("id ASC").Limit(limit).Offset(offset)
	if err := q.Find(&courses).Error; err != nil {
		return nil, 0, err
	}
	return courses, total, nil
}

func (r *CourseRepository) FindByID(id uint) (*model.Course, error) {
	var course model.Course
	if err := r.db.Preload("Lessons", func(db *gorm.DB) *gorm.DB {
		return db.Order("order_index ASC")
	}).First(&course, id).Error; err != nil {
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

func (r *CourseRepository) Delete(id uint) error {
	return r.db.Delete(&model.Course{}, id).Error
}
