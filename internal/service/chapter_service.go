package service

import (
	"errors"

	"asri-backend/internal/model"

	"gorm.io/gorm"
)

type ChapterRepo interface {
	FindAll() ([]model.Chapter, error)
	FindByID(id string) (*model.Chapter, error)
	Create(c *model.Chapter) error
	Update(c *model.Chapter) error
	Delete(id string) error
}

type ChapterService struct {
	repo ChapterRepo
}

func NewChapterService(repo ChapterRepo) *ChapterService {
	return &ChapterService{repo: repo}
}

func (s *ChapterService) List() ([]model.Chapter, error) {
	return s.repo.FindAll()
}

func (s *ChapterService) GetByID(id string) (*model.Chapter, error) {
	chapter, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return chapter, nil
}

func (s *ChapterService) Create(input model.CreateChapterInput) (*model.Chapter, error) {
	chapter := &model.Chapter{Title: input.Title, Description: input.Description, OrderIndex: input.OrderIndex}
	if err := s.repo.Create(chapter); err != nil {
		return nil, err
	}
	return chapter, nil
}

func (s *ChapterService) Update(id string, input model.UpdateChapterInput) (*model.Chapter, error) {
	chapter, err := s.GetByID(id)
	if err != nil {
		return nil, err
	}
	if input.Title != nil {
		chapter.Title = *input.Title
	}
	if input.Description != nil {
		chapter.Description = *input.Description
	}
	if input.OrderIndex != nil {
		chapter.OrderIndex = *input.OrderIndex
	}
	if err := s.repo.Update(chapter); err != nil {
		return nil, err
	}
	return chapter, nil
}

func (s *ChapterService) Delete(id string) error {
	if _, err := s.GetByID(id); err != nil {
		return err
	}
	return s.repo.Delete(id)
}
