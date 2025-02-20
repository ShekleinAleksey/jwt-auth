package repository

import (
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	AuthRepository *AuthRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		AuthRepository: NewAuthRepository(db),
	}
}
