package service

import (
	"qfzack/go-web-starter/internal/server/model"
	"qfzack/go-web-starter/internal/server/repository"
)

// define service funcs through interface
type UserService interface {
	GetUser(id string) (*model.User, error)
}

/*
define struct userService
implement the pre-defined UserService interface
*/
type userService struct {
	// user pre-defined interface in repository
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetUser(id string) (*model.User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		// determine the error type is sql.ErrNoRows
		return nil, model.ErrUserNotFound
	}
	return user, nil
}
