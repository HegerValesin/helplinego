package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/hegervalesin/helplinego/internal/models"
	"github.com/hegervalesin/helplinego/internal/repository"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{UserRepo: userRepo}
}

func (s *UserService) Create(user models.User) (*models.User, error) {
	if user.Login == "" {
		return nil, errors.New("Login é obrigatório")
	}

	// Verificar se já existe usuário com este email
	existingUser, _ := s.UserRepo.FindByLogin(user.Login)
	if existingUser != nil {
		return nil, errors.New("Login já cadastrado")
	}

	err := s.UserRepo.Create(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) GetByID(id uuid.UUID) (*models.User, error) {
	user, err := s.UserRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("usuário não encontrado")
	}
	return user, nil
}

func (s *UserService) Update(user models.User) (*models.User, error) {
	err := s.UserRepo.Update(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) Delete(id uuid.UUID) error {
	return s.UserRepo.Delete(id)
}

func (s *UserService) List(page, perPage int) ([]models.User, int64, error) {
	offset := (page - 1) * perPage
	return s.UserRepo.List(perPage, offset)
}
