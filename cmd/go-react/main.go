package main

import (
    "github.com/theus-ortiz/api-go/internal/db"
    "github.com/theus-ortiz/api-go/internal/app/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    db.InitDB()

    routes.AuthRoutes(router)
    routes.UserRoutes(router)

    router.Run(":8080")
}