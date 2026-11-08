package service

import (
	"context"
	"encoding/json"
	"os"
	"strconv"
	"strings"
	"time"

	"asri-backend/internal/groq"
	"asri-backend/internal/model"
	"asri-backend/internal/scoring"
)

type TaskReader interface {
	FindByID(id string) (*model.Task, error)
	FindByModule(moduleID string) ([]model.Task, error)
}

type SubmissionStore interface {
	Create(s *model.AudioSubmission) error
	LatestByCode(code string) ([]model.ScoreRow, error)
	LatestByStudent(code, name, chapterID string) ([]model.ScoreRow, error)
}

// strptr keeps pre-name rows NULL-shaped instead of empty-string.
func strptr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
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

// referenceFor picks what a speaking attempt is scored against. Question
// prompts ("...?") use the seeded Model answer when present — scoring a
// 2-sentence answer against a whole monologue transcript yields ~10% for a
// perfect answer. Plain repeat prompts keep module transcript, else prompt.
func referenceFor(prompt, optionsJSON, moduleTranscript string) string {
	if strings.HasSuffix(strings.TrimSpace(prompt), "?") {
		var o struct {
			ModelAnswer string `json:"Model answer"`
		}
		if json.Unmarshal([]byte(optionsJSON), &o) == nil && o.ModelAnswer != "" {
			return o.ModelAnswer
		}
	}
	if moduleTranscript != "" {
		return moduleTranscript
	}
	return prompt
}

// textDelta compensates the browser-transcript path's known downward bias
// (SR noise + wall-clock duration deflating WPM). Env-tuned, default 0.
// shortcut: flat boost, not bias modeling — recalibrate from pilot score
// distributions; per-task adaptive compensation when research demands it.
func textDelta() float64 {
	d, err := strconv.ParseFloat(os.Getenv("SUBMIT_TEXT_DELTA"), 64)
	if err != nil || d < 0 {
		return 0
	}
	return d
}

func boost(v, d float64) float64 {
	if v+d > 100 {
		return 100
	}
	return v + d
}

// SubmitAudio runs the PRD §4 pipeline: transcribe → metrics → Groq batch → log.
// Only SPEAKING_RECORDING tasks accepted; audio bytes never persisted (no-storage ruling).
func (s *SubmissionService) SubmitAudio(ctx context.Context, taskID string, audio []byte, filename, culturalTopic, studentName, accessCode string) (*SubmissionResult, error) {
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

	transcript := ""
	if module, err := s.modules.FindByID(task.ModuleID); err == nil {
		transcript = module.TargetTranscript
	}
	reference := referenceFor(task.Prompt, task.OptionsJSON, transcript)
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
		Mispronounced: scoring.Align(reference, tr.Text),
	}
	_ = s.submissions.Create(&model.AudioSubmission{
		TaskID: taskID, Source: "audio", AudioFileURL: nil, TranscribedText: tr.Text,
		WordAccuracyScore: accuracy, FluencyScore: fluency,
		PronunciationScore: fb.OverallScore, CreatedAt: time.Now(),
		StudentName: strptr(studentName), AccessCode: strptr(accessCode),
	}) // analytical log best-effort; scoring result stands regardless
	return result, nil
}

// SubmitText scores a client-transcribed utterance: pure-Go metrics, zero Groq.
// Built for 40-concurrent classrooms where Whisper-per-submit 502s. Duration
// is client wall-clock (includes thinking pauses — WPM skews low, noted).
func (s *SubmissionService) SubmitText(ctx context.Context, taskID, transcribed string, durationSec float64, studentName, accessCode string) (*SubmissionResult, error) {
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
	reference := ""
	if module, err := s.modules.FindByID(task.ModuleID); err == nil {
		reference = module.TargetTranscript
	}
	reference = referenceFor(task.Prompt, task.OptionsJSON, reference)
	wer, _, _, _, _ := scoring.WER(reference, transcribed)
	accuracy := boost(scoring.Accuracy(wer), textDelta())
	wpm := scoring.WPM(len(scoring.Words(transcribed)), durationSec)
	fluency := boost(scoring.FluencyScore(wpm), textDelta())
	words := scoring.Align(reference, transcribed)

	result := &SubmissionResult{
		TranscribedText: transcribed, WordAccuracyScore: accuracy, FluencyScore: fluency,
		WPM: wpm, Mispronounced: words,
	}
	_ = s.submissions.Create(&model.AudioSubmission{
		TaskID: taskID, Source: "text", AudioFileURL: nil, TranscribedText: transcribed,
		WordAccuracyScore: accuracy, FluencyScore: fluency,
		CreatedAt:   time.Now(),
		StudentName: strptr(studentName), AccessCode: strptr(accessCode),
	}) // analytical log best-effort; scoring result stands regardless
	return result, nil
}

// LatestByCode serves the teacher Nilai view: newest score per (student, task).
func (s *SubmissionService) LatestByCode(code string) ([]model.ScoreRow, error) {
	return s.submissions.LatestByCode(code)
}

// GateTask is one scored (speaking) task in a gate's scope.
type GateTask struct {
	TaskID      string `json:"task_id"`
	ModuleID    string `json:"module_id"`
	ModuleTitle string `json:"module_title"`
}

// GateResult is a student's progress across one gate scope: pass needs every
// scope task attempted with average word accuracy >= 80.
type GateResult struct {
	Scope     []GateTask       `json:"scope"`
	Rows      []model.ScoreRow `json:"rows"`
	Attempted int              `json:"attempted"`
	Average   float64          `json:"average"`
	Pass      bool             `json:"pass"`
}

// GateScores scopes to speaking tasks in same-course modules ordered before
// the gate module (whole chapter when beforeModuleID is empty or yields no
// tasks, e.g. the Self-Reflection overall gate). Quiz tasks score nothing
// server-side, so only SPEAKING_RECORDING tasks can gate.
func (s *SubmissionService) GateScores(code, name, chapterID, beforeModuleID string) (*GateResult, error) {
	rows, err := s.submissions.LatestByStudent(code, name, chapterID)
	if err != nil {
		return nil, err
	}
	var scope []GateTask
	collect := func(m model.Module) error {
		tasks, err := s.tasks.FindByModule(m.ID)
		if err != nil {
			return err
		}
		for _, t := range tasks {
			if t.Type == model.TaskSpeakingRecording {
				scope = append(scope, GateTask{TaskID: t.ID, ModuleID: m.ID, ModuleTitle: m.Title})
			}
		}
		return nil
	}
	if beforeModuleID != "" {
		gate, err := s.modules.FindByID(beforeModuleID)
		if err != nil {
			return nil, ErrNotFound
		}
		mods, err := s.modules.FindByCourse(gate.CourseID)
		if err != nil {
			return nil, err
		}
		for _, m := range mods {
			if m.OrderIndex < gate.OrderIndex {
				if err := collect(m); err != nil {
					return nil, err
				}
			}
		}
	}
	if len(scope) == 0 {
		courses, err := s.courses.FindByChapter(chapterID)
		if err != nil {
			return nil, err
		}
		for _, c := range courses {
			mods, err := s.modules.FindByCourse(c.ID)
			if err != nil {
				return nil, err
			}
			for _, m := range mods {
				if err := collect(m); err != nil {
					return nil, err
				}
			}
		}
	}
	byTask := map[string]model.ScoreRow{}
	for _, r := range rows {
		byTask[r.TaskID] = r
	}
	inScope := map[string]bool{}
	for _, g := range scope {
		inScope[g.TaskID] = true
	}
	var inRows []model.ScoreRow
	var sum float64
	for _, r := range rows {
		if inScope[r.TaskID] {
			inRows = append(inRows, r)
			sum += r.WordAccuracyScore
		}
	}
	avg := 0.0
	if len(inRows) > 0 {
		avg = sum / float64(len(inRows))
	}
	return &GateResult{
		Scope: scope, Rows: inRows, Attempted: len(inRows), Average: avg,
		Pass: len(scope) > 0 && len(inRows) == len(scope) && avg >= 80,
	}, nil
}

// CheckResult reveals the key only inside an attempt response —
// keys never ride GETs. Nothing stored: formative by design.
type CheckResult struct {
	Correct  bool            `json:"correct"`
	Expected any             `json:"expected"`
	Detail   map[string]bool `json:"detail,omitempty"`
}

func emptyAnswer(a any) bool {
	switch v := a.(type) {
	case nil:
		return true
	case string:
		return v == ""
	case []any:
		return len(v) == 0
	case map[string]any:
		return len(v) == 0
	}
	return false
}

// CheckAnswer grades one attempt and reveals the expected answer.
func (s *SubmissionService) CheckAnswer(ctx context.Context, taskID string, answer any) (*CheckResult, error) {
	_ = ctx
	if emptyAnswer(answer) {
		return nil, ErrValidation // no free reveal without an attempt
	}
	task, err := s.tasks.FindByID(taskID)
	if err != nil {
		return nil, ErrNotFound
	}
	ok, detail := scoring.Check(task.Type, task.AnswerKeyJSON, answer)
	var expected any
	_ = json.Unmarshal([]byte(task.AnswerKeyJSON), &expected)
	return &CheckResult{Correct: ok, Expected: expected, Detail: detail}, nil
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
