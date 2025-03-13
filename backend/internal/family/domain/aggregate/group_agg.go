package aggregate

import "family-planner/backend/internal/family/domain/entity"

type GroupAgg struct {
	Group        *entity.Group
	GroupMembers []*entity.User
}
