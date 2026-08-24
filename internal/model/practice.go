package model

type PracticeInput struct {
	Transcript string `json:"transcript" binding:"required,min=1"`
}

type PracticeResult struct {
	LessonID   uint   `json:"lesson_id"`
	TargetText string `json:"target_text"`
	Transcript string `json:"transcript"`
	Score      int    `json:"score"` // 0-100
	Feedback   string `json:"feedback"`
	Matched    int    `json:"matched_words"`
	TotalWords int    `json:"total_words"`
}
