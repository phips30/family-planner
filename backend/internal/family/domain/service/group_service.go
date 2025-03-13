package service

import (
	"errors"
	"family-planner/backend/internal/family/domain/entity"
	"family-planner/backend/internal/family/domain/repository"
	"fmt"

	"github.com/google/uuid"
)

type GroupService struct {
	repository         repository.GroupRepository
	userService        UserService
	groupMemberService GroupMemberService
}

func NewGroupService(repository repository.GroupRepository, userService UserService, groupMemberService GroupMemberService) *GroupService {
	return &GroupService{
		repository:         repository,
		userService:        userService,
		groupMemberService: groupMemberService,
	}
}

func (g *GroupService) CreateGroup(groupName string, requestedByEmail string) (*entity.Group, error) {
	validation_error := g.validate(groupName, requestedByEmail)
	if validation_error != nil {
		return nil, fmt.Errorf("%s", validation_error.Error())
	}

	requester, err := g.userService.FindByEmail(requestedByEmail)
	if err != nil {
		return nil, err
	}

	group := entity.NewGroup(groupName, *requester)

	addedGroup, err := g.repository.Save(&group)
	if err != nil {
		return nil, err
	}
	_, err = g.groupMemberService.AddMember(*addedGroup, *requester)
	if err != nil {
		return nil, err
	}

	return addedGroup, nil
}

func (g GroupService) AddGroupMember(groupId uuid.UUID, requestedByEmail string) (*entity.Group, error) {
	group, err := g.repository.Find(groupId)
	if err != nil {
		return nil, err
	}

	requester, err := g.userService.FindByEmail(requestedByEmail)
	if err != nil {
		return nil, err
	}

	userAlreadyInGroup := g.repository.FindGroupForUser(requestedByEmail)
	if userAlreadyInGroup != nil {
		return nil, errors.New("user is already in a group")
	}

	_, err = g.groupMemberService.AddMember(*group, *requester)
	if err != nil {
		return nil, err
	}

	return group, nil
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
