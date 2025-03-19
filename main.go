package main

import (
    "github.com/theus-ortiz/api-go/db"
    "github.com/theus-ortiz/api-go/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    db.InitDB()

    routes.AuthRoutes(router)
    routes.UserRoutes(router)

    router.Run(":8080")
}