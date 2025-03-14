package repository

import (
	"errors"
	"family-planner/backend/internal/family/domain/aggregate"
	"family-planner/backend/internal/family/domain/entity"

	"github.com/google/uuid"
)

var (
	ErrGroupNotFound = errors.New("group not found")
)

type GroupRepository interface {
	Find(groupId uuid.UUID) (*aggregate.GroupAgg, error)
	FindGroupForUser(email string) *entity.Group
	Save(group *entity.Group) (*entity.Group, error)
	SaveGroupMember(group *entity.Group, groupMembers []*entity.GroupMember) error
}
