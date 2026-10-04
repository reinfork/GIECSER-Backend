package handler

import (
	"errors"
	"net/http"

	"asri-backend/internal/model"
	"asri-backend/internal/service"

	"github.com/gin-gonic/gin"
)

type ChapterHandler struct {
	service *service.ChapterService
}

func NewChapterHandler(service *service.ChapterService) *ChapterHandler {
	return &ChapterHandler{service: service}
}

// List returns all chapters. One classroom code opens the whole book.
func (h *ChapterHandler) List(c *gin.Context) {
	chapters, err := h.service.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": chapters})
}

func (h *ChapterHandler) GetByID(c *gin.Context) {
	chapter, err := h.service.GetByID(c.Param("id"))
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "chapter not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, chapter)
}

func (h *ChapterHandler) Create(c *gin.Context) {
	var input model.CreateChapterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	chapter, err := h.service.Create(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, chapter)
}

func (h *ChapterHandler) Update(c *gin.Context) {
	var input model.UpdateChapterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	chapter, err := h.service.Update(c.Param("id"), input)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "chapter not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, chapter)
}

func (h *ChapterHandler) Delete(c *gin.Context) {
	if err := h.service.Delete(c.Param("id")); err != nil {
		if errors.Is(err, service.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "chapter not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "chapter deleted"})
}
