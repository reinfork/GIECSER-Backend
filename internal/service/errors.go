package service

import "errors"

var (
	ErrNotFound       = errors.New("not found")
	ErrConflict       = errors.New("conflict")
	ErrUnauthorized   = errors.New("unauthorized")
	ErrCourseNotFound = errors.New("course not found")
	ErrExpired        = errors.New("expired")
	ErrUnsupported    = errors.New("unsupported task type")
)
