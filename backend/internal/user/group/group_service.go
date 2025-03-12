package group

import (
	"errors"
	"fmt"

	"family-planner/backend/internal/user"
	"family-planner/backend/internal/user/groupmember"
)

type GroupService struct {
	repository         GroupRepository
	userService        user.UserService
	groupMemberService groupmember.GroupMemberService
}

func NewGroupService(repository GroupRepository, userService user.UserService, groupMemberService groupmember.GroupMemberService) *GroupService {
	return &GroupService{
		repository:         repository,
		userService:        userService,
		groupMemberService: groupMemberService,
	}
}

func (g *GroupService) CreateGroup(groupName string, requestedByEmail string) (*Group, error) {
	validation_error := g.validate(groupName, requestedByEmail)
	if validation_error != nil {
		return nil, fmt.Errorf("%s", validation_error.Error())
	}

	requester, err := g.userService.FindByEmail(requestedByEmail)
	if err != nil {
		return nil, err
	}

	group := NewGroup(groupName, *requester)

	addedGroup, err := g.repository.Save(&group)
	if err != nil {
		return nil, err
	}
	_, err = g.groupMemberService.AddMember(addedGroup, *requester)
	if err != nil {
		return nil, err
	}

	return addedGroup, nil
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
