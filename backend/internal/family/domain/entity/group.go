package entity

import (
	"time"

	"github.com/google/uuid"
)

type Group struct {
	Id        uuid.UUID
	Name      string
	CreatedBy User
	CreatedAt time.Time
}

func NewGroup(name string, createdBy User) Group {
	return Group{
		Id:        uuid.New(),
		Name:      name,
		CreatedBy: createdBy,
		CreatedAt: time.Now(),
	}
}
