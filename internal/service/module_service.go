package service

import (
	"errors"

	"asri-backend/internal/model"

	"gorm.io/gorm"
)

type ModuleRepo interface {
	FindByCourse(courseID string) ([]model.Module, error)
	FindByID(id string) (*model.Module, error)
	Create(m *model.Module) error
	Update(m *model.Module) error
	Delete(id string) error
}

type ModuleService struct {
	repo    ModuleRepo
	courses CourseRepo
}

func NewModuleService(repo ModuleRepo, courses CourseRepo) *ModuleService {
	return &ModuleService{repo: repo, courses: courses}
}

func (s *ModuleService) ListByCourse(courseID string) ([]model.Module, error) {
	return s.repo.FindByCourse(courseID)
}

func (s *ModuleService) GetByID(id string) (*model.Module, error) {
	module, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return module, nil
}

func (s *ModuleService) Create(input model.CreateModuleInput) (*model.Module, error) {
	if _, err := s.courses.FindByID(input.CourseID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	module := &model.Module{
		CourseID: input.CourseID, Title: input.Title, Type: input.Type,
		ContentText: input.ContentText, MediaURL: input.MediaURL,
		TargetTranscript: input.TargetTranscript, AudioModelURL: input.AudioModelURL,
		OrderIndex: input.OrderIndex,
	}
	if err := s.repo.Create(module); err != nil {
		return nil, err
	}
	return module, nil
}

func (s *ModuleService) Update(id string, input model.UpdateModuleInput) (*model.Module, error) {
	module, err := s.GetByID(id)
	if err != nil {
		return nil, err
	}
	if input.Title != nil {
		module.Title = *input.Title
	}
	if input.Type != nil {
		module.Type = *input.Type
	}
	if input.ContentText != nil {
		module.ContentText = *input.ContentText
	}
	if input.MediaURL != nil {
		module.MediaURL = *input.MediaURL
	}
	if input.TargetTranscript != nil {
		module.TargetTranscript = *input.TargetTranscript
	}
	if input.AudioModelURL != nil {
		module.AudioModelURL = *input.AudioModelURL
	}
	if input.OrderIndex != nil {
		module.OrderIndex = *input.OrderIndex
	}
	if err := s.repo.Update(module); err != nil {
		return nil, err
	}
	return module, nil
}

func (s *ModuleService) Delete(id string) error {
	if _, err := s.GetByID(id); err != nil {
		return err
	}
	return s.repo.Delete(id)
}
