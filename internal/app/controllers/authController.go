package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/theus-ortiz/api-go/internal/app/config/logger"
	"github.com/theus-ortiz/api-go/internal/app/models/requests"
	"github.com/theus-ortiz/api-go/internal/app/models/responses"
	"github.com/theus-ortiz/api-go/internal/app/services"
	"go.uber.org/zap"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (c *AuthController) Login(ctx *gin.Context) {
	var loginReq requests.LoginRequest

	// Faz o bind dos dados da requisição para a struct LoginRequest
	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		logger.Error("Error trying to validate login info", err,
			zap.String("journey", "login"),
		)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Realiza o login usando o serviço de autenticação
	token, err := c.authService.Login(loginReq)
	if err != nil {
		logger.Error("Error trying to login", err,
			zap.String("journey", "login"),
		)
		ctx.JSON(err.Code, err)
		return
	}

	// Configurar o cookie
	ctx.SetSameSite(http.SameSiteStrictMode)
	ctx.SetCookie("jwt", token, int(time.Hour*72), "/", "localhost", false, true)

	// Retornar uma resposta de sucesso
	logger.Info("Login Successful!",
		zap.String("journey", "login"))
	ctx.JSON(http.StatusOK, responses.AuthResponse{
		Message: "Login Successful!",
		Token:   token,
	})
}