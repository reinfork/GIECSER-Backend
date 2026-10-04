package model

import "time"

// Chapter is the top-level container (one per cultural topic).
// Pure container — practicable content lives in modules below.
type Chapter struct {
	ID          string    `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
	OrderIndex  int       `json:"order_index" gorm:"not null"`
	Courses     []Course  `json:"courses,omitempty" gorm:"foreignKey:ChapterID"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateChapterInput struct {
	Title       string `json:"title" binding:"required,min=3"`
	Description string `json:"description" binding:"required"`
	OrderIndex  int    `json:"order_index" binding:"gte=1"`
}

type UpdateChapterInput struct {
	Title       *string `json:"title" binding:"omitempty,min=3"`
	Description *string `json:"description" binding:"omitempty,min=1"`
	OrderIndex  *int    `json:"order_index" binding:"omitempty,gte=1"`
}
