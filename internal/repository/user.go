package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) DeleteUser(userId string) error {
	query := fmt.Sprintf("DELETE FROM users WHERE id = $1")
	_, err := r.db.Exec(query, userId)
	return err
}
