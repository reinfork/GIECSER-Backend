package scoring

import "testing"

func TestCheckMCQ(t *testing.T) {
	if ok, _ := Check("MULTIPLE_CHOICE", `{"correct":"B"}`, "b"); !ok {
		t.Fatal("case-insensitive letter should match")
	}
	if ok, _ := Check("MULTIPLE_CHOICE", `{"correct":"B"}`, "C"); ok {
		t.Fatal("wrong letter should not match")
	}
}

func TestCheckMatching(t *testing.T) {
	ok, d := Check("MATCHING", `{"1":"B","2":"E"}`, map[string]any{"1": "b", "2": "C"})
	if ok || !d["1"] || d["2"] {
		t.Fatalf("per-item detail wrong: ok=%v detail=%v", ok, d)
	}
	if ok, _ := Check("MATCHING", `{"1":"B"}`, "B"); ok {
		t.Fatal("non-map answer should not match")
	}
}

func TestCheckOpenEnded(t *testing.T) {
	if ok, _ := Check("OPEN_ENDED", `{"answers":["Riau"]}`, "  riau. "); !ok {
		t.Fatal("normalized answer should match")
	}
	if ok, _ := Check("OPEN_ENDED", `{"blanks":["Riau","Water"]}`, []any{"riau", "wind"}); ok {
		t.Fatal("one wrong blank should fail")
	}
	if ok, _ := Check("OPEN_ENDED", `{"order":["B","F"]}`, []any{"B", "F"}); !ok {
		t.Fatal("exact sequence should match")
	}
}
