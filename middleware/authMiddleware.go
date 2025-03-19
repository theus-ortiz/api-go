package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/theus-ortiz/api-go/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("jwt")
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Token de autenticação necessário"})
            c.Abort()
            return
        }

		token, err := utils.ValidateJWT(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		c.Next()
	}
}
