package service

import (
	"errors"
	"slices"
	"sp-backend/internal/domain"
)

type UserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAll() ([]domain.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) UpdateRole(id string, role string) (*domain.User, error) {
	if !slices.Contains(domain.ValidRoles, role) {
		return nil, errors.New("invalid role")
	}
	return s.repo.UpdateRole(id, role)
}
