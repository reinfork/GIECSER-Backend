package handler

import (
	"errors"
	"net/http"

	"asri-backend/internal/model"
	"asri-backend/internal/service"

	"github.com/gin-gonic/gin"
)

type ModuleHandler struct {
	service *service.ModuleService
}

func NewModuleHandler(service *service.ModuleService) *ModuleHandler {
	return &ModuleHandler{service: service}
}

// ListByCourse lists a course's modules. Any valid session sees everything.
func (h *ModuleHandler) ListByCourse(c *gin.Context) {
	courseID := c.Param("id")
	modules, err := h.service.ListByCourse(courseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": modules})
}

func (h *ModuleHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	module, err := h.service.GetByID(id)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "module not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, module)
}

func (h *ModuleHandler) Create(c *gin.Context) {
	var input model.CreateModuleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	module, err := h.service.Create(input)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "course not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, module)
}

func (h *ModuleHandler) Update(c *gin.Context) {
	var input model.UpdateModuleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	module, err := h.service.Update(c.Param("id"), input)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "module not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, module)
}

func (h *ModuleHandler) Delete(c *gin.Context) {
	if err := h.service.Delete(c.Param("id")); err != nil {
		if errors.Is(err, service.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "module not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "module deleted"})
}
