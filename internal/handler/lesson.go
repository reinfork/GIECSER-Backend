package handler

import (
	"errors"
	"net/http"
	"strconv"

	"asri-backend/internal/model"
	"asri-backend/internal/service"

	"github.com/gin-gonic/gin"
)

type LessonHandler struct {
	service *service.LessonService
}

func NewLessonHandler(service *service.LessonService) *LessonHandler {
	return &LessonHandler{service: service}
}

func (h *LessonHandler) Create(c *gin.Context) {
	var input model.CreateLessonInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	lesson, err := h.service.Create(input)
	if err != nil {
		if errors.Is(err, service.ErrCourseNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "course not found"})
			return
		}
		if errors.Is(err, service.ErrConflict) {
			c.JSON(http.StatusConflict, gin.H{"error": "order_index already exists for this course"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, lesson)
}

func (h *LessonHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	lesson, err := h.service.GetByID(uint(id))
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "lesson not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, lesson)
}

func (h *LessonHandler) ListByCourse(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid course id"})
		return
	}
	// Pagination support ?page=&limit=
	if c.Query("page") != "" || c.Query("limit") != "" {
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
		lessons, total, err := h.service.ListByCoursePaginated(uint(id), page, limit)
		if err != nil {
			if errors.Is(err, service.ErrCourseNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "course not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": lessons, "meta": gin.H{"total": total, "page": page, "limit": limit}})
		return
	}
	lessons, err := h.service.ListByCourse(uint(id))
	if err != nil {
		if errors.Is(err, service.ErrCourseNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "course not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": lessons, "meta": gin.H{"total": len(lessons), "page": 1, "limit": len(lessons)}})
}

func (h *LessonHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var input model.UpdateLessonInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	lesson, err := h.service.Update(uint(id), input)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "lesson not found"})
			return
		}
		if errors.Is(err, service.ErrConflict) {
			c.JSON(http.StatusConflict, gin.H{"error": "order_index already exists for this course"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, lesson)
}

func (h *LessonHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := h.service.Delete(uint(id)); err != nil {
		if errors.Is(err, service.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "lesson not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "lesson deleted"})
}

func (h *LessonHandler) Practice(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var input model.PracticeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := h.service.Practice(uint(id), input.Transcript)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "lesson not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
