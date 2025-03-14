package entity

import (
	"time"
)

type EntityStatus int

const (
	Exists EntityStatus = iota
	New
	Delete
)

type GroupMember struct {
	User      *User
	CreatedAt time.Time
	Status    EntityStatus
}

func NewGroupMember(user *User) *GroupMember {
	return &GroupMember{
		User:      user,
		CreatedAt: time.Now(),
		Status:    New,
	}
}
