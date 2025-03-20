package services

import (
	rest_err "github.com/theus-ortiz/api-go/internal/app/config/restErr"
	"github.com/theus-ortiz/api-go/internal/app/models/requests"
	"github.com/theus-ortiz/api-go/internal/app/repositories"
	"github.com/theus-ortiz/api-go/internal/app/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(loginReq requests.LoginRequest) (string, *rest_err.RestErr)
}

type authService struct {
	authRepo repositories.AuthRepository
}

func NewAuthService(authRepo repositories.AuthRepository) AuthService {
	return &authService{authRepo: authRepo}
}

func (s *authService) Login(loginReq requests.LoginRequest) (string, *rest_err.RestErr) {
	// Buscar o usuário no banco de dados pelo email
	user, err := s.authRepo.FindUserByEmail(loginReq.Email)
	if err != nil {
		return "", rest_err.NewInternalServerError("Failed to find user")
	}
	if user == nil {
		return "", rest_err.NewUnauthorizedError("Invalid credentials", []rest_err.Causes{
			{
				Field:   "email",
				Message: "The email or password you have entered is invalid.",
			},
		})
	}

	// Comparar a senha fornecida com a senha criptografada no banco de dados
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password))
	if err != nil {
	return "", rest_err.NewUnauthorizedError("Invalid credentials", []rest_err.Causes{
		{
			Field:   "password",
			Message: "The email or password you have entered is invalid.",
		},
	})
}

	// Gerar o token JWT
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return "", rest_err.NewInternalServerError("Failed to generate token")
	}

	return token, nil
}
