package service

import (
	"crypto/sha256"
	"fmt"
	hikarilibbackend "github.com/yuminekosan/hikariLibBackend"
	"github.com/yuminekosan/hikariLibBackend/internal/repository"
)

const salt = "dfsdgksg234k2ladas"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user hikarilibbackend.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
