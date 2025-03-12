package groupmember

import (
	"family-planner/backend/internal/user"
	"family-planner/backend/internal/user/group"
)

type GroupMemberService struct {
	repository GroupMemberRepository
}

func NewGroupMemberService(repository GroupMemberRepository) *GroupMemberService {
	return &GroupMemberService{
		repository: repository,
	}
}

func (g *GroupMemberService) AddMember(group group.Group, user user.User) (*GroupMember, error) {
	newGroupMember := NewGroupMember(group, user)
	return g.repository.Save(&newGroupMember)
}
