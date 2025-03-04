package user

import (
	"fmt"
)

type UserService struct {
	repository UserRepository
}

func NewUserService(repository UserRepository) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) CreateUser(name string, deviceId string) (*User, error) {
	if user, err := s.repository.FindByNameAndDeviceId(name, deviceId); err != nil && user != nil {
		fmt.Printf("An error occured")
	}

	return nil, nil
}
