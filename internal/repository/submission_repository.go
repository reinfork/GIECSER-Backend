package repository

import (
	"asri-backend/internal/model"

	"gorm.io/gorm"
)

type SubmissionRepository struct {
	db *gorm.DB
}

func NewSubmissionRepository(db *gorm.DB) *SubmissionRepository {
	return &SubmissionRepository{db: db}
}

func (r *SubmissionRepository) Create(s *model.AudioSubmission) error {
	return r.db.Create(s).Error
}

// LatestByStudent returns the newest score per task for one student in one
// class session, optionally narrowed to a chapter. Powers student section
// gates; scoped strictly by the session identity sealed in the JWT.
// ponytail: DISTINCT ON (postgres-only) — window fn when portability matters.
func (r *SubmissionRepository) LatestByStudent(code, name, chapterID string) ([]model.ScoreRow, error) {
	var rows []model.ScoreRow
	err := r.db.Raw(`
		SELECT DISTINCT ON (s.task_id)
			s.student_name, s.task_id, t.type AS task_type,
			ch.title AS chapter_title, c.title AS course_title, m.title AS module_title,
			s.transcribed_text, s.word_accuracy_score, s.fluency_score,
			s.pronunciation_score, s.created_at
		FROM audio_submissions s
		JOIN tasks t ON t.id = s.task_id
		JOIN modules m ON m.id = t.module_id
		JOIN courses c ON c.id = m.course_id
		JOIN chapters ch ON ch.id = c.chapter_id
		WHERE s.access_code = ? AND s.student_name = ? AND ch.id = ?
		ORDER BY s.task_id, s.created_at DESC`, code, name, chapterID).Scan(&rows).Error
	return rows, err
}

// LatestByCode returns the newest score per (student, task) for one class
// code, with chapter path attached for teacher navigation.
// ponytail: DISTINCT ON (postgres-only) — window fn when portability matters.
func (r *SubmissionRepository) LatestByCode(code string) ([]model.ScoreRow, error) {
	var rows []model.ScoreRow
	err := r.db.Raw(`
		SELECT DISTINCT ON (s.student_name, s.task_id)
			s.student_name, s.task_id, t.type AS task_type,
			ch.title AS chapter_title, c.title AS course_title, m.title AS module_title,
			s.transcribed_text, s.word_accuracy_score, s.fluency_score,
			s.pronunciation_score, s.created_at
		FROM audio_submissions s
		JOIN tasks t ON t.id = s.task_id
		JOIN modules m ON m.id = t.module_id
		JOIN courses c ON c.id = m.course_id
		JOIN chapters ch ON ch.id = c.chapter_id
		WHERE s.access_code = ? AND s.student_name IS NOT NULL
		ORDER BY s.student_name, s.task_id, s.created_at DESC`, code).Scan(&rows).Error
	return rows, err
}
