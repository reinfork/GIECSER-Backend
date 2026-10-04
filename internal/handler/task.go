package handler

import (
	"errors"
	"io"
	"net/http"
	"strings"

	"asri-backend/internal/model"
	"asri-backend/internal/service"

	"github.com/gin-gonic/gin"
)

// 5MB / short clips only: sentences, not lectures (pedagogy + queue math agree).
const maxAudioBytes = 5 << 20

var allowedAudioCT = map[string]bool{
	"audio/wav": true, "audio/x-wav": true, "audio/mpeg": true,
	"audio/mp4": true, "audio/ogg": true, "audio/webm": true,
}

type TaskHandler struct {
	tasks       *service.TaskService
	submissions *service.SubmissionService
}

func NewTaskHandler(tasks *service.TaskService, submissions *service.SubmissionService) *TaskHandler {
	return &TaskHandler{tasks: tasks, submissions: submissions}
}

// ListByModule returns tasks with answer keys stripped (student-safe).
// Unknown module id yields an empty list.
func (h *TaskHandler) ListByModule(c *gin.Context) {
	moduleID := c.Param("id")
	tasks, err := h.tasks.ListPublicByModule(moduleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

func (h *TaskHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	task, _, err := h.tasks.ContextForTask(id)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task.Public())
}

// SubmitAudio runs the PRD §4 pipeline. Temp file deleted before return (no-storage ruling).
func (h *TaskHandler) SubmitAudio(c *gin.Context) {
	taskID := c.Param("id")
	_, chapterTitle, err := h.tasks.ContextForTask(taskID)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fh, err := c.FormFile("audio")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "audio file required (multipart field 'audio')"})
		return
	}
	if fh.Size > maxAudioBytes {
		c.JSON(http.StatusRequestEntityTooLarge, gin.H{"error": "audio too large, keep clips under a minute"})
		return
	}
	f, err := fh.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot read audio"})
		return
	}
	defer f.Close()
	audio, err := io.ReadAll(io.LimitReader(f, maxAudioBytes+1))
	if err != nil || len(audio) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "empty audio"})
		return
	}
	if ct := fh.Header.Get("Content-Type"); ct != "" && !allowedAudioCT[strings.Split(ct, ";")[0]] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unsupported audio format"})
		return
	}

	result, err := h.submissions.SubmitAudio(c.Request.Context(), taskID, audio, fh.Filename, chapterTitle)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
			return
		}
		if errors.Is(err, service.ErrUnsupported) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "task does not accept audio"})
			return
		}
		c.JSON(http.StatusBadGateway, gin.H{"error": "assessment failed, try again"})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *TaskHandler) Create(c *gin.Context) {
	var input model.CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task, err := h.tasks.Create(input)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "module not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) Update(c *gin.Context) {
	var input model.UpdateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task, err := h.tasks.Update(c.Param("id"), input)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) Delete(c *gin.Context) {
	if err := h.tasks.Delete(c.Param("id")); err != nil {
		if errors.Is(err, service.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "task deleted"})
}
