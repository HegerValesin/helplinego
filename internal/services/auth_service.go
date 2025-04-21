package services

import (
	"errors"

	"github.com/hegervalesin/helplinego/internal/models"
	"github.com/hegervalesin/helplinego/internal/repository"
	"github.com/hegervalesin/helplinego/pkg/auth"
)

type AuthService struct {
	UserRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{UserRepo: userRepo}
}

func (s *AuthService) Login(login, password string) (string, *models.User, error) {
	user, err := s.UserRepo.FindByLogin(login)
	if err != nil {
		return "", nil, errors.New("invalid credentials")
	}

	if !user.CheckPassword(password) {
		return "", nil, errors.New("invalid credentials")
	}

	token, err := auth.GenerateJWT(*user)
	if err != nil {
		return "", nil, errors.New("failed to generate token")
	}

	return token, user, nil
}

func (s *AuthService) Register(user models.User) error {
	// Verificar se o usuário já existe
	existingUser, err := s.UserRepo.FindByLogin(user.Login)
	if err == nil && existingUser != nil {
		return errors.New("Login already registered")
	}

	// Define um papel padrão se não for especificado
	if user.Type == "" {
		user.Type = "user"
	}

	// Salva o novo usuário
	return s.UserRepo.Create(&user)
}
