package model

import "time"

// ScoreRow is one latest-per-(student, task) score with its chapter path
// for the teacher Nilai view. Read-only projection, never migrated.
type ScoreRow struct {
	StudentName        string    `json:"student_name"`
	TaskID             string    `json:"task_id"`
	TaskType           string    `json:"task_type"`
	ChapterTitle       string    `json:"chapter_title"`
	CourseTitle        string    `json:"course_title"`
	ModuleTitle        string    `json:"module_title"`
	TranscribedText    string    `json:"transcribed_text"`
	WordAccuracyScore  float64   `json:"word_accuracy_score"`
	FluencyScore       float64   `json:"fluency_score"`
	PronunciationScore float64   `json:"pronunciation_score"`
	WPM                float64   `json:"wpm"`
	CreatedAt          time.Time `json:"created_at"`
}

// AudioSubmission is an analytical evaluation log (PRD §2), now also the
// per-student score store: student_name + access_code attribute each row.
// AudioFileURL stays NULL: clips are scored then deleted (privacy + storage ruling).
// Teacher Nilai reads via LatestByCode — the research-direct-read ruling it replaces.
// Source marks the scoring path ("audio" = Whisper+Groq, "text" = browser
// transcript + pure-Go metrics + SUBMIT_TEXT_DELTA) so research can split them.
type AudioSubmission struct {
	ID                 string  `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	TaskID             string  `json:"task_id" gorm:"type:uuid;not null;index"`
	Source             string  `json:"source"`
	AudioFileURL       *string `json:"audio_file_url"`
	TranscribedText    string  `json:"transcribed_text" gorm:"type:text"`
	WordAccuracyScore  float64 `json:"word_accuracy_score"`
	FluencyScore       float64 `json:"fluency_score"`
	PronunciationScore float64 `json:"pronunciation_score"`
	WPM                float64 `json:"wpm"`
	// ponytail: nullable student key + code string (no FK — codes expire
	// and rows must survive them); pre-name rows read back NULL.
	StudentName *string   `json:"student_name"`
	AccessCode  *string   `json:"access_code" gorm:"index"`
	CreatedAt   time.Time `json:"created_at"`
}
