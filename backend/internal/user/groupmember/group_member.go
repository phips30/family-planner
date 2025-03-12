package groupmember

import (
	"family-planner/backend/internal/user"
	"family-planner/backend/internal/user/group"
	"time"
)

type GroupMember struct {
	Group     group.Group
	User      user.User
	CreatedAt time.Time
}

func NewGroupMember(group group.Group, user user.User) GroupMember {
	return GroupMember{
		Group:     group,
		User:      user,
		CreatedAt: time.Now(),
	}
}
