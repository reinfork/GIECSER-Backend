package scoring

import (
	"math"
	"testing"
)

func TestWER(t *testing.T) {
	cases := []struct {
		name     string
		ref, hyp string
		wantWER  float64
		wantS    int
		wantD    int
		wantI    int
	}{
		{"identical", "i eat rice", "i eat rice", 0, 0, 0, 0},
		{"one substitution", "i eat rice", "i eat bread", 1.0 / 3, 1, 0, 0},
		{"one deletion", "i eat rice", "i eat", 1.0 / 3, 0, 1, 0},
		{"one insertion", "i eat", "i eat rice", 0.5, 0, 0, 1},
		{"empty transcript", "i eat rice", "", 1, 0, 3, 0},
		{"both empty", "", "", 0, 0, 0, 0},
		{"case and punctuation ignored", "I eat, rice!", "i EAT rice", 0, 0, 0, 0},
		{"malay word kept", "saya makan satay", "saya makan satay", 0, 0, 0, 0},
		{"malay misheard", "saya makan satay", "saya makan sate", 1.0 / 3, 1, 0, 0},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			wer, s, d, i, n := WER(tc.ref, tc.hyp)
			if math.Abs(wer-tc.wantWER) > 1e-9 {
				t.Errorf("WER = %v, want %v", wer, tc.wantWER)
			}
			if s != tc.wantS || d != tc.wantD || i != tc.wantI {
				t.Errorf("S/D/I = %d/%d/%d, want %d/%d/%d", s, d, i, tc.wantS, tc.wantD, tc.wantI)
			}
			_ = n
		})
	}
}

func TestAccuracy(t *testing.T) {
	if got := Accuracy(0); got != 100 {
		t.Errorf("Accuracy(0) = %v, want 100", got)
	}
	if got := Accuracy(1); got != 0 {
		t.Errorf("Accuracy(1) = %v, want 0", got)
	}
	if got := Accuracy(2); got != 0 {
		t.Errorf("Accuracy(2) = %v, want 0 (clamped)", got)
	}
}

func TestWPM(t *testing.T) {
	if got := WPM(130, 60); got != 130 {
		t.Errorf("WPM = %v, want 130", got)
	}
	if got := WPM(10, 0); got != 0 {
		t.Errorf("WPM zero-duration = %v, want 0", got)
	}
}

func TestFluencyScore(t *testing.T) {
	for _, wpm := range []float64{110, 130, 150} {
		if got := FluencyScore(wpm); got != 100 {
			t.Errorf("FluencyScore(%v) = %v, want 100", wpm, got)
		}
	}
	if got := FluencyScore(0); got != 0 {
		t.Errorf("FluencyScore(0) = %v, want 0", got)
	}
}

func TestSilenceStats(t *testing.T) {
	segs := []Segment{{Start: 0, End: 2}, {Start: 3, End: 5}} // 1s gap > 0.5s
	pause, active := SilenceStats(segs, 5)
	if pause != 1 || active != 4 {
		t.Errorf("pause/active = %v/%v, want 1/4", pause, active)
	}
}

func TestSyllables(t *testing.T) {
	cases := map[string]int{"rice": 1, "satay": 2, "english": 2, "a": 1}
	for w, want := range cases {
		if got := Syllables(w); got != want {
			t.Errorf("Syllables(%q) = %v, want %v", w, got, want)
		}
	}
}

func TestAlign(t *testing.T) {
	cases := []struct {
		name     string
		ref, hyp string
		want     []string
	}{
		{"perfect", "root sap wound", "root sap wound", nil},
		{"substitution", "jong layo is a game", "jong layo is uh games", []string{"a", "game"}},
		{"deletion", "i eat rice", "i eat", []string{"rice"}},
		{"insertion ignored", "i eat", "i eat rice", nil},
		{"deduped in order", "a a b", "x x b", []string{"a"}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Align(tc.ref, tc.hyp)
			if len(got) != len(tc.want) {
				t.Fatalf("Align = %v, want %v", got, tc.want)
			}
			for i := range got {
				if got[i] != tc.want[i] {
					t.Fatalf("Align = %v, want %v", got, tc.want)
				}
			}
		})
	}
}
