// Package groq holds stdlib HTTP clients for Groq APIs (no SDK dep).
package groq

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"asri-backend/internal/scoring"
)

const baseURL = "https://api.groq.com/openai/v1"

// whisperModel is overridable for tests; Groq hosts whisper-large-v3.
var whisperModel = envOr("GROQ_WHISPER_MODEL", "whisper-large-v3")

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

type WhisperClient struct {
	apiKey string
	http   *http.Client
}

func NewWhisperClient(apiKey string) *WhisperClient {
	return &WhisperClient{apiKey: apiKey, http: &http.Client{Timeout: 120 * time.Second}}
}

type whisperVerbose struct {
	Text     string  `json:"text"`
	Duration float64 `json:"duration"`
	Segments []struct {
		Start float64 `json:"start"`
		End   float64 `json:"end"`
	} `json:"segments"`
}

type Transcript struct {
	Text        string
	DurationSec float64
	Segments    []scoring.Segment
}

// Transcribe sends audio bytes and returns transcript + timings.
func (c *WhisperClient) Transcribe(ctx context.Context, audio []byte, filename string) (*Transcript, error) {
	var body bytes.Buffer
	w := multipart.NewWriter(&body)
	fw, err := w.CreateFormFile("file", filename)
	if err != nil {
		return nil, err
	}
	if _, err := fw.Write(audio); err != nil {
		return nil, err
	}
	_ = w.WriteField("model", whisperModel)
	_ = w.WriteField("response_format", "verbose_json")
	if err := w.Close(); err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, baseURL+"/audio/transcriptions", &body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", w.FormDataContentType())

	res, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	raw, _ := io.ReadAll(io.LimitReader(res.Body, 4<<20))
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("whisper: HTTP %d: %s", res.StatusCode, truncate(raw, 200))
	}
	var v whisperVerbose
	if err := json.Unmarshal(raw, &v); err != nil {
		return nil, err
	}
	tr := &Transcript{Text: v.Text, DurationSec: v.Duration}
	for _, s := range v.Segments {
		tr.Segments = append(tr.Segments, scoring.Segment{Start: s.Start, End: s.End})
	}
	return tr, nil
}

func truncate(b []byte, n int) string {
	if len(b) > n {
		return string(b[:n])
	}
	return string(b)
}
