package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/theus-ortiz/api-go/internal/app/controllers"
	"github.com/theus-ortiz/api-go/internal/app/services"
	"github.com/theus-ortiz/api-go/internal/app/repositories"
	"github.com/theus-ortiz/api-go/internal/db"
)

func AuthRoutes(router *gin.Engine) {
	// Inicializa o repositório e o serviço
	authRepo := repositories.NewAuthRepository(db.InitDB())
	authService := services.NewAuthService(authRepo)
	authController := controllers.NewAuthController(authService)

	// Define as rotas
	auth := router.Group("/auth")
	{
		auth.POST("/login", authController.Login)
	}
}