package entity

import (
	"time"
)

type GroupMember struct {
	Group     Group
	User      User
	CreatedAt time.Time
}

func NewGroupMember(group Group, user User) GroupMember {
	return GroupMember{
		Group:     group,
		User:      user,
		CreatedAt: time.Now(),
	}
}
