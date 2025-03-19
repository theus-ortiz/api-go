package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/theus-ortiz/api-go/config/logger"
	"github.com/theus-ortiz/api-go/config/validation"
	"github.com/theus-ortiz/api-go/models"
	"github.com/theus-ortiz/api-go/models/requests"
	"github.com/theus-ortiz/api-go/models/responses"
	"github.com/theus-ortiz/api-go/utils"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"github.com/theus-ortiz/api-go/config/restErr"
)


// Função auxiliar para retornar erro de credenciais inválidas
func invalidCredentialsError() *rest_err.RestErr {
	return rest_err.NewUnauthorizedError("Invalid credentials", []rest_err.Causes{
		{
			Field:   "Invalid credentials",
			Message: "The email or password you have entered is invalid.",
		},
	})
}

func Login(c *gin.Context) {
	var loginReq requests.LoginRequest

	// Faz o bind dos dados da requisição para a struct LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		logger.Error("Error trying to validate login info", err,
			zap.String("journey", "login"),
		)
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	// Valida os dados da requisição
	if err := validate.Struct(loginReq); err != nil {
		logger.Error("Error trying to validate login info", err,
			zap.String("journey", "login"),
		)
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	// Buscar o usuário no banco de dados pelo email
	user, err := models.FindUserByEmail(loginReq.Email)
	if err != nil {
		logger.Error("Error trying to find user by email", err,
			zap.String("journey", "login"),
		)

		// Retorna erro de credenciais inválidas
		c.JSON(http.StatusUnauthorized, invalidCredentialsError())
		return
	}

	// Comparar a senha fornecida com a senha criptografada no banco de dados
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password))
	if err != nil {
		logger.Error("Invalid credentials", err,
			zap.String("journey", "login"),
		)

		// Retorna erro de credenciais inválidas
		c.JSON(http.StatusUnauthorized, invalidCredentialsError())
		return
	}

	// Gerar o token JWT
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		logger.Error("Error trying to generate JWT token", err,
			zap.String("journey", "tokenJWT"),
		)
		c.JSON(http.StatusInternalServerError, rest_err.NewInternalServerError("Error generating token"))
		return
	}

	// Configurar o cookie
	c.SetSameSite(http.SameSiteStrictMode) // envia o cookie apenas para o mesmo site
	c.SetCookie("jwt", token, int(time.Hour*72), "/", "localhost", false, true)

	// Retornar uma resposta de sucesso
	logger.Info("Login Successful!",
		zap.String("journey", "login"))
	c.JSON(http.StatusOK, responses.AuthResponse{
		Message: "Login Successful!",
	})
}