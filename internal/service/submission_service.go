package service

import (
	"context"
	"time"

	"asri-backend/internal/groq"
	"asri-backend/internal/model"
	"asri-backend/internal/scoring"
)

type TaskReader interface {
	FindByID(id string) (*model.Task, error)
}

type SubmissionStore interface {
	Create(s *model.AudioSubmission) error
}

type WhisperClient interface {
	Transcribe(ctx context.Context, audio []byte, filename string) (*groq.Transcript, error)
}

type FeedbackClient interface {
	Assess(ctx context.Context, req groq.AssessmentRequest) (*groq.Feedback, error)
}

type SubmissionService struct {
	tasks       TaskReader
	modules     ModuleRepo
	submissions SubmissionStore
	whisper     WhisperClient
	feedback    FeedbackClient
}

func NewSubmissionService(tasks TaskReader, modules ModuleRepo, submissions SubmissionStore, whisper WhisperClient, feedback FeedbackClient) *SubmissionService {
	return &SubmissionService{tasks: tasks, modules: modules, submissions: submissions, whisper: whisper, feedback: feedback}
}

type SubmissionResult struct {
	TranscribedText    string         `json:"transcribed_text"`
	WordAccuracyScore  float64        `json:"word_accuracy_score"`
	FluencyScore       float64        `json:"fluency_score"`
	PronunciationScore float64        `json:"pronunciation_score"`
	Feedback           *groq.Feedback `json:"feedback"`
}

// SubmitAudio runs the PRD §4 pipeline: transcribe → metrics → Groq batch → log.
// Only SPEAKING_RECORDING tasks accepted; audio bytes never persisted (no-storage ruling).
func (s *SubmissionService) SubmitAudio(ctx context.Context, taskID string, audio []byte, filename, culturalTopic string) (*SubmissionResult, error) {
	task, err := s.tasks.FindByID(taskID)
	if err != nil {
		return nil, ErrNotFound
	}
	if task.Type != model.TaskSpeakingRecording {
		return nil, ErrUnsupported
	}

	tr, err := s.whisper.Transcribe(ctx, audio, filename)
	if err != nil {
		return nil, err
	}

	reference := task.Prompt
	if module, err := s.modules.FindByID(task.ModuleID); err == nil && module.TargetTranscript != "" {
		reference = module.TargetTranscript
	}
	wer, _, _, _, _ := scoring.WER(reference, tr.Text)
	accuracy := scoring.Accuracy(wer)
	words := scoring.Words(tr.Text)
	wpm := scoring.WPM(len(words), tr.DurationSec)
	fluency := scoring.FluencyScore(wpm)

	fb, err := s.feedback.Assess(ctx, groq.AssessmentRequest{
		Transcript: tr.Text, Reference: reference, WPM: wpm, Accuracy: accuracy,
		TargetPhonetics: task.TargetPhonetics, CulturalTopic: culturalTopic,
	})
	if err != nil {
		return nil, err
	}

	result := &SubmissionResult{
		TranscribedText: tr.Text, WordAccuracyScore: accuracy, FluencyScore: fluency,
		PronunciationScore: fb.OverallScore, Feedback: fb,
	}
	_ = s.submissions.Create(&model.AudioSubmission{
		TaskID: taskID, AudioFileURL: nil, TranscribedText: tr.Text,
		WordAccuracyScore: accuracy, FluencyScore: fluency,
		PronunciationScore: fb.OverallScore, CreatedAt: time.Now(),
	}) // analytical log best-effort; scoring result stands regardless
	return result, nil
}
