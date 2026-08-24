package service

import (
	"errors"
	"os"
	"time"

	"asri-backend/internal/model"
	"asri-backend/internal/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Register(input model.RegisterInput) (*model.User, error) {
	if _, err := s.userRepo.FindByEmail(input.Email); err == nil {
		return nil, ErrConflict
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &model.User{
		Name:         input.Name,
		Email:        input.Email,
		PasswordHash: string(hash),
		Role:         input.Role,
	}
	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *AuthService) Login(input model.LoginInput) (string, *model.User, error) {
	user, err := s.userRepo.FindByEmail(input.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", nil, ErrUnauthorized
		}
		return "", nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
		return "", nil, ErrUnauthorized
	}
	token, err := s.generateToken(user)
	if err != nil {
		return "", nil, err
	}
	return token, user, nil
}

func (s *AuthService) ListUsers() ([]model.User, error) {
	return s.userRepo.FindAll()
}

func (s *AuthService) ListUsersPaginated(page, limit int) ([]model.User, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}
	offset := (page - 1) * limit
	return s.userRepo.FindAllPaginated(limit, offset)
}

func (s *AuthService) GetUserByID(id uint) (*model.User, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return user, nil
}

func (s *AuthService) DeleteUser(id uint) error {
	if _, err := s.GetUserByID(id); err != nil {
		return err
	}
	return s.userRepo.Delete(id)
}

func (s *AuthService) generateToken(user *model.User) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "dev-secret-change-me"
	}
	claims := jwt.MapClaims{
		"sub":  user.ID,
		"email": user.Email,
		"role": user.Role,
		"exp":  time.Now().Add(72 * time.Hour).Unix(),
		"iat":  time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
