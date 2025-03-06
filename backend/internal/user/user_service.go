package user

import (
	"errors"
	"fmt"
)

// This is the user domain service

type UserService struct {
	repository UserRepository
}

func NewUserService(repository UserRepository) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) CreateUser(name string, email string) (*User, error) {
	validation_error := s.validate(email)
	if validation_error != nil {
		return nil, fmt.Errorf("%s", validation_error.Error())
	}

	user, err := NewUser(name, email)
	if err != nil {
		return nil, err
	}
	return s.repository.Create(user)
}

func (s *UserService) FindByEmail(email string) (*User, error) {
	return s.repository.FindByEmail(email)
}

func (s *UserService) validate(email string) error {
	user, err := s.FindByEmail(email)
	if err != nil {
		fmt.Printf("Error querying user")
	}
	if user != nil {
		return errors.New("User already exists")
	}
	return nil
}
