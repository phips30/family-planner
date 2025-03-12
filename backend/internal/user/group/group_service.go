package group

import (
	"errors"
	"fmt"

	"family-planner/backend/internal/user"
)

type GroupService struct {
	repository     GroupRepository
	userRepository user.UserRepository
}

func NewGroupService(repository GroupRepository, userRepository user.UserRepository) *GroupService {
	return &GroupService{
		repository:     repository,
		userRepository: userRepository,
	}
}

func (g *GroupService) CreateGroup(groupName string, requestedByEmail string) (*Group, error) {
	validation_error := g.validate(groupName, requestedByEmail)
	if validation_error != nil {
		return nil, fmt.Errorf("%s", validation_error.Error())
	}

	requester, err := g.userRepository.FindByEmail(requestedByEmail)
	if err != nil {
		return nil, err
	}

	group := NewGroup(groupName, *requester)

	return g.repository.Save(&group)
}

func (g *GroupService) validate(groupName string, requestedByEmail string) error {
	userAlreadyInGroup := g.repository.FindGroupForUser(requestedByEmail)

	if userAlreadyInGroup != nil {
		return errors.New("user is already in a group")
	}
	if groupName == "" {
		return errors.New("group name cannot be empty")
	}
	return nil
}
