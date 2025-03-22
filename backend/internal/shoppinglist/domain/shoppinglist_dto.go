package domain

import (
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
