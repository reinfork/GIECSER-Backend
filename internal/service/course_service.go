package service

import (
	"errors"

	"asri-backend/internal/model"
	"asri-backend/internal/repository"

	"gorm.io/gorm"
)

type CourseService struct {
	repo *repository.CourseRepository
}

func NewCourseService(repo *repository.CourseRepository) *CourseService {
	return &CourseService{repo: repo}
}

func (s *CourseService) List() ([]model.Course, error) {
	return s.repo.FindAll()
}

func (s *CourseService) ListPaginated(page, limit int) ([]model.Course, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}
	offset := (page - 1) * limit
	return s.repo.FindAllPaginated(limit, offset)
}

func (s *CourseService) GetByID(id uint) (*model.Course, error) {
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
	course := &model.Course{
		Title:       input.Title,
		Description: input.Description,
	}
	if err := s.repo.Create(course); err != nil {
		return nil, err
	}
	return course, nil
}

func (s *CourseService) Update(id uint, input model.UpdateCourseInput) (*model.Course, error) {
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
	if err := s.repo.Update(course); err != nil {
		return nil, err
	}
	return course, nil
}

func (s *CourseService) Delete(id uint) error {
	if _, err := s.GetByID(id); err != nil {
		return err
	}
	return s.repo.Delete(id)
}
