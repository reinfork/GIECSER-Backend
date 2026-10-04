package model

import "time"

// AccessCode is a 6-character PIN a teacher issues. One code opens the whole
// book — no per-chapter scoping (classroom reality: shared code on the board).
// Redis caches code validity; this table is authoritative.
type AccessCode struct {
	ID        string    `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Code      string    `json:"code" gorm:"uniqueIndex;not null"`
	CreatedBy *string   `json:"created_by" gorm:"type:uuid"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

type ValidatePINInput struct {
	Code string `json:"code" binding:"required,len=6"`
}

type IssuePINInput struct {
	TTLHours int `json:"ttl_hours"` // optional, defaults to 24
}
