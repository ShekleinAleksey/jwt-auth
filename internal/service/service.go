package service

import "github.com/ShekleinAleksey/jwt-auth/internal/repository"

type Service struct {
	AuthService *AuthService
	UserService *UserService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		AuthService: NewAuthService(*repo.AuthRepository),
		UserService: NewUserService(*repo.UserRepository),
	}
}
