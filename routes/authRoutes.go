package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/theus-ortiz/api-go/controllers"
)

func AuthRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/login", controllers.Login)
	}
}
