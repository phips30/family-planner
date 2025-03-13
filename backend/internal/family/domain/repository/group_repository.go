package repository

import (
	"errors"
	"family-planner/backend/internal/family/domain/entity"

	"github.com/google/uuid"
)

var (
	ErrGroupNotFound = errors.New("group not found")
)

type GroupRepository interface {
	Save(group *entity.Group) (*entity.Group, error)
	Find(groupId uuid.UUID) (*entity.Group, error)
	FindGroupForUser(email string) *entity.Group
}
