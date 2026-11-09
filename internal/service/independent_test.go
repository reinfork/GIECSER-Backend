package service

import (
	"context"
	"testing"

	"asri-backend/internal/groq"
	"asri-backend/internal/model"
)

type fakeTasks struct{ task *model.Task }

func (f fakeTasks) FindByID(id string) (*model.Task, error) { return f.task, nil }
func (f fakeTasks) FindByModule(moduleID string) ([]model.Task, error) {
	return []model.Task{*f.task}, nil
}

type fakeModules struct{ module *model.Module }

func (f fakeModules) FindByCourse(courseID string) ([]model.Module, error) {
	return []model.Module{*f.module}, nil
}
func (f fakeModules) FindByID(id string) (*model.Module, error) { return f.module, nil }
func (f fakeModules) Create(m *model.Module) error              { return nil }
func (f fakeModules) Update(m *model.Module) error              { return nil }
func (f fakeModules) Delete(id string) error                    { return nil }

type fakeCourses struct{}

func (f fakeCourses) FindByChapter(chapterID string) ([]model.Course, error) { return nil, nil }
func (f fakeCourses) FindByID(id string) (*model.Course, error) {
	return &model.Course{ID: id}, nil
}
func (f fakeCourses) Create(c *model.Course) error { return nil }
func (f fakeCourses) Update(c *model.Course) error { return nil }
func (f fakeCourses) Delete(id string) error       { return nil }

type fakeChapters struct{}

func (f fakeChapters) FindByID(id string) (*model.Chapter, error) {
	return &model.Chapter{ID: id}, nil
}

type fakeSubs struct{ last *model.AudioSubmission }

func (f *fakeSubs) Create(s *model.AudioSubmission) error { f.last = s; return nil }
func (f *fakeSubs) LatestByCode(code string) ([]model.ScoreRow, error) {
	return nil, nil
}
func (f *fakeSubs) LatestByStudent(code, name, chapterID string) ([]model.ScoreRow, error) {
	return nil, nil
}

type fakeWhisper struct{ text string }

func (f fakeWhisper) Transcribe(ctx context.Context, audio []byte, filename string) (*groq.Transcript, error) {
	return &groq.Transcript{Text: f.text, DurationSec: 60}, nil
}

type fakeFeedback struct{ score float64 }

func (f fakeFeedback) Assess(ctx context.Context, req groq.AssessmentRequest) (*groq.Feedback, error) {
	return &groq.Feedback{OverallScore: f.score}, nil
}
func (f fakeFeedback) ChapterRollup(ctx context.Context, req groq.ChapterRollupRequest) (*groq.ChapterVerdict, error) {
	return &groq.ChapterVerdict{}, nil
}

func independentSvc(task *model.Task, module *model.Module, subs *fakeSubs) *SubmissionService {
	return NewSubmissionService(
		fakeTasks{task: task}, fakeModules{module: module}, fakeCourses{}, fakeChapters{}, subs,
		fakeWhisper{text: "completely unrelated free speech about my holiday"},
		fakeFeedback{score: 82},
	)
}

func TestSubmitAudioIndependentUsesAIScore(t *testing.T) {
	task := &model.Task{ID: "t1", ModuleID: "m1", Type: model.TaskSpeakingRecording, Prompt: "Record your monologue."}
	module := &model.Module{ID: "m1", Type: model.ModuleOralTest, TargetTranscript: ""}
	subs := &fakeSubs{}
	res, err := independentSvc(task, module, subs).SubmitAudio(context.Background(), "t1", []byte("x"), "a.webm", "topic", "S", "C")
	if err != nil {
		t.Fatal(err)
	}
	if res.WordAccuracyScore != 82 {
		t.Fatalf("independent accuracy must be the AI verdict, got %v", res.WordAccuracyScore)
	}
	if len(res.Mispronounced) != 0 {
		t.Fatalf("no alignment without a reference, got %v", res.Mispronounced)
	}
	if subs.last == nil || subs.last.WordAccuracyScore != 82 {
		t.Fatalf("stored row must carry the AI score")
	}
}

func TestSubmitTextIndependentUsesFluency(t *testing.T) {
	task := &model.Task{ID: "t1", ModuleID: "m1", Type: model.TaskSpeakingRecording, Prompt: "Record your monologue."}
	module := &model.Module{ID: "m1", Type: model.ModuleOralTest, TargetTranscript: ""}
	subs := &fakeSubs{}
	res, err := independentSvc(task, module, subs).SubmitText(context.Background(), "t1", "free speech with no reference at all", 60, "S", "C")
	if err != nil {
		t.Fatal(err)
	}
	if res.WordAccuracyScore != res.FluencyScore {
		t.Fatalf("text-path independent accuracy must equal fluency, got %v vs %v", res.WordAccuracyScore, res.FluencyScore)
	}
}
