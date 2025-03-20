package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/theus-ortiz/api-go/internal/app/controllers"
	"github.com/theus-ortiz/api-go/internal/app/services"
	"github.com/theus-ortiz/api-go/internal/app/repositories"
	"github.com/theus-ortiz/api-go/internal/db"
)

func UserRoutes(router *gin.Engine) {
	// Inicializa o repositório e o serviço
	userRepo := repositories.NewUserRepository(db.InitDB())
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	// Define as rotas
	user := router.Group("/user")
	{
		user.POST("/", userController.CreateUser)
		user.GET("/", userController.GetAllUsers)
	}
}