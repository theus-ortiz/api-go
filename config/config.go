package config

import (
	"os"

	"github.com/theus-ortiz/api-go/config/logger"
)

func JwtSecret() []byte {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		logger.Error("Chave JWT não encontrada", nil)
	}

	logger.Info("Chave JWT encontrada")
	return []byte(jwtSecret)
}