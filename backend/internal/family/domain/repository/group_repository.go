package repository

import "family-planner/backend/internal/family/domain/entity"

type GroupRepository interface {
	Save(group *entity.Group) (*entity.Group, error)
	FindGroupForUser(email string) *entity.Group
}
