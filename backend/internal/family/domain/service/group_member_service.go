package service

import (
	"family-planner/backend/internal/family/domain/entity"
	"family-planner/backend/internal/family/domain/repository"
)

type GroupMemberService struct {
	repository repository.GroupMemberRepository
}

func NewGroupMemberService(repository repository.GroupMemberRepository) *GroupMemberService {
	return &GroupMemberService{
		repository: repository,
	}
}

func (g *GroupMemberService) AddMember(group entity.Group, user entity.User) (*entity.GroupMember, error) {
	newGroupMember := entity.NewGroupMember(group, user)
	return g.repository.Save(&newGroupMember)
}
