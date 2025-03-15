package service

import (
	"errors"
	"family-planner/backend/internal/family/domain/aggregate"
	"family-planner/backend/internal/family/domain/entity"
	"family-planner/backend/internal/family/domain/repository"

	"github.com/google/uuid"
)

type GroupService struct {
	repository  repository.GroupRepository
	userService UserService
}

var (
	ErrUserAlreadyInGroup = errors.New("user already in group")
)

func NewGroupService(repository repository.GroupRepository, userService UserService) *GroupService {
	return &GroupService{
		repository:  repository,
		userService: userService,
	}
}

func (g *GroupService) CreateGroup(groupName string, requestedByEmail string) (*entity.Group, error) {
	userAlreadyInGroup, _ := g.repository.FindGroupForUser(requestedByEmail)
	if userAlreadyInGroup != nil {
		return nil, errors.New("user is already in a group")
	}

	requester, err := g.userService.FindByEmail(requestedByEmail)
	if err != nil {
		return nil, err
	}
	groupAgg, err := aggregate.CreateNewGroupAgg(groupName, requester)
	if err != nil {
		return nil, err
	}

	addedGroup, err := g.repository.Save(groupAgg.Group)
	if err != nil {
		return nil, err
	}
	err = g.repository.SaveGroupMember(groupAgg.Group, groupAgg.GetNewGroupMembers())
	if err != nil {
		return nil, err
	}

	return addedGroup, nil
}

func (g GroupService) AddGroupMember(groupId uuid.UUID, requestedByEmail string) (*aggregate.GroupAgg, error) {
	groupAgg, err := g.repository.Find(groupId)
	if err != nil {
		return nil, err
	}

	requester, err := g.userService.FindByEmail(requestedByEmail)
	if err != nil {
		return nil, err
	}

	userAlreadyInAnyGroup, _ := g.repository.FindGroupForUser(requestedByEmail)
	if userAlreadyInAnyGroup != nil {
		return nil, ErrUserAlreadyInGroup
	}

	err = groupAgg.AddGroupMember(requester)
	if err != nil {
		return groupAgg, err
	}
	err = g.repository.SaveGroupMember(groupAgg.Group, groupAgg.GetNewGroupMembers())
	return groupAgg, err
}

func (g GroupService) FindGroupMembers(groupId uuid.UUID) (*aggregate.GroupAgg, error) {
	return g.repository.Find(groupId)
}

func (g GroupService) FindGroupMembersByEmail(email string) (*aggregate.GroupAgg, error) {
	return g.repository.FindGroupForUser(email)
}
