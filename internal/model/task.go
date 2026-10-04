package model

import "time"

const (
	TaskMultipleChoice    = "MULTIPLE_CHOICE"
	TaskMatching          = "MATCHING"
	TaskOpenEnded         = "OPEN_ENDED"
	TaskSpeakingRecording = "SPEAKING_RECORDING"
)

// Task is one exercisable unit. Options/keys ride as TEXT holding JSON
// (stdlib encoding/json; no extra driver dep for jsonb).
// AnswerKeyJSON is NEVER serialized to session tokens — see TaskPublic.
type Task struct {
	ID              string    `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	ModuleID        string    `json:"module_id" gorm:"type:uuid;not null;index"`
	Type            string    `json:"type" gorm:"not null"`
	Prompt          string    `json:"prompt" gorm:"type:text"`
	OptionsJSON     string    `json:"options_json" gorm:"type:text"`
	AnswerKeyJSON   string    `json:"answer_key_json" gorm:"type:text"`
	TargetPhonetics string    `json:"target_phonetics" gorm:"type:text"`
	OrderIndex      int       `json:"order_index" gorm:"not null"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// TaskPublic is the student-facing shape: answer key stripped.
type TaskPublic struct {
	ID              string `json:"id"`
	ModuleID        string `json:"module_id"`
	Type            string `json:"type"`
	Prompt          string `json:"prompt"`
	OptionsJSON     string `json:"options_json"`
	TargetPhonetics string `json:"target_phonetics"`
	OrderIndex      int    `json:"order_index"`
}

func (t Task) Public() TaskPublic {
	return TaskPublic{
		ID: t.ID, ModuleID: t.ModuleID, Type: t.Type, Prompt: t.Prompt,
		OptionsJSON: t.OptionsJSON, TargetPhonetics: t.TargetPhonetics, OrderIndex: t.OrderIndex,
	}
}

type CreateTaskInput struct {
	ModuleID        string `json:"module_id" binding:"required"`
	Type            string `json:"type" binding:"required,oneof=MULTIPLE_CHOICE MATCHING OPEN_ENDED SPEAKING_RECORDING"`
	Prompt          string `json:"prompt" binding:"required"`
	OptionsJSON     string `json:"options_json"`
	AnswerKeyJSON   string `json:"answer_key_json"`
	TargetPhonetics string `json:"target_phonetics"`
	OrderIndex      int    `json:"order_index" binding:"gte=1"`
}

type UpdateTaskInput struct {
	Type            *string `json:"type" binding:"omitempty,oneof=MULTIPLE_CHOICE MATCHING OPEN_ENDED SPEAKING_RECORDING"`
	Prompt          *string `json:"prompt" binding:"omitempty,min=1"`
	OptionsJSON     *string `json:"options_json"`
	AnswerKeyJSON   *string `json:"answer_key_json"`
	TargetPhonetics *string `json:"target_phonetics"`
	OrderIndex      *int    `json:"order_index" binding:"omitempty,gte=1"`
}
