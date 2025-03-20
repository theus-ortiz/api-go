package repositories

import (
	"database/sql"

	"github.com/theus-ortiz/api-go/internal/app/models"
)

type AuthRepository interface {
	FindUserByEmail(email string) (*models.User, error)
}

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	query := "SELECT id, email, password FROM users WHERE email = ?"
	row := r.db.QueryRow(query, email)
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}
