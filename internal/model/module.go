package model

import "time"

// Chapter/module content types (PRD §2). Lives on modules — the practicable level.
const (
	ModuleVideo     = "VIDEO"
	ModuleMonologue = "MONOLOGUE"
	ModuleDialogue  = "DIALOGUE"
	ModuleQuiz      = "QUIZ"
	ModuleOralTest  = "ORAL_TEST"
)

// Module is the practicable content unit: text to read, audio model, target transcript.
type Module struct {
	ID               string    `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	CourseID         string    `json:"course_id" gorm:"type:uuid;not null;index"`
	Title            string    `json:"title" gorm:"not null"`
	Type             string    `json:"type" gorm:"not null"`
	ContentText      string    `json:"content_text" gorm:"type:text"`
	MediaURL         string    `json:"media_url"`
	TargetTranscript string    `json:"target_transcript" gorm:"type:text"`
	AudioModelURL    string    `json:"audio_model_url"`
	OrderIndex       int       `json:"order_index" gorm:"not null"`
	Tasks            []Task    `json:"tasks,omitempty" gorm:"foreignKey:ModuleID"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type CreateModuleInput struct {
	CourseID         string `json:"course_id" binding:"required"`
	Title            string `json:"title" binding:"required"`
	Type             string `json:"type" binding:"required,oneof=VIDEO MONOLOGUE DIALOGUE QUIZ ORAL_TEST"`
	ContentText      string `json:"content_text"`
	MediaURL         string `json:"media_url"`
	TargetTranscript string `json:"target_transcript"`
	AudioModelURL    string `json:"audio_model_url"`
	OrderIndex       int    `json:"order_index" binding:"gte=1"`
}

type UpdateModuleInput struct {
	Title            *string `json:"title" binding:"omitempty,min=1"`
	Type             *string `json:"type" binding:"omitempty,oneof=VIDEO MONOLOGUE DIALOGUE QUIZ ORAL_TEST"`
	ContentText      *string `json:"content_text"`
	MediaURL         *string `json:"media_url"`
	TargetTranscript *string `json:"target_transcript"`
	AudioModelURL    *string `json:"audio_model_url"`
	OrderIndex       *int    `json:"order_index" binding:"omitempty,gte=1"`
}
