package repository

import "family-planner/backend/internal/family/domain/entity"

type GroupMemberRepository interface {
	Save(groupMember *entity.GroupMember) (*entity.GroupMember, error)
}
