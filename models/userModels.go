package models

import (
	"time"

	"github.com/theus-ortiz/api-go/config/restErr"
	"github.com/theus-ortiz/api-go/config/validation"
	"github.com/theus-ortiz/api-go/db"
	"github.com/theus-ortiz/api-go/models/requests"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

func GetAllUsers() ([]User, error) {
	// Prepara o statement para selecionar todos os usuários
	stmt, err := db.InitDB().Prepare("SELECT id, name, email, created_at FROM users")
	if err != nil {
		return nil, err
	}
	defer stmt.Close() // Fecha o statement após o uso

	// Executa o statement e guarda o resultado em rows
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Fecha o resultado após o uso

	// Cria um slice de usuários para guardar o resultado
	users := []User{}
	for rows.Next() {
		var u User
		// Faz o scan dos valores do resultado para a struct User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt); err != nil {
			return nil, err
		}
		// Adiciona o usuário ao slice
		users = append(users, u)
	}
	return users, nil
}

func (u *User) CreateUser(userReq requests.UserRequest) *rest_err.RestErr {
	// Gera o hash da senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
	if err != nil {
		return rest_err.NewInternalServerError("Failed to generate password hash")
	}

	// Prepara o statement para inserir o usuário
	stmt, err := db.InitDB().Prepare("INSERT INTO users (name, email, password) VALUES (?, ?, ?)")
	if err != nil {
		return rest_err.NewInternalServerError("Failed to prepare SQL statement")
	}
	defer stmt.Close()

	// Executa o statement
	_, err = stmt.Exec(userReq.Name, userReq.Email, string(hashedPassword))
	if err != nil {
		// Trata o erro usando a função ValidateUserError
		return validation.ValidateUserError(err)
	}

	return nil

}
