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

func (s *UserService) CreateUser(name string, deviceId string) (*User, error) {
	validation_error := s.validate(name, deviceId)
	if validation_error != nil {
		return nil, fmt.Errorf("%s", validation_error.Error())
	}

	user, err := NewUser(name, deviceId)
	if err != nil {
		return nil, err
	}
	return s.repository.Create(user)
}

func (s *UserService) validate(name string, deviceId string) error {
	user, err := s.repository.FindByNameAndDeviceId(name, deviceId)
	if err != nil {
		fmt.Printf("Error querying user")
	}
	if user != nil {
		return errors.New("User already exists in db")
	}
	return nil
}
