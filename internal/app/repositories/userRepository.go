package repositories

import (
	"database/sql"

	"github.com/theus-ortiz/api-go/internal/app/models"
	"github.com/theus-ortiz/api-go/internal/app/models/requests"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	GetAll() ([]models.User, error)
	GetByEmail(email string) (*models.User, error)
	Create(userReq requests.UserRequest) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetAll() ([]models.User, error) {
	stmt, err := r.db.Prepare("SELECT id, name, email, created_at FROM users")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *userRepository) GetByEmail(email string) (*models.User, error) {
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

func (r *userRepository) Create(userReq requests.UserRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	stmt, err := r.db.Prepare("INSERT INTO users (name, email, password) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userReq.Name, userReq.Email, string(hashedPassword))
	return err
}
