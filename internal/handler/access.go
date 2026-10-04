package handler

import (
	"errors"
	"net/http"

	"asri-backend/internal/model"
	"asri-backend/internal/service"

	"github.com/gin-gonic/gin"
)

type AccessHandler struct {
	service *service.AccessService
}

func NewAccessHandler(service *service.AccessService) *AccessHandler {
	return &AccessHandler{service: service}
}

// ValidatePIN exchanges a 6-char classroom code for a session JWT
// opening the whole book.
func (h *AccessHandler) ValidatePIN(c *gin.Context) {
	var input model.ValidatePINInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, expiresAt, err := h.service.ValidatePIN(c.Request.Context(), input.Code)
	if err != nil {
		if errors.Is(err, service.ErrUnauthorized) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired code"})
			return
		}
		if errors.Is(err, service.ErrExpired) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "code expired, ask your teacher for a new one"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token, "expires_at": expiresAt})
}

// IssuePIN mints a code opening the whole book (teacher only).
func (h *AccessHandler) IssuePIN(c *gin.Context) {
	var input model.IssuePINInput
	_ = c.ShouldBindJSON(&input) // TTL optional
	var teacherID *string
	if id, ok := c.Get("userID"); ok {
		if s, ok := id.(string); ok {
			teacherID = &s
		}
	}
	ac, err := h.service.IssuePIN(c.Request.Context(), input.TTLHours, teacherID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, ac)
}

// ListPINs shows active codes (teacher only).
func (h *AccessHandler) ListPINs(c *gin.Context) {
	codes, err := h.service.ListActivePINs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": codes})
}
