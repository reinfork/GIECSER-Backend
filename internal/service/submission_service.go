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
	ChapterRollup(ctx context.Context, req groq.ChapterRollupRequest) (*groq.ChapterVerdict, error)
}

type SubmissionService struct {
	tasks       TaskReader
	modules     ModuleRepo
	courses     CourseRepo
	chapters    ChapterStore
	submissions SubmissionStore
	whisper     WhisperClient
	feedback    FeedbackClient
}

func NewSubmissionService(tasks TaskReader, modules ModuleRepo, courses CourseRepo, chapters ChapterStore, submissions SubmissionStore, whisper WhisperClient, feedback FeedbackClient) *SubmissionService {
	return &SubmissionService{tasks: tasks, modules: modules, courses: courses, chapters: chapters, submissions: submissions, whisper: whisper, feedback: feedback}
}

type SubmissionResult struct {
	TranscribedText    string         `json:"transcribed_text"`
	WordAccuracyScore  float64        `json:"word_accuracy_score"`
	FluencyScore       float64        `json:"fluency_score"`
	PronunciationScore float64        `json:"pronunciation_score"`
	WPM                float64        `json:"wpm"`
	Mispronounced      []string       `json:"mispronounced,omitempty"`
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
		PronunciationScore: fb.OverallScore, WPM: wpm, Feedback: fb,
	}
	_ = s.submissions.Create(&model.AudioSubmission{
		TaskID: taskID, AudioFileURL: nil, TranscribedText: tr.Text,
		WordAccuracyScore: accuracy, FluencyScore: fluency,
		PronunciationScore: fb.OverallScore, CreatedAt: time.Now(),
	}) // analytical log best-effort; scoring result stands regardless
	return result, nil
}

// SubmitText scores a client-transcribed utterance: pure-Go metrics, zero Groq.
// Built for 40-concurrent classrooms where Whisper-per-submit 502s. Duration
// is client wall-clock (includes thinking pauses — WPM skews low, noted).
func (s *SubmissionService) SubmitText(ctx context.Context, taskID, transcribed string, durationSec float64) (*SubmissionResult, error) {
	_ = ctx
	if len(transcribed) == 0 || len(transcribed) > 2000 {
		return nil, ErrValidation
	}
	task, err := s.tasks.FindByID(taskID)
	if err != nil {
		return nil, ErrNotFound
	}
	if task.Type != model.TaskSpeakingRecording {
		return nil, ErrUnsupported
	}
	reference := task.Prompt
	if module, err := s.modules.FindByID(task.ModuleID); err == nil && module.TargetTranscript != "" {
		reference = module.TargetTranscript
	}
	wer, _, _, _, _ := scoring.WER(reference, transcribed)
	accuracy := scoring.Accuracy(wer)
	wpm := scoring.WPM(len(scoring.Words(transcribed)), durationSec)
	fluency := scoring.FluencyScore(wpm)
	words := scoring.Align(reference, transcribed)

	result := &SubmissionResult{
		TranscribedText: transcribed, WordAccuracyScore: accuracy, FluencyScore: fluency,
		WPM: wpm, Mispronounced: words,
	}
	_ = s.submissions.Create(&model.AudioSubmission{
		TaskID: taskID, AudioFileURL: nil, TranscribedText: transcribed,
		WordAccuracyScore: accuracy, FluencyScore: fluency,
		CreatedAt: time.Now(),
	}) // analytical log best-effort; scoring result stands regardless
	return result, nil
}

// ChapterAttempt is one client-held practice result (latest per task; see
// frontend aggregate). Reference/transcribed ride along so the rollup prompt
// needs no log reads — the log has no student key by design.
type ChapterAttempt struct {
	TaskID        string   `json:"task_id"`
	Reference     string   `json:"reference"`
	Transcribed   string   `json:"transcribed"`
	WordAccuracy  float64  `json:"word_accuracy"`
	Fluency       float64  `json:"fluency"`
	Pronunciation float64  `json:"pronunciation"`
	WPM           float64  `json:"wpm"`
	PhoneticWords []string `json:"phonetic_words"`
}

// ChapterFeedback rolls client-collected attempts into one Groq verdict.
// Stateless: stores nothing, validates task→chapter linkage, one LLM call.
func (s *SubmissionService) ChapterFeedback(ctx context.Context, chapterID string, attempts []ChapterAttempt) (*groq.ChapterVerdict, error) {
	if len(attempts) == 0 || len(attempts) > 20 {
		return nil, ErrValidation
	}
	chapter, err := s.chapters.FindByID(chapterID)
	if err != nil {
		return nil, ErrNotFound
	}
	var sumAcc, sumFlu float64
	wordHits := map[string]int{}
	var lines []groq.ChapterAttemptLine
	for _, a := range attempts {
		task, err := s.tasks.FindByID(a.TaskID)
		if err != nil {
			return nil, ErrNotFound
		}
		module, err := s.modules.FindByID(task.ModuleID)
		if err != nil {
			return nil, ErrNotFound
		}
		course, err := s.courses.FindByID(module.CourseID)
		if err != nil {
			return nil, ErrNotFound
		}
		if course.ChapterID != chapterID {
			return nil, ErrValidation // cross-chapter stuffing
		}
		sumAcc += a.WordAccuracy
		sumFlu += a.Fluency
		for _, w := range a.PhoneticWords {
			wordHits[w]++
		}
		lines = append(lines, groq.ChapterAttemptLine{
			Reference: a.Reference, Transcribed: a.Transcribed,
			Accuracy: a.WordAccuracy, WPM: a.WPM,
		})
	}
	return s.feedback.ChapterRollup(ctx, groq.ChapterRollupRequest{
		ChapterTitle: chapter.Title, Attempts: lines,
		AvgAccuracy: sumAcc / float64(len(attempts)),
		AvgFluency:  sumFlu / float64(len(attempts)),
		WordHits:    wordHits,
	})
}
