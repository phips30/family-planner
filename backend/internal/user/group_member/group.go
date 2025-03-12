package group

import (
	"family-planner/backend/internal/user"
	"time"

	"go.mongodb.org/mongo-driver/internal/uuid"
)

type Group struct {
	Id        uuid.UUID
	Name      string
	CreatedBy user.User
	CreatedAt time.Time
}

func NewGroup(name string, createdBy user.User) Group {
	return Group{
		Id:        uuid.New(),
		Name:      name,
		CreatedBy: createdBy,
		CreatedAt: time.Now(),
	}
}
