package service

import (
	"errors"
	"os"
	"time"

	"asri-backend/internal/model"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserStore is what teacher auth needs — fakes welcome, no DB required in tests.
type UserStore interface {
	FindByEmail(email string) (*model.User, error)
	Create(user *model.User) error
}

type AuthService struct {
	users UserStore
}

func NewAuthService(users UserStore) *AuthService {
	return &AuthService{users: users}
}

func jwtSecret() string {
	if s := os.Getenv("JWT_SECRET"); s != "" {
		return s
	}
	return "dev-secret-change-me"
}

// TeacherLogin authenticates a teacher, issuing a kind=teacher JWT (24h).
func (s *AuthService) TeacherLogin(input model.TeacherLoginInput) (string, *model.User, error) {
	user, err := s.users.FindByEmail(input.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", nil, ErrUnauthorized
		}
		return "", nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
		return "", nil, ErrUnauthorized
	}
	token, err := signToken(jwt.MapClaims{
		"sub":   user.ID,
		"email": user.Email,
		"kind":  "teacher",
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
		"iat":   time.Now().Unix(),
	})
	if err != nil {
		return "", nil, err
	}
	return token, user, nil
}

// SessionToken mints a short-lived kind=session JWT. Unscoped by design:
// one classroom code opens the whole book (no per-chapter gating).
func SessionToken(ttlHours int) (string, error) {
	if ttlHours < 1 {
		ttlHours = 12
	}
	return signToken(jwt.MapClaims{
		"kind": "session",
		"exp":  time.Now().Add(time.Duration(ttlHours) * time.Hour).Unix(),
		"iat":  time.Now().Unix(),
	})
}

// SeedTeacher creates the teacher account on first boot when env provides credentials.
func (s *AuthService) SeedTeacher(email, password string) (*model.User, error) {
	if email == "" || password == "" {
		return nil, nil // nothing configured — skip silently
	}
	if u, err := s.users.FindByEmail(email); err == nil {
		return u, nil // already seeded
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &model.User{Email: email, PasswordHash: string(hash)}
	if err := s.users.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func signToken(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret()))
}
