package groq

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
	"time"
)

// feedbackModel is the assessor (PRD §4). Pinned via env, verified at build.
// ponytail: was llama-3.3-70b-versatile until Groq retired it (instant 404s →
// our 502s); qwen3.8-27b honors response_format json_object, verified live.
var feedbackModel = envOr("GROQ_FEEDBACK_MODEL", "qwen/qwen3.8-27b")

type PhonemeNote struct {
	Word   string `json:"word"`
	Target string `json:"target"`
	Actual string `json:"actual"`
	Tip    string `json:"tip"`
}

type GrammarNote struct {
	Original    string `json:"original"`
	Corrected   string `json:"corrected"`
	Explanation string `json:"explanation"`
}

// Feedback is the PRD §4.2 JSON output contract.
type Feedback struct {
	OverallScore       float64       `json:"overall_score"`
	PhoneticFeedback   []PhonemeNote `json:"phonetic_feedback"`
	GrammarCorrections []GrammarNote `json:"grammar_corrections"`
	CulturalNotes      string        `json:"cultural_context_notes"`
	ActionableTips     []string      `json:"actionable_tips"`
}

type AssessmentRequest struct {
	Transcript      string
	Reference       string
	WPM             float64
	Accuracy        float64
	TargetPhonetics string
	CulturalTopic   string
}

type FeedbackClient struct {
	apiKey string
	http   *http.Client
}

func NewFeedbackClient(apiKey string) *FeedbackClient {
	return &FeedbackClient{apiKey: apiKey, http: &http.Client{Timeout: 90 * time.Second}}
}

func (c *FeedbackClient) Assess(ctx context.Context, req AssessmentRequest) (*Feedback, error) {
	prompt := fmt.Sprintf(`You assess English speaking practice for Indonesian students learning Riau Malay culture in English.
Reference sentence: %q
Student said: %q
Measured: WPM %.1f (ideal 110-150), word accuracy %.1f%%.
Target phonetics lexicon: %s
Cultural topic: %s
Reply with JSON only, exactly these keys: overall_score (0-100 float), phonetic_feedback (array of {word,target,actual,tip}), grammar_corrections (array of {original,corrected,explanation}), cultural_context_notes (string), actionable_tips (array of max 3 short strings).`,
		req.Reference, req.Transcript, req.WPM, req.Accuracy, req.TargetPhonetics, req.CulturalTopic)

	payload, _ := json.Marshal(map[string]any{
		"model": feedbackModel,
		"messages": []map[string]string{
			{"role": "system", "content": "You reply with JSON only. No markdown, no commentary."},
			{"role": "user", "content": prompt},
		},
		"response_format": map[string]string{"type": "json_object"},
		"temperature":     0.2,
	})

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, baseURL+"/chat/completions", bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Authorization", "Bearer "+c.apiKey)
	httpReq.Header.Set("Content-Type", "application/json")

	res, err := c.http.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	raw, _ := io.ReadAll(io.LimitReader(res.Body, 1<<20))
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("groq feedback: HTTP %d: %s", res.StatusCode, truncate(raw, 200))
	}
	var envelope struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}
	if err := json.Unmarshal(raw, &envelope); err != nil {
		return nil, err
	}
	if len(envelope.Choices) == 0 {
		return nil, fmt.Errorf("groq feedback: empty choices")
	}
	var fb Feedback
	if err := json.Unmarshal([]byte(envelope.Choices[0].Message.Content), &fb); err != nil {
		return nil, fmt.Errorf("groq feedback: bad JSON contract: %w", err)
	}
	return &fb, nil
}

// ChapterVerdict is the per-chapter rollup contract. Averages and recurring
// words are computed in Go — the model only writes band + next steps.
type ChapterVerdict struct {
	OverallBand    string   `json:"overall_band"`
	AvgAccuracy    float64  `json:"avg_accuracy"`
	AvgFluency     float64  `json:"avg_fluency"`
	RecurringWords []string `json:"recurring_words"`
	NextSteps      []string `json:"next_steps"`
}

type ChapterAttemptLine struct {
	Reference   string
	Transcribed string
	Accuracy    float64
	WPM         float64
}

type ChapterRollupRequest struct {
	ChapterTitle string
	Attempts     []ChapterAttemptLine
	AvgAccuracy  float64
	AvgFluency   float64
	WordHits     map[string]int
}

func (c *FeedbackClient) ChapterRollup(ctx context.Context, req ChapterRollupRequest) (*ChapterVerdict, error) {
	var sb strings.Builder
	for i, a := range req.Attempts {
		fmt.Fprintf(&sb, "%d. target %q, said %q (acc %.0f%%, %.0f wpm)\n",
			i+1, truncate([]byte(a.Reference), 200), truncate([]byte(a.Transcribed), 200), a.Accuracy, a.WPM)
	}
	prompt := fmt.Sprintf(`You recap English speaking practice across one chapter (%s) for an Indonesian student.
Measured averages: accuracy %.1f%%, fluency %.1f/100. Per-attempt log:
%s
Reply with JSON only, exactly: overall_band (one of "Needs practice", "Getting there", "Strong"), next_steps (array of max 3 short concrete actions).`,
		req.ChapterTitle, req.AvgAccuracy, req.AvgFluency, sb.String())

	payload, _ := json.Marshal(map[string]any{
		"model": feedbackModel,
		"messages": []map[string]string{
			{"role": "system", "content": "You reply with JSON only. No markdown, no commentary."},
			{"role": "user", "content": prompt},
		},
		"response_format": map[string]string{"type": "json_object"},
		"temperature":     0.2,
		"max_tokens":      500, // verdict is small by contract; bounds latency
	})

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, baseURL+"/chat/completions", bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Authorization", "Bearer "+c.apiKey)
	httpReq.Header.Set("Content-Type", "application/json")

	res, err := c.http.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	raw, _ := io.ReadAll(io.LimitReader(res.Body, 1<<20))
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("groq chapter rollup: HTTP %d: %s", res.StatusCode, truncate(raw, 200))
	}
	var envelope struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}
	if err := json.Unmarshal(raw, &envelope); err != nil {
		return nil, err
	}
	if len(envelope.Choices) == 0 {
		return nil, fmt.Errorf("groq chapter rollup: empty choices")
	}
	var partial struct {
		OverallBand string   `json:"overall_band"`
		NextSteps   []string `json:"next_steps"`
	}
	if err := json.Unmarshal([]byte(envelope.Choices[0].Message.Content), &partial); err != nil {
		return nil, fmt.Errorf("groq chapter rollup: bad JSON contract: %w", err)
	}
	words := topHits(req.WordHits, 8)
	return &ChapterVerdict{
		OverallBand: partial.OverallBand, AvgAccuracy: req.AvgAccuracy,
		AvgFluency: req.AvgFluency, RecurringWords: words, NextSteps: partial.NextSteps,
	}, nil
}

// topHits returns the most frequent words, most-frequent first.
func topHits(hits map[string]int, n int) []string {
	type kv struct {
		w string
		c int
	}
	var all []kv
	for w, c := range hits {
		all = append(all, kv{w, c})
	}
	sort.Slice(all, func(i, j int) bool { return all[i].c > all[j].c })
	var out []string
	for i := 0; i < len(all) && i < n; i++ {
		out = append(out, all[i].w)
	}
	return out
}
