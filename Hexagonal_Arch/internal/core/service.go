// internal/core/service.go
package core

import "hexagonal-architecture/internal/ports"

type UserService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *User) error {
	return s.repo.Save(user)
}

func (s *UserService) GetUser(id int) (*User, error) {
	return s.repo.FindByID(id)
}

// OCP (Open-Closed Principle): The service can support different repositories without modification.
