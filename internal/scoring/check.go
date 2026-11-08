package scoring

import (
	"encoding/json"
	"strconv"
	"strings"
)

// Check compares a student attempt against a task's answer_key_json.
// Pure function — no DB, no Groq — so it unit-tests like WER.
func Check(taskType, keyJSON string, answer any) (correct bool, detail map[string]bool) {
	norm := func(s string) string {
		s = strings.ToLower(strings.TrimSpace(s))
		s = strings.Trim(s, ".")
		return strings.Join(strings.Fields(s), " ")
	}
	switch taskType {
	case "MULTIPLE_CHOICE":
		var k struct {
			Correct string `json:"correct"`
		}
		if json.Unmarshal([]byte(keyJSON), &k) != nil {
			return false, nil
		}
		if s, ok := answer.(string); ok {
			return norm(s) == norm(k.Correct), nil
		}
		return false, nil
	case "MATCHING":
		var k map[string]string
		if json.Unmarshal([]byte(keyJSON), &k) != nil {
			return false, nil
		}
		got, ok := answer.(map[string]any)
		if !ok {
			return false, nil
		}
		detail = map[string]bool{}
		all := true
		for key, want := range k {
			hit := false
			if v, ok := got[key]; ok {
				if s, ok := v.(string); ok {
					hit = norm(s) == norm(want)
				}
			}
			detail[key] = hit
			all = all && hit
		}
		return all, detail
	default: // OPEN_ENDED: answers (any-of), blanks (per-index), order (sequence)
		var k struct {
			Answers []string `json:"answers"`
			Blanks  []string `json:"blanks"`
			Order   []string `json:"order"`
		}
		if json.Unmarshal([]byte(keyJSON), &k) != nil {
			return false, nil
		}
		if len(k.Blanks) > 0 {
			arr, ok := answer.([]any)
			if !ok || len(arr) != len(k.Blanks) {
				return false, nil
			}
			detail = map[string]bool{}
			all := true
			for i, want := range k.Blanks {
				hit := false
				if s, ok := arr[i].(string); ok {
					hit = norm(s) == norm(want)
				}
				// ponytail: positional index keys — real per-blank IDs when UI needs them.
				detail[strconv.Itoa(i)] = hit
				all = all && hit
			}
			return all, detail
		}
		if len(k.Order) > 0 {
			arr, ok := answer.([]any)
			if !ok || len(arr) != len(k.Order) {
				return false, nil
			}
			for i, want := range k.Order {
				s, ok := arr[i].(string)
				if !ok || norm(s) != norm(want) {
					return false, nil
				}
			}
			return true, nil
		}
		if s, ok := answer.(string); ok {
			for _, want := range k.Answers {
				if norm(s) == norm(want) {
					return true, nil
				}
			}
		}
		return false, nil
	}
}
