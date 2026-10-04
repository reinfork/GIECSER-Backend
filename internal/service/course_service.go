package service

import (
	"errors"

	"asri-backend/internal/model"

	"gorm.io/gorm"
)

type CourseRepo interface {
	FindByChapter(chapterID string) ([]model.Course, error)
	FindByID(id string) (*model.Course, error)
	Create(c *model.Course) error
	Update(c *model.Course) error
	Delete(id string) error
}

// ChapterStore is what content services need for parent lookups — fakes welcome.
type ChapterStore interface {
	FindByID(id string) (*model.Chapter, error)
}

type CourseService struct {
	repo     CourseRepo
	chapters ChapterStore
}

func NewCourseService(repo CourseRepo, chapters ChapterStore) *CourseService {
	return &CourseService{repo: repo, chapters: chapters}
}

func (s *CourseService) ListByChapter(chapterID string) ([]model.Course, error) {
	return s.repo.FindByChapter(chapterID)
}

func (s *CourseService) GetByID(id string) (*model.Course, error) {
	course, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return course, nil
}

func (s *CourseService) Create(input model.CreateCourseInput) (*model.Course, error) {
	if _, err := s.chapters.FindByID(input.ChapterID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	course := &model.Course{ChapterID: input.ChapterID, Title: input.Title, Description: input.Description, OrderIndex: input.OrderIndex}
	if err := s.repo.Create(course); err != nil {
		return nil, err
	}
	return course, nil
}

func (s *CourseService) Update(id string, input model.UpdateCourseInput) (*model.Course, error) {
	course, err := s.GetByID(id)
	if err != nil {
		return nil, err
	}
	if input.Title != nil {
		course.Title = *input.Title
	}
	if input.Description != nil {
		course.Description = *input.Description
	}
	if input.OrderIndex != nil {
		course.OrderIndex = *input.OrderIndex
	}
	if err := s.repo.Update(course); err != nil {
		return nil, err
	}
	return course, nil
}

func (s *CourseService) Delete(id string) error {
	if _, err := s.GetByID(id); err != nil {
		return err
	}
	return s.repo.Delete(id)
}
