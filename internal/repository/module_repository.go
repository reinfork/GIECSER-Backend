package repository

import (
	"asri-backend/internal/model"

	"gorm.io/gorm"
)

type ModuleRepository struct {
	db *gorm.DB
}

func NewModuleRepository(db *gorm.DB) *ModuleRepository {
	return &ModuleRepository{db: db}
}

func (r *ModuleRepository) FindByCourse(courseID string) ([]model.Module, error) {
	var modules []model.Module
	if err := r.db.Where("course_id = ?", courseID).Order("order_index ASC").Find(&modules).Error; err != nil {
		return nil, err
	}
	return modules, nil
}

func (r *ModuleRepository) FindByID(id string) (*model.Module, error) {
	var module model.Module
	if err := r.db.First(&module, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &module, nil
}

func (r *ModuleRepository) Create(module *model.Module) error {
	return r.db.Create(module).Error
}

func (r *ModuleRepository) Update(module *model.Module) error {
	return r.db.Save(module).Error
}

func (r *ModuleRepository) Delete(id string) error {
	return r.db.Delete(&model.Module{}, "id = ?", id).Error
}
