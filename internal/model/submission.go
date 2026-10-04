package model

import "time"

// AudioSubmission is an analytical evaluation log (PRD §2), not progress tracking.
// AudioFileURL stays NULL: clips are scored then deleted (privacy + storage ruling).
// Research analysis reads this table directly — no list endpoint by design.
type AudioSubmission struct {
	ID                 string    `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	TaskID             string    `json:"task_id" gorm:"type:uuid;not null;index"`
	AudioFileURL       *string   `json:"audio_file_url"`
	TranscribedText    string    `json:"transcribed_text" gorm:"type:text"`
	WordAccuracyScore  float64   `json:"word_accuracy_score"`
	FluencyScore       float64   `json:"fluency_score"`
	PronunciationScore float64   `json:"pronunciation_score"`
	CreatedAt          time.Time `json:"created_at"`
}
