package service

import (
	"github.com/ShekleinAleksey/jwt-auth/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) DeleteUser(userId string) error {
	return s.repo.DeleteUser(userId)
}
