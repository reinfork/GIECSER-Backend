package service

import (
	"errors"

	"asri-backend/internal/model"

	"gorm.io/gorm"
)

type TaskRepo interface {
	FindByModule(moduleID string) ([]model.Task, error)
	FindByID(id string) (*model.Task, error)
	Create(t *model.Task) error
	Update(t *model.Task) error
	Delete(id string) error
}

type TaskService struct {
	repo     TaskRepo
	modules  ModuleRepo
	courses  CourseRepo
	chapters ChapterStore
}

func NewTaskService(repo TaskRepo, modules ModuleRepo, courses CourseRepo, chapters ChapterStore) *TaskService {
	return &TaskService{repo: repo, modules: modules, courses: courses, chapters: chapters}
}

func (s *TaskService) ListByModule(moduleID string) ([]model.Task, error) {
	return s.repo.FindByModule(moduleID)
}

// ListPublicByModule strips answer keys — the only shape session tokens may see.
func (s *TaskService) ListPublicByModule(moduleID string) ([]model.TaskPublic, error) {
	tasks, err := s.repo.FindByModule(moduleID)
	if err != nil {
		return nil, err
	}
	out := make([]model.TaskPublic, 0, len(tasks))
	for _, t := range tasks {
		out = append(out, t.Public())
	}
	return out, nil
}

func (s *TaskService) GetByID(id string) (*model.Task, error) {
	task, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return task, nil
}

func (s *TaskService) Create(input model.CreateTaskInput) (*model.Task, error) {
	if _, err := s.modules.FindByID(input.ModuleID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	task := &model.Task{
		ModuleID: input.ModuleID, Type: input.Type, Prompt: input.Prompt,
		OptionsJSON: input.OptionsJSON, AnswerKeyJSON: input.AnswerKeyJSON,
		TargetPhonetics: input.TargetPhonetics, OrderIndex: input.OrderIndex,
	}
	if err := s.repo.Create(task); err != nil {
		return nil, err
	}
	return task, nil
}

func (s *TaskService) Update(id string, input model.UpdateTaskInput) (*model.Task, error) {
	task, err := s.GetByID(id)
	if err != nil {
		return nil, err
	}
	if input.Type != nil {
		task.Type = *input.Type
	}
	if input.Prompt != nil {
		task.Prompt = *input.Prompt
	}
	if input.OptionsJSON != nil {
		task.OptionsJSON = *input.OptionsJSON
	}
	if input.AnswerKeyJSON != nil {
		task.AnswerKeyJSON = *input.AnswerKeyJSON
	}
	if input.TargetPhonetics != nil {
		task.TargetPhonetics = *input.TargetPhonetics
	}
	if input.OrderIndex != nil {
		task.OrderIndex = *input.OrderIndex
	}
	if err := s.repo.Update(task); err != nil {
		return nil, err
	}
	return task, nil
}

func (s *TaskService) Delete(id string) error {
	if _, err := s.GetByID(id); err != nil {
		return err
	}
	return s.repo.Delete(id)
}

// ContextForTask resolves the task plus the chapter title for Groq cultural
// grounding. No scope involved: one code opens the whole book.
func (s *TaskService) ContextForTask(id string) (task *model.Task, chapterTitle string, err error) {
	task, err = s.GetByID(id)
	if err != nil {
		return nil, "", err
	}
	var module *model.Module
	// modules reader is the ModuleRepo interface: FindByID exists on it.
	module, err = s.modules.FindByID(task.ModuleID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "", ErrNotFound
		}
		return nil, "", err
	}
	course, err := s.courses.FindByID(module.CourseID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "", ErrNotFound
		}
		return nil, "", err
	}
	chapter, err := s.chapters.FindByID(course.ChapterID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "", ErrNotFound
		}
		return nil, "", err
	}
	return task, chapter.Title, nil
}
