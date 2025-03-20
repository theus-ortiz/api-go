package config

import (
	"os"

	"github.com/theus-ortiz/api-go/internal/app/config/logger"
)

const jwtSecretEnv = "JWT_SECRET"

// JwtSecret retorna a chave JWT a partir da variável de ambiente
func JwtSecret() []byte {
	jwtSecret := os.Getenv(jwtSecretEnv)
	if jwtSecret == "" {
		logger.Error("Chave JWT não encontrada", nil)
		return nil
	}

	return []byte(jwtSecret)
}
