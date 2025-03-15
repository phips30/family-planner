package aggregate

import (
	"errors"
	"family-planner/backend/internal/family/domain/entity"
)

var (
	ErrNoName             = errors.New("group has no name")
	ErrNoCreator          = errors.New("group has no creator")
	ErrUserAlreadyInGroup = errors.New("user is already in group")
)

type GroupAgg struct {
	Group        *entity.Group
	GroupMembers []*entity.GroupMember
}

func CreateNewGroupAgg(name string, createdBy *entity.User) (*GroupAgg, error) {
	if name == "" {
		return nil, ErrNoName
	}
	if createdBy == nil {
		return nil, ErrNoCreator
	}

	groupAgg := &GroupAgg{
		Group:        entity.NewGroup(name, *createdBy),
		GroupMembers: make([]*entity.GroupMember, 0),
	}
	groupAgg.AddGroupMember(createdBy)

	return groupAgg, nil
}

func FromExistingGroupAgg(group *entity.Group, creator *entity.User, groupMembers []*entity.GroupMember) *GroupAgg {
	for _, groupMember := range groupMembers {
		groupMember.Status = entity.Exists
	}

	group.CreatedBy = *creator
	return &GroupAgg{
		Group:        group,
		GroupMembers: groupMembers,
	}
}

func (g *GroupAgg) AddGroupMember(newMember *entity.User) error {
	for _, groupMember := range g.GroupMembers {
		if groupMember.User.Id == newMember.Id {
			return ErrUserAlreadyInGroup
		}
	}
	g.GroupMembers = append(g.GroupMembers, entity.NewGroupMember(newMember))
	return nil
}

func (g *GroupAgg) GetNewGroupMembers() []*entity.GroupMember {
	var newGroupMembers []*entity.GroupMember
	for _, groupMember := range g.GroupMembers {
		if groupMember.Status == entity.New {
			newGroupMembers = append(newGroupMembers, groupMember)
		}
	}
	return newGroupMembers
}
