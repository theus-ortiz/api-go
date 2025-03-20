package services

import (
	rest_err "github.com/theus-ortiz/api-go/internal/app/config/restErr"
	"github.com/theus-ortiz/api-go/internal/app/models"
	"github.com/theus-ortiz/api-go/internal/app/models/requests"
	"github.com/theus-ortiz/api-go/internal/app/repositories"
)

type UserService interface {
	GetAllUsers() ([]models.User, error)
	CreateUser(userReq requests.UserRequest) *rest_err.RestErr
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.userRepo.GetAll()
}

func (s *userService) CreateUser(userReq requests.UserRequest) *rest_err.RestErr {
	// Verifica se o e-mail já está em uso
	existingUser, err := s.userRepo.GetByEmail(userReq.Email)
	if err != nil {
		return rest_err.NewInternalServerError("Failed to check email availability")
	}
	if existingUser != nil {
		return rest_err.NewBadRequestError("Email already in use")
	}

	// Cria o usuário no banco de dados
	if err := s.userRepo.Create(userReq); err != nil {
		return rest_err.NewInternalServerError("Failed to create user")
	}

	return nil
}
