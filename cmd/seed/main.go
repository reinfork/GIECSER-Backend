// Command seed provisions the teacher account and curriculum content.
// Speaking-first: passages, monologues, dialogues, pronunciation models and
// speaking tasks, plus comprehension quizzes where the book provides keys.
// Idempotent per chapter (title-keyed skip).
// Usage: TEACHER_EMAIL=... TEACHER_PASSWORD=... go run ./cmd/seed
// Revert Chapter 1 to supervisor FINAL: go run ./cmd/seed -reseed-ch1
// Revert Chapter 2 to supervisor FINAL: go run ./cmd/seed -reseed-ch2
// Revert Chapter 3 to supervisor FINAL: go run ./cmd/seed -reseed-ch3
package main

import (
	"flag"
	"log"

	"asri-backend/internal/config"
	"asri-backend/internal/model"
	"asri-backend/internal/repository"
	"asri-backend/internal/service"
)

func must(err error) {
	if err != nil {
		log.Fatalf("seed: %v", err)
	}
}

type seedTask struct {
	typ, prompt, phonetics, options, answer string
	order                                   int
}

type seedModule struct {
	title, typ, content, media, transcript, audio string
	order                                         int
	tasks                                         []seedTask
}

type seedCourse struct {
	title, desc string
	modules     []seedModule
}

type seedChapter struct {
	title, desc string
	courses     []seedCourse
}

func speakTask(prompt, phonetics string) seedTask {
	return seedTask{typ: model.TaskSpeakingRecording, prompt: prompt, phonetics: phonetics}
}

// quizTask carries machine-checkable keys. Conventions:
// MULTIPLE_CHOICE: options {"A":"..","B":".."}, answer {"correct":"B"}.
// MATCHING: options {"statements":{"1":".."},"endings":{"A":".."}}, answer {"1":"B",...}.
// OPEN_ENDED: answer {"answers"|"blanks"|"order":[...]}.
func quizTask(typ, prompt, options, answer string) seedTask {
	return seedTask{typ: typ, prompt: prompt, options: options, answer: answer}
}

type stores struct {
	chapter *repository.ChapterRepository
	course  *repository.CourseRepository
	module  *repository.ModuleRepository
	task    *repository.TaskRepository
}

func insertChapter(st stores, ch seedChapter, order int) (courses, modules, tasks int) {
	chapter := &model.Chapter{Title: ch.title, Description: ch.desc, OrderIndex: order}
	must(st.chapter.Create(chapter))
	return insertCourses(st, chapter.ID, ch.courses)
}

// insertCourses (re)populates courses under an existing chapter row.
// Reseeds use this so a re-run never duplicates the chapter itself.
func insertCourses(st stores, chapterID string, courses_ []seedCourse) (courses, modules, tasks int) {
	for j, co := range courses_ {
		course := &model.Course{ChapterID: chapterID, Title: co.title, Description: co.desc, OrderIndex: j + 1}
		must(st.course.Create(course))
		courses++
		for k, m := range co.modules {
			module := &model.Module{
				CourseID: course.ID, Title: m.title, Type: m.typ,
				ContentText: m.content, MediaURL: m.media, TargetTranscript: m.transcript,
				AudioModelURL: m.audio, OrderIndex: k + 1,
			}
			must(st.module.Create(module))
			modules++
			for l, t := range m.tasks {
				must(st.task.Create(&model.Task{
					ModuleID: module.ID, Type: t.typ, Prompt: t.prompt,
					OptionsJSON: t.options, AnswerKeyJSON: t.answer,
					TargetPhonetics: t.phonetics, OrderIndex: l + 1,
				}))
				tasks++
			}
		}
	}
	return courses, modules, tasks
}

func main() {
	reseed := flag.Bool("reseed-ch1", false, "delete and re-seed the Traditional Games chapter")
	reseedCh2 := flag.Bool("reseed-ch2", false, "delete and re-seed the Medicinal Plants chapter")
	reseedCh3 := flag.Bool("reseed-ch3", false, "delete and re-seed the Maritime Transport chapter")
	reseedCh4 := flag.Bool("reseed-ch4", false, "delete and re-seed the Traditional Attire chapter")
	flag.Parse()

	cfg := config.Load()
	db, err := config.InitDB(cfg.DSN())
	must(err)

	authSvc := service.NewAuthService(repository.NewUserRepository(db))
	teacher, err := authSvc.SeedTeacher(cfg.TeacherEmail, cfg.TeacherPassword)
	must(err)
	if teacher == nil {
		log.Fatal("seed: set TEACHER_EMAIL and TEACHER_PASSWORD (created_by linkage)")
	}

	st := stores{
		chapter: repository.NewChapterRepository(db),
		course:  repository.NewCourseRepository(db),
		module:  repository.NewModuleRepository(db),
		task:    repository.NewTaskRepository(db),
	}

	if *reseed {
		var ch model.Chapter
		if err := db.Where("title = ?", "Traditional Games").First(&ch).Error; err != nil {
			log.Fatalf("seed: chapter not found: %v", err)
		}
		// ponytail: raw bottom-up deletes — GORM has no cascade here; a model
		// method would hide the blast radius. Scoped to one chapter by id.
		must(db.Exec(`DELETE FROM tasks WHERE module_id IN (SELECT m.id FROM modules m JOIN courses c ON c.id = m.course_id WHERE c.chapter_id = ?)`, ch.ID).Error)
		must(db.Exec(`DELETE FROM modules WHERE course_id IN (SELECT id FROM courses WHERE chapter_id = ?)`, ch.ID).Error)
		must(db.Exec(`DELETE FROM courses WHERE chapter_id = ?`, ch.ID).Error)
		co, mo, ta := insertCourses(st, ch.ID, ch1Final().courses)
		db.Model(&model.Chapter{}).Where("id = ?", ch.ID).Update("description", ch1Final().desc)
		log.Printf("seed: re-seeded Traditional Games: %d courses, %d modules, %d tasks", co, mo, ta)
		return
	}

	if *reseedCh2 {
		var ch model.Chapter
		if err := db.Where("title = ?", "Medicinal Plants").First(&ch).Error; err != nil {
			log.Fatalf("seed: chapter not found: %v", err)
		}
		must(db.Exec(`DELETE FROM tasks WHERE module_id IN (SELECT m.id FROM modules m JOIN courses c ON c.id = m.course_id WHERE c.chapter_id = ?)`, ch.ID).Error)
		must(db.Exec(`DELETE FROM modules WHERE course_id IN (SELECT id FROM courses WHERE chapter_id = ?)`, ch.ID).Error)
		must(db.Exec(`DELETE FROM courses WHERE chapter_id = ?`, ch.ID).Error)
		co, mo, ta := insertCourses(st, ch.ID, ch2Final().courses)
		db.Model(&model.Chapter{}).Where("id = ?", ch.ID).Update("description", ch2Final().desc)
		log.Printf("seed: re-seeded Medicinal Plants: %d courses, %d modules, %d tasks", co, mo, ta)
		return
	}

	if *reseedCh3 {
		var ch model.Chapter
		if err := db.Where("title = ?", "Maritime Transport").First(&ch).Error; err != nil {
			log.Fatalf("seed: chapter not found: %v", err)
		}
		must(db.Exec(`DELETE FROM tasks WHERE module_id IN (SELECT m.id FROM modules m JOIN courses c ON c.id = m.course_id WHERE c.chapter_id = ?)`, ch.ID).Error)
		must(db.Exec(`DELETE FROM modules WHERE course_id IN (SELECT id FROM courses WHERE chapter_id = ?)`, ch.ID).Error)
		must(db.Exec(`DELETE FROM courses WHERE chapter_id = ?`, ch.ID).Error)
		co, mo, ta := insertCourses(st, ch.ID, ch3Final().courses)
		db.Model(&model.Chapter{}).Where("id = ?", ch.ID).Update("description", ch3Final().desc)
		log.Printf("seed: re-seeded Maritime Transport: %d courses, %d modules, %d tasks", co, mo, ta)
		return
	}

	if *reseedCh4 {
		var ch model.Chapter
		if err := db.Where("title = ?", "Traditional Attire").First(&ch).Error; err != nil {
			log.Fatalf("seed: chapter not found: %v", err)
		}
		must(db.Exec(`DELETE FROM tasks WHERE module_id IN (SELECT m.id FROM modules m JOIN courses c ON c.id = m.course_id WHERE c.chapter_id = ?)`, ch.ID).Error)
		must(db.Exec(`DELETE FROM modules WHERE course_id IN (SELECT id FROM courses WHERE chapter_id = ?)`, ch.ID).Error)
		must(db.Exec(`DELETE FROM courses WHERE chapter_id = ?`, ch.ID).Error)
		co, mo, ta := insertCourses(st, ch.ID, ch4Final().courses)
		db.Model(&model.Chapter{}).Where("id = ?", ch.ID).Update("description", ch4Final().desc)
		log.Printf("seed: re-seeded Traditional Attire: %d courses, %d modules, %d tasks", co, mo, ta)
		return
	}

	// ponytail: title-keyed skip lets content grow chapter-by-chapter;
	// replace with versioned migrations when edits (not just additions) are needed.
	var existing []model.Chapter
	db.Model(&model.Chapter{}).Find(&existing)
	present := map[string]bool{}
	for _, c := range existing {
		present[c.Title] = true
	}

	nCo, nMo, nTa := 0, 0, 0
	for i, ch := range chapters() {
		if present[ch.title] {
			log.Printf("seed: chapter %q already present, skipping", ch.title)
			continue
		}
		co, mo, ta := insertChapter(st, ch, i+1)
		nCo, nMo, nTa = nCo+co, nMo+mo, nTa+ta
	}
	log.Printf("seed: done — %d courses, %d modules, %d tasks created this run", nCo, nMo, nTa)
}
