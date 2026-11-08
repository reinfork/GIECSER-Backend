package service

import (
	"testing"
)

func TestReferenceFor(t *testing.T) {
	mono := "Riau coastal traditional games have been part of Malay community life"
	model := `{"Model answer":"Riau coastal traditional games are played by Malay communities."}`

	if got := referenceFor("What are Riau coastal traditional games?", model, mono); got != "Riau coastal traditional games are played by Malay communities." {
		t.Fatalf("question should score against Model answer, got %q", got)
	}
	if got := referenceFor("Read the monologue aloud.", "", mono); got != mono {
		t.Fatalf("repeat should score against transcript, got %q", got)
	}
	if got := referenceFor("What is it?", `{}`, mono); got != mono {
		t.Fatalf("question without key falls back to transcript, got %q", got)
	}
	if got := referenceFor("Say something.", "", ""); got != "Say something." {
		t.Fatalf("bare prompt falls back to itself, got %q", got)
	}
}

func TestBoost(t *testing.T) {
	if got := boost(70, 10); got != 80 {
		t.Fatalf("got %v", got)
	}
	if got := boost(95, 10); got != 100 {
		t.Fatalf("must cap at 100, got %v", got)
	}
	if got := textDelta(); got != 0 {
		t.Fatalf("default delta must be 0, got %v", got)
	}
	t.Setenv("SUBMIT_TEXT_DELTA", "abc")
	if got := textDelta(); got != 0 {
		t.Fatalf("invalid delta must be 0, got %v", got)
	}
	t.Setenv("SUBMIT_TEXT_DELTA", "-5")
	if got := textDelta(); got != 0 {
		t.Fatalf("negative delta must be 0, got %v", got)
	}
	t.Setenv("SUBMIT_TEXT_DELTA", "7.5")
	if got := textDelta(); got != 7.5 {
		t.Fatalf("got %v", got)
	}
}
