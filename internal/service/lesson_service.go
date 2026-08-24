package service

import (
	"errors"

	"asri-backend/internal/model"
	"asri-backend/internal/repository"

	"gorm.io/gorm"
)

type LessonService struct {
	lessonRepo *repository.LessonRepository
	courseRepo *repository.CourseRepository
}

func NewLessonService(lessonRepo *repository.LessonRepository, courseRepo *repository.CourseRepository) *LessonService {
	return &LessonService{lessonRepo: lessonRepo, courseRepo: courseRepo}
}

func (s *LessonService) GetByID(id uint) (*model.Lesson, error) {
	lesson, err := s.lessonRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return lesson, nil
}

func (s *LessonService) ListByCourse(courseID uint) ([]model.Lesson, error) {
	// Verify course exists
	if _, err := s.courseRepo.FindByID(courseID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrCourseNotFound
		}
		return nil, err
	}
	return s.lessonRepo.FindByCourseID(courseID)
}

func (s *LessonService) ListByCoursePaginated(courseID uint, page, limit int) ([]model.Lesson, int64, error) {
	if _, err := s.courseRepo.FindByID(courseID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 0, ErrCourseNotFound
		}
		return nil, 0, err
	}
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}
	offset := (page - 1) * limit
	return s.lessonRepo.FindByCourseIDPaginated(courseID, limit, offset)
}

func (s *LessonService) Create(input model.CreateLessonInput) (*model.Lesson, error) {
	if _, err := s.courseRepo.FindByID(input.CourseID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrCourseNotFound
		}
		return nil, err
	}
	exists, err := s.lessonRepo.ExistsByCourseIDAndOrder(input.CourseID, input.OrderIndex, nil)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrConflict
	}
	lesson := &model.Lesson{
		CourseID:   input.CourseID,
		Title:      input.Title,
		OrderIndex: input.OrderIndex,
		TargetText: input.TargetText,
	}
	if err := s.lessonRepo.Create(lesson); err != nil {
		return nil, err
	}
	return lesson, nil
}

func (s *LessonService) Update(id uint, input model.UpdateLessonInput) (*model.Lesson, error) {
	lesson, err := s.GetByID(id)
	if err != nil {
		return nil, err
	}
	if input.Title != nil {
		lesson.Title = *input.Title
	}
	if input.TargetText != nil {
		lesson.TargetText = *input.TargetText
	}
	if input.OrderIndex != nil {
		exists, err := s.lessonRepo.ExistsByCourseIDAndOrder(lesson.CourseID, *input.OrderIndex, &id)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, ErrConflict
		}
		lesson.OrderIndex = *input.OrderIndex
	}
	if err := s.lessonRepo.Update(lesson); err != nil {
		return nil, err
	}
	return lesson, nil
}

func (s *LessonService) Delete(id uint) error {
	if _, err := s.GetByID(id); err != nil {
		return err
	}
	return s.lessonRepo.Delete(id)
}

func (s *LessonService) Practice(id uint, transcript string) (*model.PracticeResult, error) {
	lesson, err := s.GetByID(id)
	if err != nil {
		return nil, err
	}
	// Normalize and compute simple word-match score
	target := normalize(lesson.TargetText)
	spoken := normalize(transcript)
	targetWords := splitWords(target)
	spokenWords := splitWords(spoken)
	spokenSet := make(map[string]bool, len(spokenWords))
	for _, w := range spokenWords {
		spokenSet[w] = true
	}
	matched := 0
	for _, w := range targetWords {
		if spokenSet[w] {
			matched++
		}
	}
	total := len(targetWords)
	score := 0
	if total > 0 {
		score = matched * 100 / total
	}
	feedback := feedbackForScore(score)
	return &model.PracticeResult{
		LessonID:   lesson.ID,
		TargetText: lesson.TargetText,
		Transcript: transcript,
		Score:      score,
		Feedback:   feedback,
		Matched:    matched,
		TotalWords: total,
	}, nil
}

func normalize(s string) string {
	// lower case and trim; keep simple for now
	out := ""
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			r += 'a' - 'A'
		}
		out += string(r)
	}
	return out
}

func splitWords(s string) []string {
	var words []string
	cur := ""
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			cur += string(r)
		} else if cur != "" {
			words = append(words, cur)
			cur = ""
		}
	}
	if cur != "" {
		words = append(words, cur)
	}
	return words
}

func feedbackForScore(score int) string {
	switch {
	case score >= 90:
		return "Excellent! Pronunciation is very accurate."
	case score >= 70:
		return "Good! Minor corrections needed."
	case score >= 50:
		return "Fair — keep practicing pronunciation."
	default:
		return "Needs practice — try listening and repeating again."
	}
}
