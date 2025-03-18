package routes

import (
    "github.com/theus-ortiz/api-go/controllers"
    "github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
    user := router.Group("/user")
    {
        user.POST("/", controllers.CreateUser)
        user.GET("/", controllers.GetAllUsers)
        // Add routes for Read, Update, Delete
    }
}