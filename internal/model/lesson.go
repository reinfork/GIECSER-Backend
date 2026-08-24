package model

import "time"

type Lesson struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	CourseID   uint      `json:"course_id" gorm:"index;not null"`
	Title      string    `json:"title" gorm:"not null"`
	OrderIndex int       `json:"order_index" gorm:"not null"`
	TargetText string    `json:"target_text" gorm:"type:text"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CreateLessonInput struct {
	CourseID   uint   `json:"course_id" binding:"required"`
	Title      string `json:"title" binding:"required"`
	OrderIndex int    `json:"order_index" binding:"required,gte=1"`
	TargetText string `json:"target_text" binding:"required"`
}
