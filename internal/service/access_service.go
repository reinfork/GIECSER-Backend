package service

import (
	"context"
	"crypto/rand"
	"errors"
	"time"

	"asri-backend/internal/model"

	"gorm.io/gorm"
)

// CodeStore + PinCache are what PIN issue/validate needs — fakes welcome.
type CodeStore interface {
	Create(code *model.AccessCode) error
	FindByCode(code string) (*model.AccessCode, error)
	ListActive(now time.Time) ([]model.AccessCode, error)
}

type PinCache interface {
	SetPIN(ctx context.Context, code string, ttl time.Duration) error
	GetPIN(ctx context.Context, code string) (bool, error)
}

type AccessService struct {
	codes           CodeStore
	cache           PinCache
	sessionTTLHours int
}

func NewAccessService(codes CodeStore, cache PinCache, sessionTTLHours int) *AccessService {
	if sessionTTLHours < 1 {
		sessionTTLHours = 12
	}
	return &AccessService{codes: codes, cache: cache, sessionTTLHours: sessionTTLHours}
}

// alphabet drops confusables (0/O/1/I) — codes get read aloud in classrooms.
const pinAlphabet = "ABCDEFGHJKMNPQRSTUVWXYZ23456789"

func generatePIN() (string, error) {
	b := make([]byte, 6)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	for i := range b {
		b[i] = pinAlphabet[int(b[i])%len(pinAlphabet)]
	}
	return string(b), nil
}

// IssuePIN mints a 6-char code opening the whole book (default 24h TTL).
func (s *AccessService) IssuePIN(ctx context.Context, ttlHours int, teacherID *string) (*model.AccessCode, error) {
	if ttlHours < 1 {
		ttlHours = 24
	}
	code, err := generatePIN()
	if err != nil {
		return nil, err
	}
	now := time.Now()
	ac := &model.AccessCode{
		Code: code, CreatedBy: teacherID,
		CreatedAt: now, ExpiresAt: now.Add(time.Duration(ttlHours) * time.Hour),
	}
	if err := s.codes.Create(ac); err != nil {
		return nil, err
	}
	_ = s.cache.SetPIN(ctx, code, time.Until(ac.ExpiresAt)) // cache best-effort; DB authoritative
	return ac, nil
}

// ValidatePIN checks cache → DB, then mints an unscoped session JWT.
func (s *AccessService) ValidatePIN(ctx context.Context, code string) (token string, expiresAt time.Time, err error) {
	if ok, err := s.cache.GetPIN(ctx, code); err == nil && ok {
		return s.mintSession()
	}
	ac, err := s.codes.FindByCode(code)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", time.Time{}, ErrUnauthorized
		}
		return "", time.Time{}, err
	}
	if time.Now().After(ac.ExpiresAt) {
		return "", time.Time{}, ErrExpired
	}
	_ = s.cache.SetPIN(ctx, code, time.Until(ac.ExpiresAt))
	return s.mintSession()
}

func (s *AccessService) mintSession() (string, time.Time, error) {
	token, err := SessionToken(s.sessionTTLHours)
	if err != nil {
		return "", time.Time{}, err
	}
	return token, time.Now().Add(time.Duration(s.sessionTTLHours) * time.Hour), nil
}

func (s *AccessService) ListActivePINs() ([]model.AccessCode, error) {
	return s.codes.ListActive(time.Now())
}
