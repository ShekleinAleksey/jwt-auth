package repository

import (
	"fmt"

	"github.com/ShekleinAleksey/jwt-auth/internal/entity"
	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}
func (r *AuthRepository) CreateUser(user entity.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO users (first_name, last_name, email, password_hash) values ($1, $2, $3, $4) RETURNING id")

	row := r.db.QueryRow(query, user.FirstName, user.LastName, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthRepository) GetUser(email, password string) (entity.User, error) {
	var user entity.User
	query := fmt.Sprintf("SELECT id FROM users WHERE email=$1 AND password_hash=$2")
	err := r.db.Get(&user, query, email, password)

	return user, err
}
