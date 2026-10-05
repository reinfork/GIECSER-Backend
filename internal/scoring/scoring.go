// Package scoring holds dependency-free speech metrics (PRD §3).
// Pure functions: unit-tested, no network, no models.
package scoring

import (
	"math"
	"strings"
	"unicode"
)

// Words lowercases and splits on non-letters/digits. Malay + English safe.
func Words(s string) []string {
	return strings.FieldsFunc(strings.ToLower(s), func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	})
}

// WER returns Word Error Rate = (S+D+I)/N via Levenshtein on word slices,
// plus the raw substitution/deletion/insertion counts and N.
func WER(reference, transcript string) (wer float64, s, d, i, n int) {
	ref, hyp := Words(reference), Words(transcript)
	n = len(ref)
	if n == 0 {
		if len(hyp) == 0 {
			return 0, 0, 0, 0, 0
		}
		return 1, 0, 0, len(hyp), 0
	}
	prev := make([]int, len(hyp)+1)
	curr := make([]int, len(hyp)+1)
	for j := range prev {
		prev[j] = j
	}
	for a := 1; a <= n; a++ {
		curr[0] = a
		for b := 1; b <= len(hyp); b++ {
			cost := 0
			if ref[a-1] != hyp[b-1] {
				cost = 1
			}
			curr[b] = min3(prev[b]+1, curr[b-1]+1, prev[b-1]+cost)
		}
		prev, curr = curr, prev
	}
	dist := prev[len(hyp)]
	// Backtrack-lite: count ops by re-walking is O(n*m) memory; instead derive
	// S/D/I approximately is wrong — do the full matrix once (clips are short).
	return fullWER(ref, hyp, dist)
}

func fullWER(ref, hyp []string, dist int) (float64, int, int, int, int) {
	n, m := len(ref), len(hyp)
	dp := make([][]int, n+1)
	for a := range dp {
		dp[a] = make([]int, m+1)
		dp[a][0] = a
	}
	for b := 0; b <= m; b++ {
		dp[0][b] = b
	}
	for a := 1; a <= n; a++ {
		for b := 1; b <= m; b++ {
			cost := 0
			if ref[a-1] != hyp[b-1] {
				cost = 1
			}
			dp[a][b] = min3(dp[a-1][b]+1, dp[a][b-1]+1, dp[a-1][b-1]+cost)
		}
	}
	var s, d, ins int
	for a, b := n, m; a > 0 || b > 0; {
		switch {
		case a > 0 && b > 0 && ref[a-1] == hyp[b-1]:
			a, b = a-1, b-1
		case a > 0 && b > 0 && dp[a][b] == dp[a-1][b-1]+1:
			s, a, b = s+1, a-1, b-1
		case b > 0 && dp[a][b] == dp[a][b-1]+1:
			ins, b = ins+1, b-1
		default:
			d, a = d+1, a-1
		}
	}
	_ = dist
	return float64(s+d+ins) / float64(n), s, d, ins, n
}

func min3(a, b, c int) int {
	return min(a, min(b, c))
}

// Align returns reference words hit by substitution or deletion, deduped in
// order. Same normalization and backtrack priority as WER, so chips agree
// with Accuracy by construction.
// ponytail: O(n*m) DP per submit — clips are short; Hirschberg when they aren't.
func Align(reference, transcript string) []string {
	ref, hyp := Words(reference), Words(transcript)
	n, m := len(ref), len(hyp)
	dp := make([][]int, n+1)
	for a := range dp {
		dp[a] = make([]int, m+1)
		dp[a][0] = a
	}
	for b := 0; b <= m; b++ {
		dp[0][b] = b
	}
	for a := 1; a <= n; a++ {
		for b := 1; b <= m; b++ {
			cost := 0
			if ref[a-1] != hyp[b-1] {
				cost = 1
			}
			dp[a][b] = min3(dp[a-1][b]+1, dp[a][b-1]+1, dp[a-1][b-1]+cost)
		}
	}
	var bad []string
	seen := map[string]bool{}
	for a, b := n, m; a > 0 || b > 0; {
		switch {
		case a > 0 && b > 0 && ref[a-1] == hyp[b-1]:
			a, b = a-1, b-1
		case a > 0 && b > 0 && dp[a][b] == dp[a-1][b-1]+1:
			if !seen[ref[a-1]] {
				seen[ref[a-1]] = true
				bad = append([]string{ref[a-1]}, bad...)
			}
			a, b = a-1, b-1
		case b > 0 && dp[a][b] == dp[a][b-1]+1:
			b--
		default:
			if !seen[ref[a-1]] {
				seen[ref[a-1]] = true
				bad = append([]string{ref[a-1]}, bad...)
			}
			a--
		}
	}
	return bad
}

// Accuracy derives PRD §3.1: max(0, 100*(1-WER)).
func Accuracy(wer float64) float64 {
	return math.Max(0, 100*(1-wer))
}

// WPM is words per minute of speech (PRD target 110–150 for L2 learners).
func WPM(wordCount int, durationSec float64) float64 {
	if durationSec <= 0 {
		return 0
	}
	return float64(wordCount) / durationSec * 60
}

// Segment is one transcribed span with timings (from Whisper verbose_json).
type Segment struct {
	Start float64
	End   float64
}

// SilenceStats splits total time into paused (>0.5s gaps) vs active speech.
func SilenceStats(segments []Segment, totalSec float64) (pauseSec, activeSec float64) {
	var cursor float64
	for _, sg := range segments {
		if gap := sg.Start - cursor; gap > 0.5 {
			pauseSec += gap
		}
		if dur := sg.End - sg.Start; dur > 0 {
			activeSec += dur
		}
		cursor = sg.End
	}
	if tail := totalSec - cursor; tail > 0.5 {
		pauseSec += tail
	}
	return pauseSec, activeSec
}

// Syllables estimates via vowel groups — approximate, not linguistic truth.
// ponytail: naive estimator; swap for a phoneme aligner when pronunciation research demands it.
func Syllables(word string) int {
	word = strings.ToLower(word)
	count, prevVowel := 0, false
	isVowel := func(r rune) bool { return strings.ContainsRune("aeiou", r) }
	for _, r := range word {
		v := isVowel(r)
		if v && !prevVowel {
			count++
		}
		prevVowel = v
	}
	if strings.HasSuffix(word, "e") && count > 1 {
		count--
	}
	return max(count, 1)
}

// ArticulationRate is syllables per active second (PRD target 3.5–4.8).
func ArticulationRate(words []string, activeSec float64) float64 {
	if activeSec <= 0 {
		return 0
	}
	syll := 0
	for _, w := range words {
		syll += Syllables(w)
	}
	return float64(syll) / activeSec
}

// FluencyScore maps WPM onto 0–100 against the 110–150 ideal band.
func FluencyScore(wpm float64) float64 {
	const lo, hi = 110.0, 150.0
	if wpm >= lo && wpm <= hi {
		return 100
	}
	dist := math.Min(math.Abs(wpm-lo), math.Abs(wpm-hi))
	return math.Max(0, 100-dist)
}
