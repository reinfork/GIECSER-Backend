// chapters() aggregates the full chapter seeders, all frozen from GIECSER_FINAL.
// Chapter 1 is seeded separately (teacher-created); Chapters 2-4 live in
// chapter2.go / chapter3.go / chapter4.go.
package main

func chapters() []seedChapter {
	return []seedChapter{
		ch2Final(),
		ch3Final(),
		ch4Final(),
	}
}
