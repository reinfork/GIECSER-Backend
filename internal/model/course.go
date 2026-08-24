package model

import "time"

type Course struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
	Lessons     []Lesson  `json:"lessons" gorm:"foreignKey:CourseID"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateCourseInput struct {
	Title       string `json:"title" binding:"required,min=3"`
	Description string `json:"description" binding:"required"`
}
