package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/theus-ortiz/api-go/config"
)

func GenerateJWT(userID int) (string, error) {
	// Criar o token com claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 72).Unix(), // Expira em 72 horas
	})

	// Assinar o token com a chave secreta
	tokenString, err := token.SignedString(config.JwtSecret())
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	// Validar o token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verificar o método de assinatura
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		// Retornar a chave secreta para validação
		return config.JwtSecret(), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}