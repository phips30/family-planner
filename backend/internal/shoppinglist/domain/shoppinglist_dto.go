package domain

import (
	"family-planner/backend/internal/common/models"
	"time"

	"github.com/google/uuid"
)

type ShoppinglistDto struct {
	Id      string
	Name    string
	UserId  uuid.UUID
	GroupId uuid.UUID
	AddedAt time.Time
	Bought  bool
}

type ShoppinglistItemResponse struct {
	Id      string
	Name    string
	AddedAt time.Time
	User    models.UserDto
	Bought  bool
}
